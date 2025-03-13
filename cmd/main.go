package main

import (
	"fmt"
	"github.com/Babahasko/stat_api/configs"
	"github.com/Babahasko/stat_api/internal/auth"
	"github.com/Babahasko/stat_api/internal/link"
	"github.com/Babahasko/stat_api/internal/stat"
	"github.com/Babahasko/stat_api/internal/user"
	"github.com/Babahasko/stat_api/pkg/db"
	"github.com/Babahasko/stat_api/pkg/event"
	"github.com/Babahasko/stat_api/pkg/middleware"
	"net/http"
)
func App() http.Handler{
	conf := configs.LoadConfig()
	db := db.NewDB(conf)
	router := http.NewServeMux()
	eventBus := event.NewEventBus()

	//Repositories
	linkRepository := link.NewLinkRepository(db)

	userReposiory := user.NewUserRepository(db)
	statRepository := stat.NewStatRepository(db)

	//Services

	authService := auth.NewAuthService(userReposiory)
	statService := stat.NewStatService(&stat.StatServiceDeps{
		EventBus:       eventBus,
		StatRepository: statRepository,
	})

	go statService.AddClick()

	// Handler
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
		Config:         conf,
		EventBus:       eventBus,
	})
	stat.NewStatHandler(router, &stat.StatHandlerDeps{
		StatRepository: statRepository,
	})
	// Middlewares
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)
	return stack(router)
}

func main() {
	app := App()
	server := http.Server{
		Addr:    ":8081",
		Handler: app,
	}

	fmt.Println("Server is listening on port 8081")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
