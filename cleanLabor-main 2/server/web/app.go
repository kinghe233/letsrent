package web

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"my-app/db"
	"my-app/model"
	"net/http"
	"time"
)

type App struct {
	d        db.DB
	handlers map[string]http.HandlerFunc
}

type Taskers []model.Tasker

func allTaskers(w http.ResponseWriter, r *http.Request) {
	taskers := Taskers{
		model.Tasker{UserName: "Ray", Desc: "Test", WorkContent: "Hello", Phone: "111-111-1111"},
		model.Tasker{UserName: "Robin", Desc: "Test", WorkContent: "Hello", Phone: "111-111-1111"},
	}

	fmt.Println("Endpoint Hit: All Taskers EndPoint")
	json.NewEncoder(w).Encode(taskers)
}

func allTaskCreators(w http.ResponseWriter, r *http.Request) {
	taskers := Taskers{
		model.Tasker{UserName: "Jack", Desc: "Test", WorkContent: "Hello", Phone: "111-111-1111"},
		model.Tasker{UserName: "Lucy", Desc: "Test", WorkContent: "Hello", Phone: "111-111-1111"},
	}

	fmt.Println("Endpoint Hit: All allTaskCreators EndPoint")
	json.NewEncoder(w).Encode(taskers)
}

func collectionExists(client *mongo.Client, database, collection string) bool {
	collections, err := client.Database(database).ListCollectionNames(context.Background(), nil)
	if err != nil {
		return false
	}
	for _, coll := range collections {
		if coll == collection {
			return true
		}
	}
	return false
}

// Return JSON struct [Yeshan Li]
type retResult struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
}

func registerTaskCreatorsUser(w http.ResponseWriter, r *http.Request) {
	// 解析请求体中的JSON数据并将其存储在User结构体中
	fmt.Println("Endpoint Hit: registerTaskCreatorsUser")
	taskMan := model.TaskCreator{
		UserName: "Ray", Desc: "Test", WorkContent: "Hello", Phone: "111-111-1111",
	}
	fmt.Println(taskMan)
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	fmt.Println(client, "Client")
	if err != nil {
		fmt.Println(err, "after mongo.NewClient")
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err, "error after client.Connect")
		fmt.Println(err)
		log.Fatal(err)
	}

	// err = client.Database("mydb").CreateCollection(context.Background(), "mycollection", nil)
	// if err != nil {
	// 	fmt.Println("Failed to create collection:", err)
	// 	return
	// }

	// List database names
	databases, err := client.ListDatabaseNames(context.Background(), nil)
	if err != nil {
		fmt.Println("here error list datebase")
		log.Fatal(err)
	}

	// Print database names
	for _, db := range databases {
		fmt.Println(db, "list datebase")
	}

	collections, err := client.Database("mydb").ListCollectionNames(ctx, nil)
	fmt.Println(collections, "after client.Database.ListCollectionNames(ctx, nil) collections")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "TaskCreator registered successfully!")
}

func (a *App) registerTaskerUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	// 解析请求体中的JSON数据并将其存储在User结构体中
	var taskMan model.Tasker
	err := json.NewDecoder(r.Body).Decode(&taskMan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: 在数据库中创建新用户
	// Initiate response data [Yeshan Li]
	w.Header().Set("WorkContent-Type", "application/json")
	var ret retResult
	// Check whether the inputted user information valid or not [Yeshan Li]
	var validateResult = a.ValidateTaskerRegister(taskMan)
	if validateResult != "success" {
		ret.Status = "fail"
		ret.Reason = validateResult
		json.NewEncoder(w).Encode(ret)
		return
	}
	// Add the new user account to the Database [Yeshan Li]
	a.AddTasker(taskMan)
	ret.Status = "success"
	json.NewEncoder(w).Encode(ret)

	// 返回成功响应
	w.WriteHeader(http.StatusCreated)
}

