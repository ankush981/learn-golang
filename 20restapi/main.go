package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Todo struct {
	Id     string  `json:"id"`
	Text   string  `json:"text"`
	Author *Author `json:"author"`
}

type Author struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

func (t *Todo) IsEmptyTodo() bool {
	return t.Text == ""
}

// storage (database) for the todos
var todos []Todo

func main() {
	r := mux.NewRouter()
	seed()

	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/todos", getAllTodos).Methods("GET")
	r.HandleFunc("/todos/{id}", getTodo).Methods("GET")
	r.HandleFunc("/todos", createTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", updateTodo).Methods("PATCH")
	r.HandleFunc("/todos/{id}", deleteTodo).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func seed() {
	todos = append(todos, Todo{Id: "1", Text: "Walk the dog", Author: &Author{Name: "Ankush Thakur", Location: "Delhi"}})
	todos = append(todos, Todo{Id: "2", Text: "Read the book", Author: &Author{Name: "Ankush", Location: "Remote"}})
}

// controllers
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Todo Maker</h1>"))
}

func getAllTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, todo := range todos {
		if todo.Id == params["id"] {
			json.NewEncoder(w).Encode(todo)
			return
		}
	}

	json.NewEncoder(w).Encode("No todo found")
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("No request body")
	}

	var todo Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	if todo.IsEmptyTodo() {
		json.NewEncoder(w).Encode("No data inside the request")
		return
	}

	// generarte "unique" id
	rand.Seed(time.Now().UnixNano())
	todo.Id = strconv.Itoa(rand.Intn(100))
	todos = append(todos, todo)
	json.NewEncoder(w).Encode(todo)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parmas := mux.Vars(r)

	// remove by id and add a new (updated) one
	for index, todo := range todos {
		if todo.Id == parmas["id"] {
			todos = append(todos[:index], todos[index+1:]...)
			var t Todo
			_ = json.NewDecoder(r.Body).Decode(&t)
			t.Id = parmas["id"]
			todos = append(todos, t)
			json.NewEncoder(w).Encode(t)
			return
		}
	}
	json.NewEncoder(w).Encode("No matching Todo found!")
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, todo := range todos {
		if todo.Id == params["id"] {
			todos = append(todos[:index], todos[index+1:]...)
			json.NewEncoder(w).Encode("Deleted")
			return
		}
	}
}
