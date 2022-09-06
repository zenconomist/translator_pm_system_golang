package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/handlers"

	envir "environment"

	"github.com/gorilla/mux"
)

type Response struct {
	Message string `json:"message,omitempty"`
}

type ServerHandler struct {
	env envir.Environment
}

func EmptyHttpServer(env envir.Environment) {}

func InitiateServer(env envir.Environment) *http.Server {
	r := InitRouter(env)
	corsObj := handlers.AllowedOrigins([]string{"*"})
	myHandler := handlers.CORS(corsObj)(r)

	srv := &http.Server{
		Handler: myHandler,
		// Addr:    "0.0.0.0:443", // prod ip - dockerben is és windows gépen is a 0.0.0.0-s ip működik local-ként
		Addr: "0.0.0.0:80", // dev ip
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	return srv
}

func InitRouter(env envir.Environment) *mux.Router {
	r := mux.NewRouter()
	ish := ServerHandler{env}
	srGet := r.Methods("GET").Subrouter()

	// get all items
	srGet.HandleFunc("/customer/all", AuthReqAny(ish.GetAllCustomer))
	srGet.HandleFunc("/firm/all", AuthReqAny(ish.GetAllFirm))
	srGet.HandleFunc("/user/all", AuthReqAny(ish.GetAllUser))
	srGet.HandleFunc("/project/all", AuthReqAny(ish.GetAllProject))
	srGet.HandleFunc("/batch/all", AuthReqAny(ish.GetAllBatch))
	srGet.HandleFunc("/task/all", AuthReqAny(ish.GetAllTask))
	srGet.HandleFunc("/upmlogger/all", AuthReqAny(ish.GetAllUPMLogger))
	srGet.HandleFunc("/taskconfig/all", AuthReqAny(ish.GetAllTaskConfig))
	srGet.HandleFunc("/taskoffered/all", AuthReqAny(ish.GetAllTaskOffered))
	srGet.HandleFunc("/clientoffer/all", AuthReqAny(ish.GetAllClientOffer))
	srGet.HandleFunc("/clientoffertask/all", AuthReqAny(ish.GetAllClientOfferTask))

	// get By ID
	srGet.HandleFunc("/customer/{id}", AuthReqAny(ish.GetCustomerByID))
	srGet.HandleFunc("/firm/{id}", AuthReqAny(ish.GetFirmByID))
	srGet.HandleFunc("/user/{id}", AuthReqAny(ish.GetUserByID))
	srGet.HandleFunc("/project/{id}", AuthReqAny(ish.GetProjectByID))
	srGet.HandleFunc("/batch/{id}", AuthReqAny(ish.GetBatchByID))
	srGet.HandleFunc("/task/{id}", AuthReqAny(ish.GetTaskByID))
	srGet.HandleFunc("/upmlogger/{id}", AuthReqAny(ish.GetUPMLoggerByID))
	srGet.HandleFunc("/taskconfig/{id}", AuthReqAny(ish.GetTaskConfigByID))
	srGet.HandleFunc("/taskoffered/{id}", AuthReqAny(ish.GetTaskOfferedByID))
	srGet.HandleFunc("/clientoffer/{id}", AuthReqAny(ish.GetClientOfferByID))
	srGet.HandleFunc("/clientoffertask/{id}", AuthReqAny(ish.GetClientOfferTaskByID))

	// add new items
	srPost := r.Methods("POST").Subrouter()
	srPost.HandleFunc("/customer", AuthReqAny(ish.AddNewCustomer))
	srPost.HandleFunc("/firm", AuthReqAny(ish.AddNewFirm))
	srPost.HandleFunc("/user", AuthReqAny(ish.AddNewUser))
	srPost.HandleFunc("/project", AuthReqAny(ish.AddNewProject))
	srPost.HandleFunc("/batch", AuthReqAny(ish.AddNewBatch))
	srPost.HandleFunc("/task", AuthReqAny(ish.AddNewTask))
	srPost.HandleFunc("/taskconfig", AuthReqAny(ish.AddNewTaskConfig))
	srPost.HandleFunc("/taskoffered", AuthReqAny(ish.AddNewTaskOffered))
	srPost.HandleFunc("/clientoffer", AuthReqAny(ish.AddNewClientOffer))
	srPost.HandleFunc("/clientoffertask", AuthReqAny(ish.AddNewClientOfferTask))

	// update items
	srPut := r.Methods("PUT").Subrouter()
	srPut.HandleFunc("/customer", AuthReqAny(ish.UpdateCustomer))
	srPut.HandleFunc("/firm", AuthReqAny(ish.UpdateFirm))
	srPut.HandleFunc("/user", AuthReqAny(ish.UpdateUser))
	srPut.HandleFunc("/project", AuthReqAny(ish.UpdateProject))
	srPut.HandleFunc("/batch", AuthReqAny(ish.UpdateBatch))
	srPut.HandleFunc("/task", AuthReqAny(ish.UpdateTask))
	srPut.HandleFunc("/taskconfig", AuthReqAny(ish.UpdateTaskConfig))
	srPut.HandleFunc("/taskoffered", AuthReqAny(ish.UpdateTaskOffered))
	srPut.HandleFunc("/taskoffered/state", AuthReqAny(ish.SetTaskOfferedState))
	srPut.HandleFunc("/clientoffer", AuthReqAny(ish.UpdateClientOffer))
	srPut.HandleFunc("/clientoffertask", AuthReqAny(ish.UpdateClientOfferTask))

	// delete items
	srDel := r.Methods("DELETE").Subrouter()
	srDel.HandleFunc("/customer", AuthReqAny(ish.DeleteCustomer))
	srDel.HandleFunc("/firm", AuthReqAny(ish.DeleteFirm))
	srDel.HandleFunc("/user", AuthReqAny(ish.DeleteUser))
	srDel.HandleFunc("/project", AuthReqAny(ish.DeleteProject))
	srDel.HandleFunc("/batch", AuthReqAny(ish.DeleteBatch))
	srDel.HandleFunc("/task", AuthReqAny(ish.DeleteTask))
	srDel.HandleFunc("/taskconfig", AuthReqAny(ish.DeleteTaskConfig))
	srDel.HandleFunc("/taskoffered", AuthReqAny(ish.DeleteTaskOffered))
	srDel.HandleFunc("/clientoffer", AuthReqAny(ish.DeleteClientOffer))
	srDel.HandleFunc("/clientoffertask", AuthReqAny(ish.DeleteClientOfferTask))

	return r

}

func AuthReqAny(handler http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// authentication
		handler.ServeHTTP(w, r)
	}
}

func (sh ServerHandler) recoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered ", r)
	}
}
