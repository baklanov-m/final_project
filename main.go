package maingodoc

import (
	"log"
	"os"

	"todo/pkg/db"
	"todo/pkg/server"
)

func main() {
	dbFile := os.Getenv("TODO_DBFILE")
	if dbFile == "" {
		dbFile = "scheduler.db" // путь к бд
	}

	err := db.Init(dbFile) // инициализация бд
	if err != nil {
		log.Fatal("Ошибка инициализации базы данных:", err)
	}

	server.StartServer() // старт сервера
}
