package application

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Imanr2/Restaurant_API/internal/database"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Application struct {
	Router *mux.Router
}

func (app *Application) Initialize(dbConfig database.DBConfig) {
	_, err := getDB(dbConfig)

	if err != nil {
		log.Fatal(err)
	}

	app.Router = mux.NewRouter()

	app.InitializeRoutes()
}

func (app *Application) InitializeRoutes() {

}

func (app *Application) Run() {
	fmt.Println("application running")
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", app.Router))
}

func getDB(dbConfig database.DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.GetUsername(), dbConfig.GetPassword(), dbConfig.GetNet(), dbConfig.GetPort(), dbConfig.GetDBName())
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
