// There code be run
package main

import (
	app2 "fyne.io/fyne/v2/app"
	"library/config"
	"library/internal/repository"
	"library/internal/service"
	"library/internal/ui"
	"log"
)

func main() {
	a := app2.New()
	w := a.NewWindow("Library App")

	//Run the config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Cant load config file.Error:%v", err)
	}

	//Запускаю БД
	db, err := repository.Connect(ConnectString(cfg))
	if err != nil {
		log.Fatalf("Cant connect to database:%v", err)
	}
	defer db.Close()

	//Инициализация репозитория
	repo := repository.NewRepository(db)

	//Инициализация сервиса
	srv := service.NewService(repo)

	//Инициализация графического интерфейса
	uiManager := ui.NewUIManager(w)

}

func ConnectString(cfg *config.Config) string {
	return "postgres://" + cfg.DB_Username + ":" + cfg.DB_Password + "@" + cfg.DB_Host + ":" + cfg.DB_Port + "/" + cfg.DB_Name + "?sslmode=" + cfg.DB_SSlMODE
}
