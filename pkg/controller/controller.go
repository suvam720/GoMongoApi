package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/suvam720/mongoapi/pkg/dbhelper"
	"github.com/suvam720/mongoapi/pkg/model"
)

func InsertOneTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var todo model.Task
	_ = json.NewDecoder(r.Body).Decode(&todo)
	dbhelper.InsertTask(todo)
	json.NewEncoder(w).Encode(todo)
}

func GetAllTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allTasks := dbhelper.GetTasks()
	json.NewEncoder(w).Encode(allTasks)
}

func MarkCompleted(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	params := mux.Vars(r)
	dbhelper.UpdateTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func DeleteOneTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	dbhelper.DeleteTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := dbhelper.DeleteTasks()
	json.NewEncoder(w).Encode(count)
}