// Login function [Yeshan Li]
func (a *App) loginTaskerUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	// 解析请求体中的JSON数据并将其存储在User结构体中
	var taskMan model.Tasker
	err := json.NewDecoder(r.Body).Decode(&taskMan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Initiate response data [Yeshan Li]
	w.Header().Set("WorkContent-Type", "application/json")
	var ret retResult
	// Check whether the inputted user information valid or not [Yeshan Li]
	var validateResult = a.ValidateTaskerLogin(taskMan)
	if validateResult != "success" {
		ret.Status = "fail"
		ret.Reason = validateResult
		json.NewEncoder(w).Encode(ret)
		return
	}
	ret.Status = "success"
	json.NewEncoder(w).Encode(ret)
	w.WriteHeader(http.StatusCreated)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func NewApp(d db.DB, cors bool) App {
	app := App{
		d:        d,
		handlers: make(map[string]http.HandlerFunc),
	}
	techHandler := app.GetTechnologies
	if !cors {
		techHandler = disableCors(techHandler)
	}
	taskerHandler := app.GetallTaskers
	if !cors {
		taskerHandler = disableCors(taskerHandler)
	}
	app.handlers["/api/technologies"] = techHandler
	app.handlers["/"] = http.FileServer(http.Dir("/webapp")).ServeHTTP
	app.handlers["/allTaskers"] = allTaskers
	// API to display all the Taskers in the Database (For test use) [Yeshan Li]
	app.handlers["/GetallTaskers"] = taskerHandler
	app.handlers["/allTaskCreators"] = allTaskCreators
	// API to register the account [Yeshan Li]
	app.handlers["/registerTaskerUser"] = app.registerTaskerUser
	// API to login the account [Yeshan Li]
	app.handlers["/loginTaskerUser"] = app.loginTaskerUser
	app.handlers["/registerTaskCreator"] = registerTaskCreatorsUser
	return app
}

func (a *App) Serve() error {
	for path, handler := range a.handlers {
		http.Handle(path, handler)
	}
	log.Println("Web server is available on port 8080")
	return http.ListenAndServe(":8080", nil)
}

func (a *App) GetTechnologies(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("WorkContent-Type", "application/json")
	technologies, err := a.d.GetTechnologies()
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(technologies)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

// Function to display all the Taskers in the Database (For test use) [Yeshan Li]
func (a *App) GetallTaskers(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("WorkContent-Type", "application/json")
	taskers, err := a.d.GetTaskers()
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(taskers)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
	fmt.Println("Endpoint Hit: All Taskers EndPoint")
}

// Function to validate the new user information in the Database (For test use) [Yeshan Li]
func (a *App) ValidateTaskerRegister(tasker model.Tasker) string {
	if tasker.UserName == "" {
		fmt.Println("Username is empty!")
		return "Username is empty."
	} else if tasker.Desc == "" {
		fmt.Println("Desc is empty!")
		return "Desc is empty."
	} else if tasker.WorkContent == "" {
		fmt.Println("WorkContent is empty!")
		return "WorkContent is empty."
	} else if tasker.Email == "" {
		fmt.Println("Email is empty!")
		return "Email is empty."
	} else if tasker.Phone == "" {
		fmt.Println("Phone is empty!")
		return "Phone is empty."
	} else if tasker.Password == "" {
		fmt.Println("Password is empty!")
		return "Password is empty."
	}
	var tsk []*model.Tasker
	tsk, _ = a.d.GetTaskers()
	fmt.Println("Endpoint Hit: Validate Tasker Register EndPoint")
	for i := 0; i < len(tsk); i++ {
		if tsk[i].UserName == tasker.UserName {
			fmt.Println("Username exist!")
			return "Username existed."
		}
		if tsk[i].Email == tasker.Email {
			fmt.Println("Email exist!")
			return "Email existed."
		}
		if tsk[i].Phone == tasker.Phone {
			fmt.Println("Phone number exist!")
			return "Phone number existed."
		}
	}
	return "success"
}

// Function to validate the login user information in the Database (For test use) [Yeshan Li]
func (a *App) ValidateTaskerLogin(tasker model.Tasker) string {
	if tasker.UserName == "" {
		fmt.Println("Username is empty!")
		return "Username is empty."
	} else if tasker.Password == "" {
		fmt.Println("Password is empty!")
		return "Password is empty."
	}
	var tsk []*model.Tasker
	tsk, _ = a.d.GetTaskers()
	fmt.Println("Endpoint Hit: Validate Tasker Login EndPoint")
	for i := 0; i < len(tsk); i++ {
		if tsk[i].UserName == tasker.UserName {
			if tsk[i].Password == tasker.Password {
				fmt.Println("Login Successfully!")
				return "success"
			} else {
				fmt.Println("Incorrect Password!")
				return "Incorrect Password."
			}
		}
	}
	return "Username not found in the system, please check your input."
}

// Function to add the new user information in the Database (For test use) [Yeshan Li]
func (a *App) AddTasker(tasker model.Tasker) {
	a.d.AddTasker(tasker)
	fmt.Println("Endpoint Hit: Add Taskers EndPoint")
}

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}

// Needed in order to disable CORS for local development
func disableCors(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		h(w, r)
	}
}
