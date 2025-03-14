// Здесь будут импортироваться данные из файла .env для подключения БД и прочее
package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DB_Username string
	DB_Password string
	DB_Host     string
	DB_Port     string
	DB_Name     string
	DB_SSlMODE  string
}

// Загрузка конфига
func LoadConfig() (*Config, error) {

	//Импорт данных из файла .env
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	//Возвращаем структуру в которую мы подставили данные из файла .env
	return &Config{
		DB_Username: os.Getenv("DB_USERNAME"),
		DB_Password: os.Getenv("DB_PASSWORD"),
		DB_Host:     os.Getenv("DB_HOST"),
		DB_Port:     os.Getenv("DB_PORT"),
		DB_Name:     os.Getenv("DB_NAME"),
		DB_SSlMODE:  os.Getenv("DB_SSLMODE"),
	}, nil
}

// Создание строки подключения
func ConnectString(cfg *Config) string {
	return "postgres://" + cfg.DB_Username + ":" + cfg.DB_Password + "@" + cfg.DB_Host + ":" + cfg.DB_Port + "/" + cfg.DB_Name + "?sslmode=" + cfg.DB_SSlMODE
}
