package application

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Imanr2/Restaurant_API/internal/database"
	"github.com/Imanr2/Restaurant_API/internal/menuitem"
	"github.com/Imanr2/Restaurant_API/internal/order"
	"github.com/Imanr2/Restaurant_API/internal/reservation"
	"github.com/Imanr2/Restaurant_API/internal/session"
	"github.com/Imanr2/Restaurant_API/internal/table"
	"github.com/Imanr2/Restaurant_API/internal/user"
	"github.com/go-playground/validator"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Application struct {
	Router *mux.Router
}

var userManager user.UserManager
var menuItemManager menuitem.MenuItemManager
var orderManager order.OrderManager
var reservationManager reservation.ReservationManager
var tableManager table.TableManager

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
	menuItemManager = menuitem.NewMenuItemManager(db)
	orderManager = order.NewOrderManager(db)
	reservationManager = reservation.NewReservationManager(db)
	tableManager = table.NewTableManager(db)

	app.Router = mux.NewRouter()
	fs := http.FileServer(http.Dir("../images"))
	app.Router.PathPrefix("/image/").Handler(http.StripPrefix(("/image/"), fs))
	app.InitializeRoutes()
}

func (app *Application) InitialMigration(database *gorm.DB) error {
	err := database.AutoMigrate(
		&user.User{},
		&menuitem.MenuItem{},
		&menuitem.Ingredient{},
		&menuitem.Image{},
		&order.Order{},
		&order.OrderItem{},
		&table.Table{},
		&reservation.Reservation{},
	)
	return err
}

func (app *Application) InitializeRoutes() {
	// User routes
	app.Router.HandleFunc("/register", app.Register).Methods("POST")
	app.Router.HandleFunc("/login", app.Login).Methods("POST")

	// Menu Item routes
	app.Router.HandleFunc("/menuitem", app.GetMenuItems).Methods("GET")
	app.Router.HandleFunc("/menuitem/{id}", app.GetMenuItem).Methods("GET")
	app.Router.HandleFunc("/menuitem", app.CreateMenuItem).Methods("POST")
	app.Router.HandleFunc("/menuitem/{id}", app.DeleteMenuItem).Methods("DELETE")
	app.Router.HandleFunc("/menuitem/{id}", app.UpdateMenuItem).Methods("PUT")
	app.Router.HandleFunc("/menuitem/image", app.SaveImage).Methods("POST")

	// Order routes
	app.Router.HandleFunc("/order", app.GetOrders).Methods("GET")
	app.Router.HandleFunc("/order/{id}", app.GetOrderWithID).Methods("GET")
	app.Router.HandleFunc("/order", app.CreateOrder).Methods("POST")
	app.Router.HandleFunc("/order/{id}", app.CompleteOrder).Methods("PUT")
	app.Router.HandleFunc("/order/{id}", app.DeleteOrder).Methods("DELETE")

	// Reservation routes
	app.Router.HandleFunc("/reservation", app.GetReservations).Methods("GET")
	app.Router.HandleFunc("/reservation/{id}", app.GetReservationWithID).Methods("GET")
	app.Router.HandleFunc("/reservation", app.CreateReservation).Methods("POST")
	app.Router.HandleFunc("/reservation/{id}", app.FulfillReservation).Methods("PUT")
	app.Router.HandleFunc("/reservation/{id}", app.DeleteReservation).Methods("DELETE")

	// Table routes
	app.Router.HandleFunc("/table", app.GetTables).Methods("GET")
}

