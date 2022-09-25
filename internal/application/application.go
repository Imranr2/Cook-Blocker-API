package application

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Imanr2/Restaurant_API/internal/database"
	"github.com/Imanr2/Restaurant_API/internal/user"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Application struct {
	Router *mux.Router
}

var userManager user.UserManager

func (app *Application) Initialize(dbConfig database.DBConfig) {
	db, err := getDB(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	err = app.InitialMigration(db)
	if err != nil {
		log.Fatal(err)
	}

	userManager = user.NewUserManager(db)

	app.Router = mux.NewRouter()

	app.InitializeRoutes()
}

func (app *Application) InitialMigration(database *gorm.DB) error {
	err := database.AutoMigrate(&user.User{})
	return err
}

func (app *Application) InitializeRoutes() {
	app.Router.HandleFunc("/register", app.Register).Methods("POST")
	app.Router.HandleFunc("/login", app.Login).Methods("POST")
}

func (app *Application) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var loginRequest user.LoginRequest
	json.NewDecoder(r.Body).Decode(&loginRequest)

	validate := validator.New()
	err := validate.Struct(loginRequest)

	if err != nil {
		resp := user.LoginResponse{
			ErrorCode: 1,
			Error:     err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	jwt, resp := userManager.Login(loginRequest)

	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode(resp)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwt.TokenString,
		Expires: jwt.ExpirationTime,
	})
}

func (app *Application) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var registerRequest user.RegisterRequest
	json.NewDecoder(r.Body).Decode(&registerRequest)

	validate := validator.New()
	err := validate.Struct(registerRequest)

	if err != nil {
		resp := user.RegisterResponse{
			ErrorCode: 1,
			Error:     err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	resp := userManager.Register(registerRequest)

	json.NewEncoder(w).Encode(resp)
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
