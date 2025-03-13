package auth_test

import (
	"bytes"
	"encoding/json"
	"github.com/Babahasko/stat_api/configs"
	"github.com/Babahasko/stat_api/internal/auth"
	"github.com/Babahasko/stat_api/internal/user"
	"github.com/Babahasko/stat_api/pkg/db"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func bootstrap() (*auth.AuthHandler, sqlmock.Sqlmock, error) {
	database, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	gormDb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: database,
	}))
	if err != nil {
		return nil, nil, err
	}
	userRepository := user.NewUserRepository(&db.Db{
		DB: gormDb,
	})
	handler := auth.AuthHandler{
		Config: &configs.Config{
			Auth: configs.AuthConfig{
				Secret: "secret",
			},
		},
		AuthService: auth.NewAuthService(userRepository),
	}
	return &handler, mock, nil
}
func TestLoginHandlerSucess(t *testing.T) {
	handler, mock, err := bootstrap()
	// Добавляем одну строку в виртуальную базу данных
	rows := sqlmock.NewRows([]string{"email", "password"}).
		AddRow("bob8@mail.ru", "$2a$10$iY3dGkePOnVpYIHIo429lezOqW0hPNDjAaKnyblCbIO4EBQJq9D6q")
	// Возвращаем эту таблицу
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	if err != nil {
		t.Fatal(err)
		return
	}
	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "bob8@mail.ru",
		Password: "11114",
	})
	reader := bytes.NewReader(data)
	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/auth/login", reader)
	handler.Login()(wr, req)
	if wr.Code != http.StatusOK {
		t.Errorf("Got %v expected %v", wr.Code, 200)
	}
}

func TestRegisterHandlerSuccess(t *testing.T) {
	handler, mock, err := bootstrap()
	//Сначала проверяем что юзера нет в БД
	rows := sqlmock.NewRows([]string{"email", "password", "name"}) // вернули пустого юзера
	// Возвращаем эту таблицу
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	mock.ExpectBegin()
	mock.ExpectQuery("INSERT").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()
	if err != nil {
		t.Fatal(err)
		return
	}
	data, _ := json.Marshal(&auth.RegisterRequest{
		Email:    "bob8@mail.ru",
		Password: "11114",
		Name: "Вася",
	})
	reader := bytes.NewReader(data)
	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/auth/register", reader)
	handler.Register()(wr, req)
	if wr.Code != http.StatusCreated {
		t.Errorf("Got %v expected %v", wr.Code, 201)
	}
}