func (app *Application) GetTables(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	_, err := app.authenticate(w, r)
	if err != nil {
		resp := reservation.GetResponse{
			ErrorCode: 2,
			Error:     err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	resp, err := tableManager.GetTables()
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(resp)
	return
}

func (app *Application) GetReservations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	_, err := app.authenticate(w, r)
	if err != nil {
		resp := reservation.GetResponse{
			ErrorCode: 2,
			Error:     err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	resp, err := reservationManager.GetReservations()
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(resp)
	return
}

func (app *Application) GetReservationWithID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	_, err := app.authenticate(w, r)
	if err != nil {
		resp := reservation.GetWithIDResponse{
			ErrorCode: 2,
			Error:     err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	params := mux.Vars(r)
	var getRequest reservation.GetWithIDRequest
	getRequest.ID = params["id"]

	resp, err := reservationManager.GetReservationWithID(getRequest)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(resp)
	return
}

func (app *Application) CreateReservation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	_, err := app.authenticate(w, r)
	if err != nil {
		resp := reservation.CreateResponse{
			ErrorCode: 2,
			Error:     err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	var createRequest reservation.CreateRequest
	json.NewDecoder(r.Body).Decode(&createRequest)

	validate := validator.New()
	err = validate.Struct(createRequest)

	if err != nil {
		resp := reservation.CreateResponse{
			ErrorCode: 1,
			Error:     err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	resp, err := reservationManager.CreateReservation(createRequest)

	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(resp)
	return
}

func (app *Application) FulfillReservation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	_, err := app.authenticate(w, r)
	if err != nil {
		resp := reservation.FulfillResponse{
			ErrorCode: 2,
			Error:     err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	params := mux.Vars(r)
	var completeRequest reservation.FulfillRequest
	completeRequest.ID = params["id"]

	resp, err := reservationManager.FulfillReservation(completeRequest)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(resp)
	return
}

func (app *Application) DeleteReservation(w http.ResponseWriter, r *http.Request) {
	_, err := app.authenticate(w, r)
	if err != nil {
		resp := reservation.DeleteResponse{
			ErrorCode: 2,
			Error:     err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	params := mux.Vars(r)
	var deleteRequest reservation.DeleteRequest
	deleteRequest.ID = params["id"]

	resp, err := reservationManager.DeleteReservation(deleteRequest)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(resp)
	return
}

func (app *Application) GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	_, err := app.authenticate(w, r)
	if err != nil {
		resp := order.GetResponse{
			ErrorCode: 2,
			Error:     err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	resp, err := orderManager.GetOrders()
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(resp)
	return
}

func (app *Application) GetOrderWithID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	_, err := app.authenticate(w, r)
	if err != nil {
		resp := order.GetWithIDResponse{
			ErrorCode: 2,
			Error:     err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	params := mux.Vars(r)
	var getRequest order.GetWithIDRequest
	getRequest.ID = params["id"]

	resp, err := orderManager.GetOrderWithID(getRequest)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(resp)
	return
}

func (app *Application) CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userId, err := app.authenticate(w, r)
	if err != nil {
		resp := order.CreateResponse{
			ErrorCode: 2,
			Error:     err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	var createRequest order.CreateRequest
	json.NewDecoder(r.Body).Decode(&createRequest)

	validate := validator.New()
	err = validate.Struct(createRequest)

	if err != nil {
		resp := order.CreateResponse{
			ErrorCode: 1,
			Error:     err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	createRequest.UserID = userId
	resp, err := orderManager.CreateOrder(createRequest)

	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(resp)
	return
}

func (app *Application) CompleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	_, err := app.authenticate(w, r)
	if err != nil {
		resp := order.CompleteResponse{
			ErrorCode: 2,
			Error:     err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	params := mux.Vars(r)
	var completeRequest order.CompleteRequest
	completeRequest.ID = params["id"]

	resp, err := orderManager.CompleteOrder(completeRequest)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(resp)
	return
}

func (app *Application) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	_, err := app.authenticate(w, r)
	if err != nil {
		resp := order.DeleteResponse{
			ErrorCode: 2,
			Error:     err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	params := mux.Vars(r)
	var deleteRequest order.DeleteRequest
	deleteRequest.ID = params["id"]

	resp, err := orderManager.DeleteOrder(deleteRequest)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(resp)
	return
}

func (app *Application) DeleteMenuItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	_, err := app.authenticate(w, r)
	if err != nil {
		resp := menuitem.DeleteResponse{
			ErrorCode: 2,
			Error:     err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	params := mux.Vars(r)
	var deleteRequest menuitem.DeleteRequest
	deleteRequest.ID = params["id"]

	resp, err := menuItemManager.DeleteMenuItem(deleteRequest)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(resp)
	return
}

func (app *Application) UpdateMenuItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	_, err := app.authenticate(w, r)
	if err != nil {
		resp := menuitem.UpdateResponse{
			ErrorCode: 2,
			Error:     err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	var updateRequest menuitem.UpdateRequest
	json.NewDecoder(r.Body).Decode(&updateRequest)
	params := mux.Vars(r)
	updateRequest.ID = params["id"]

	validate := validator.New()
	err = validate.Struct(updateRequest)

	if err != nil {
		resp := menuitem.UpdateResponse{
			ErrorCode: 1,
			Error:     err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	resp, err := menuItemManager.UpdateMenuItem(updateRequest)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(resp)
	return
}

func (app *Application) GetMenuItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	_, err := app.authenticate(w, r)
	if err != nil {
		resp := menuitem.GetResponse{
			ErrorCode: 2,
			Error:     err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	resp, err := menuItemManager.GetMenuItems()
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(resp)
	return
}

func (app *Application) GetMenuItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	_, err := app.authenticate(w, r)
	if err != nil {
		resp := menuitem.GetWithIDResponse{
			ErrorCode: 2,
			Error:     err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	params := mux.Vars(r)
	var getRequest menuitem.GetWithIDRequest
	getRequest.ID = params["id"]

	resp, err := menuItemManager.GetMenuItemWithID(getRequest)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(resp)
	return
}

func (app *Application) CreateMenuItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userId, err := app.authenticate(w, r)

	if err != nil {
		resp := menuitem.CreateResponse{
			ErrorCode: 2,
			Error:     err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	var createRequest menuitem.CreateRequest
	json.NewDecoder(r.Body).Decode(&createRequest)

	validate := validator.New()
	err = validate.Struct(createRequest)

	if err != nil {
		resp := menuitem.CreateResponse{
			ErrorCode: 1,
			Error:     err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	createRequest.UserID = userId
	resp, err := menuItemManager.CreateMenuItem(createRequest)

	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(resp)
	return
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

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwt.TokenString,
		Expires: jwt.ExpirationTime,
		Path:    "/",
	})
	json.NewEncoder(w).Encode(resp)

	fmt.Println("cookie was set")
	return
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
		fmt.Println(err.Error())
		json.NewEncoder(w).Encode(resp)
		return
	}

	resp := userManager.Register(registerRequest)

	json.NewEncoder(w).Encode(resp)
	return
}

func (app *Application) authenticate(w http.ResponseWriter, r *http.Request) (id uint, err error) {
	tkn, err := session.GetToken(r)
	if err != nil {
		return
	}

	id, err = session.VerifyToken(tkn)
	if err != nil {
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tkn,
		Expires: time.Now().Add(5 * time.Minute),
		Path:    "/",
	})
	return
}

func (app *Application) Run() {
	fmt.Println("application running")
	credentials := handlers.AllowCredentials()
	headers := handlers.AllowedHeaders([]string{"Access-Control-Allow-Headers", "X-Requested-With", "Content-Type", "Authorization", "Accept"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", handlers.CORS(credentials, headers, methods, origins)(app.Router)))
}

func getDB(dbConfig database.DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.GetUsername(), dbConfig.GetPassword(), dbConfig.GetNet(), dbConfig.GetPort(), dbConfig.GetDBName())
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func (app *Application) SaveImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	file, handler, err := r.FormFile("image")

	if err != nil {
		resp := menuitem.SaveImageResponse{
			ErrorCode: 4,
			Error:     err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	defer file.Close()
	f, err := os.OpenFile("../images/"+handler.Filename, os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		resp := menuitem.SaveImageResponse{
			ErrorCode: 4,
			Error:     err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	defer f.Close()
	io.Copy(f, file)

	resp := menuitem.SaveImageResponse{
		ImageUrl: fmt.Sprintf("http://localhost:8000/image/%s", handler.Filename),
	}
	json.NewEncoder(w).Encode(resp)
	return
}
