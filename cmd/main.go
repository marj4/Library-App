// There code be run
package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"library/config"
	"library/internal/repository"
	"library/internal/service"
	"log"
)

func main() {
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

	books, err := srv.GetAllInfoForBook()
	if err != nil {
		log.Fatalf("Cant get info for book:%v", err)
	}
	fmt.Println(books)

	//a := app2.New()
	//w := a.NewWindow("Библиотека")
	//bookList := widget.NewList(
	//	func() int {
	//		books, _ := srv.GetAllInfoForBook()
	//		return len(books)
	//	},
	//	func() fyne.CanvasObject { return widget.NewLabel("") },
	//	func(i widget.ListItemID, obj fyne.CanvasObject) {
	//		//books,_:= srv.GetAllInfoForBook()
	//		obj.(*widget.Label).SetText("book1")
	//	})
	//w.SetContent(container.NewVBox(
	//	widget.NewLabel("Список книг:"),
	//	bookList,
	//	widget.NewButton("Добавить книгу", func() {
	//		// Здесь должен быть вызов метода backend для добавления книги
	//	}),
	//	widget.NewButton("Удалить книгу", func() {
	//		// Здесь должен быть вызов метода backend для удаления книги
	//	})))
	//w.ShowAndRun()
}

func ConnectString(cfg *config.Config) string {
	return "postgres://" + cfg.DB_Username + ":" + cfg.DB_Password + "@" + cfg.DB_Host + ":" + cfg.DB_Port + "/" + cfg.DB_Name + "?sslmode=" + cfg.DB_SSlMODE
}
