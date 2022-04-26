package router

import (
	"github.com/gorilla/mux"
	"github.com/suvam720/mongoapi/pkg/controller"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/task", controller.InsertOneTask).Methods("POST")
	router.HandleFunc("/api/tasks", controller.GetAllTask).Methods("GET")
	router.HandleFunc("/api/task/{id}", controller.MarkCompleted).Methods("PUT")
	router.HandleFunc("/api/task/{id}", controller.DeleteOneTask).Methods("DELETE")
	router.HandleFunc("/api/deletetasks", controller.DeleteAllTask).Methods("DELETE")

	return router
}
