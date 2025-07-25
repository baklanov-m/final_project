package server // функции для запуска веб-сервера.

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"todo/pkg/api"
)

const (
	defaultPort = "7540"
	webDir      = "./web"
)

func StartServer() { //старт веб сервера планировщика задач
	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = defaultPort
	}

	api.Init()

	http.Handle("/", http.FileServer(http.Dir(webDir)))

	fmt.Printf("Сервер запущен на порту %s\n", port)
	fmt.Printf("Веб-интерфейс: http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
