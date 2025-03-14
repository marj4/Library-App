package models

// Модель таблицы Book из БД library
type Book struct {
	Id int `json:"id"`
	//Id_inventory string `json:"id_inventory"`
	BookName     string `json:"bookName"`
	Author       string `json:"author"`
	Genre        string `json:"genre"`
	YearRealease int    `json:"yearReale"`
}
