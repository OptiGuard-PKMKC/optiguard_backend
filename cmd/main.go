package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/OptiGuard-PKMKC/optiguard_backend/pkg/config"
	driver_db "github.com/OptiGuard-PKMKC/optiguard_backend/pkg/driver/db"

	_ "github.com/lib/pq"
)

func main() {
	env := config.LoadEnv()

	db, err := driver_db.NewConnection(env)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(db)

	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	}

	http.HandleFunc("/hello-world", helloHandler)

	log.Println("Server started at: ", env.AppPort)
	port := fmt.Sprintf(":%s", env.AppPort)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Println("Failed to start server: ", err)
	}
}
