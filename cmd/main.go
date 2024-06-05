package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/OptiGuard-PKMKC/optiguard_backend/internal/http/controllers"
	"github.com/OptiGuard-PKMKC/optiguard_backend/internal/http/routes"
	route_intf "github.com/OptiGuard-PKMKC/optiguard_backend/internal/http/routes/interfaces"
	"github.com/OptiGuard-PKMKC/optiguard_backend/pkg/config"
	driver_db "github.com/OptiGuard-PKMKC/optiguard_backend/pkg/driver/db"
	"github.com/OptiGuard-PKMKC/optiguard_backend/pkg/repositories"
	"github.com/OptiGuard-PKMKC/optiguard_backend/pkg/usecases"
)

func main() {
	log.Print("Starting server...")

	env := config.LoadEnv()

	db, err := driver_db.NewConnection(env)
	if err != nil {
		log.Println(err)
	}

	userRepo := repositories.NewDbUserRepository(db)

	authUsecase := usecases.NewAuthUsecase(env.SecretKey, userRepo)

	authController := controllers.NewAuthController(authUsecase)

	router := routes.SetupRouter(route_intf.Controllers{
		Auth: authController,
	})

	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	}
	http.HandleFunc("/hello-world", helloHandler)

	port := fmt.Sprintf(":%s", env.AppPort)

	err = http.ListenAndServe(port, router)
	if err != nil {
		log.Println("Failed to start server: ", err)
	}

	log.Println("Server started at: ", env.AppPort)
}
