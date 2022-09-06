package server

import (
	"dto"
	"encoding/json"
	"entities"
	"fmt"
	"net/http"
	"services"
	"strconv"

	"github.com/gorilla/mux"
)

func (sh ServerHandler) AddNewFirm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add firm init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "new firm adding process started",
	}

	entity := entities.Firm{}
	history := entities.FirmHistory{}
	dto := dto.FirmRequestDTO{}
	errJsonDecode := json.NewDecoder(r.Body).Decode(&dto)
	if errJsonDecode != nil {
		logger.LogError(errJsonDecode)
		res.Message = fmt.Sprintf("json decoding error: %v", errJsonDecode)
		json.NewEncoder(w).Encode(res)
		return
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entity), sh.env, &dto, entity, &history)
	id, errService := service.CreateNewItem(&dto)
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("New firm created with ID: %d", id)
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) UpdateFirm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add firm init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "firm updating process started",
	}

	entity := entities.Firm{}
	history := entities.FirmHistory{}
	dto := dto.FirmRequestDTO{}
	errJsonDecode := json.NewDecoder(r.Body).Decode(&dto)
	if errJsonDecode != nil {
		logger.LogError(errJsonDecode)
		res.Message = fmt.Sprintf("json decoding error: %v", errJsonDecode)
		json.NewEncoder(w).Encode(res)
		return
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entity), sh.env, &dto, entity, &history)
	errService := service.UpdateItem(&dto)
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("Firm updated.")
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) DeleteFirm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                             // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "delete firm init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "firm deleting process started",
	}

	vars := mux.Vars(r)
	id, _ := vars["id"]
	idInt, errConv := strconv.Atoi(id)
	if errConv != nil {
		defer logger.LogError(errConv)
		res.Message = fmt.Sprintf("couldn't convert given id, error: %v", errConv)
		json.NewEncoder(w).Encode(res)
		return
	}
	Id := uint(idInt)

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.Firm{}), sh.env, &dto.FirmRequestDTO{}, entities.Firm{}, &entities.FirmHistory{})
	errService := service.DeleteItem(Id)
	if errService != nil {
		defer logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("Firm deleted.")
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) AddNewCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                              // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add customer init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "new customer adding process started",
	}

	entity := entities.Customer{}
	history := entities.CustomerHistory{}
	dto := dto.CustomerRequestDTO{}
	errJsonDecode := json.NewDecoder(r.Body).Decode(&dto)
	if errJsonDecode != nil {
		logger.LogError(errJsonDecode)
		res.Message = fmt.Sprintf("json decoding error: %v", errJsonDecode)
		json.NewEncoder(w).Encode(res)
		return
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entity), sh.env, &dto, entity, &history)
	id, errService := service.CreateNewItem(&dto)
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("New customer created with ID: %d", id)
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                              // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add customer init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "customer updating process started",
	}

	entity := entities.Customer{}
	history := entities.CustomerHistory{}
	dto := dto.CustomerRequestDTO{}
	errJsonDecode := json.NewDecoder(r.Body).Decode(&dto)
	if errJsonDecode != nil {
		logger.LogError(errJsonDecode)
		res.Message = fmt.Sprintf("json decoding error: %v", errJsonDecode)
		json.NewEncoder(w).Encode(res)
		return
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entity), sh.env, &dto, entity, &history)
	errService := service.UpdateItem(&dto)
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("Customer updated.")
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                                 // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "delete customer init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "customer deleting process started",
	}

	vars := mux.Vars(r)
	id, _ := vars["id"]
	idInt, errConv := strconv.Atoi(id)
	if errConv != nil {
		defer logger.LogError(errConv)
		res.Message = fmt.Sprintf("couldn't convert given id, error: %v", errConv)
		json.NewEncoder(w).Encode(res)
		return
	}
	Id := uint(idInt)

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.Customer{}), sh.env, &dto.CustomerRequestDTO{}, entities.Customer{}, &entities.CustomerHistory{})
	errService := service.DeleteItem(Id)
	if errService != nil {
		defer logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("Customer deleted.")
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) AddNewDefaultSupplierPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add defaultsupplierprice init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "new defaultsupplierprice adding process started",
	}

	entity := entities.DefaultSupplierPrice{}
	history := entities.DefaultSupplierPriceHistory{}
	dto := dto.DefaultSupplierPriceRequestDTO{}
	errJsonDecode := json.NewDecoder(r.Body).Decode(&dto)
	if errJsonDecode != nil {
		logger.LogError(errJsonDecode)
		res.Message = fmt.Sprintf("json decoding error: %v", errJsonDecode)
		json.NewEncoder(w).Encode(res)
		return
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entity), sh.env, &dto, entity, &history)
	id, errService := service.CreateNewItem(&dto)
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("New defaultsupplierprice created with ID: %d", id)
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) UpdateDefaultSupplierPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add defaultsupplierprice init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "defaultsupplierprice updating process started",
	}

	entity := entities.DefaultSupplierPrice{}
	history := entities.DefaultSupplierPriceHistory{}
	dto := dto.DefaultSupplierPriceRequestDTO{}
	errJsonDecode := json.NewDecoder(r.Body).Decode(&dto)
	if errJsonDecode != nil {
		logger.LogError(errJsonDecode)
		res.Message = fmt.Sprintf("json decoding error: %v", errJsonDecode)
		json.NewEncoder(w).Encode(res)
		return
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entity), sh.env, &dto, entity, &history)
	errService := service.UpdateItem(&dto)
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("DefaultSupplierPrice updated.")
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) DeleteDefaultSupplierPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                                             // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "delete defaultsupplierprice init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "defaultsupplierprice deleting process started",
	}

	vars := mux.Vars(r)
	id, _ := vars["id"]
	idInt, errConv := strconv.Atoi(id)
	if errConv != nil {
		defer logger.LogError(errConv)
		res.Message = fmt.Sprintf("couldn't convert given id, error: %v", errConv)
		json.NewEncoder(w).Encode(res)
		return
	}
	Id := uint(idInt)

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.DefaultSupplierPrice{}), sh.env, &dto.DefaultSupplierPriceRequestDTO{}, entities.DefaultSupplierPrice{}, &entities.DefaultSupplierPriceHistory{})
	errService := service.DeleteItem(Id)
	if errService != nil {
		defer logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("DefaultSupplierPrice deleted.")
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) AddNewTaskStateChangeComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                                            // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add taskstatechangecomment init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "new taskstatechangecomment adding process started",
	}

	entity := entities.TaskStateChangeComment{}
	history := entities.TaskStateChangeCommentHistory{}
	dto := dto.TaskStateChangeCommentRequestDTO{}
	errJsonDecode := json.NewDecoder(r.Body).Decode(&dto)
	if errJsonDecode != nil {
		logger.LogError(errJsonDecode)
		res.Message = fmt.Sprintf("json decoding error: %v", errJsonDecode)
		json.NewEncoder(w).Encode(res)
		return
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entity), sh.env, &dto, entity, &history)
	id, errService := service.CreateNewItem(&dto)
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("New taskstatechangecomment created with ID: %d", id)
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) UpdateTaskStateChangeComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                                            // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add taskstatechangecomment init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "taskstatechangecomment updating process started",
	}

	entity := entities.TaskStateChangeComment{}
	history := entities.TaskStateChangeCommentHistory{}
	dto := dto.TaskStateChangeCommentRequestDTO{}
	errJsonDecode := json.NewDecoder(r.Body).Decode(&dto)
	if errJsonDecode != nil {
		logger.LogError(errJsonDecode)
		res.Message = fmt.Sprintf("json decoding error: %v", errJsonDecode)
		json.NewEncoder(w).Encode(res)
		return
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entity), sh.env, &dto, entity, &history)
	errService := service.UpdateItem(&dto)
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("TaskStateChangeComment updated.")
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) DeleteTaskStateChangeComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                                               // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "delete taskstatechangecomment init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "taskstatechangecomment deleting process started",
	}

	vars := mux.Vars(r)
	id, _ := vars["id"]
	idInt, errConv := strconv.Atoi(id)
	if errConv != nil {
		defer logger.LogError(errConv)
		res.Message = fmt.Sprintf("couldn't convert given id, error: %v", errConv)
		json.NewEncoder(w).Encode(res)
		return
	}
	Id := uint(idInt)

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.TaskStateChangeComment{}), sh.env, &dto.TaskStateChangeCommentRequestDTO{}, entities.TaskStateChangeComment{}, &entities.TaskStateChangeCommentHistory{})
	errService := service.DeleteItem(Id)
	if errService != nil {
		defer logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("TaskStateChangeComment deleted.")
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) AddNewDefaultCustomerPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add defaultcustomerprice init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "new defaultcustomerprice adding process started",
	}

	entity := entities.DefaultCustomerPrice{}
	history := entities.DefaultCustomerPriceHistory{}
	dto := dto.DefaultCustomerPriceRequestDTO{}
	errJsonDecode := json.NewDecoder(r.Body).Decode(&dto)
	if errJsonDecode != nil {
		logger.LogError(errJsonDecode)
		res.Message = fmt.Sprintf("json decoding error: %v", errJsonDecode)
		json.NewEncoder(w).Encode(res)
		return
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entity), sh.env, &dto, entity, &history)
	id, errService := service.CreateNewItem(&dto)
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("New defaultcustomerprice created with ID: %d", id)
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) UpdateDefaultCustomerPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add defaultcustomerprice init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "defaultcustomerprice updating process started",
	}

	entity := entities.DefaultCustomerPrice{}
	history := entities.DefaultCustomerPriceHistory{}
	dto := dto.DefaultCustomerPriceRequestDTO{}
	errJsonDecode := json.NewDecoder(r.Body).Decode(&dto)
	if errJsonDecode != nil {
		logger.LogError(errJsonDecode)
		res.Message = fmt.Sprintf("json decoding error: %v", errJsonDecode)
		json.NewEncoder(w).Encode(res)
		return
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entity), sh.env, &dto, entity, &history)
	errService := service.UpdateItem(&dto)
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("DefaultCustomerPrice updated.")
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) DeleteDefaultCustomerPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                                             // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "delete defaultcustomerprice init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "defaultcustomerprice deleting process started",
	}

	vars := mux.Vars(r)
	id, _ := vars["id"]
	idInt, errConv := strconv.Atoi(id)
	if errConv != nil {
		defer logger.LogError(errConv)
		res.Message = fmt.Sprintf("couldn't convert given id, error: %v", errConv)
		json.NewEncoder(w).Encode(res)
		return
	}
	Id := uint(idInt)

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.DefaultCustomerPrice{}), sh.env, &dto.DefaultCustomerPriceRequestDTO{}, entities.DefaultCustomerPrice{}, &entities.DefaultCustomerPriceHistory{})
	errService := service.DeleteItem(Id)
	if errService != nil {
		defer logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("DefaultCustomerPrice deleted.")
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) AddNewTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add task init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "new task adding process started",
	}

	entity := entities.Task{}
	history := entities.TaskHistory{}
	dto := dto.TaskRequestDTO{}
	errJsonDecode := json.NewDecoder(r.Body).Decode(&dto)
	if errJsonDecode != nil {
		logger.LogError(errJsonDecode)
		res.Message = fmt.Sprintf("json decoding error: %v", errJsonDecode)
		json.NewEncoder(w).Encode(res)
		return
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entity), sh.env, &dto, entity, &history)
	id, errService := service.CreateNewItem(&dto)
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("New task created with ID: %d", id)
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add task init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "task updating process started",
	}

	entity := entities.Task{}
	history := entities.TaskHistory{}
	dto := dto.TaskRequestDTO{}
	errJsonDecode := json.NewDecoder(r.Body).Decode(&dto)
	if errJsonDecode != nil {
		logger.LogError(errJsonDecode)
		res.Message = fmt.Sprintf("json decoding error: %v", errJsonDecode)
		json.NewEncoder(w).Encode(res)
		return
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entity), sh.env, &dto, entity, &history)
	errService := service.UpdateItem(&dto)
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("Task updated.")
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                             // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "delete task init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "task deleting process started",
	}

	vars := mux.Vars(r)
	id, _ := vars["id"]
	idInt, errConv := strconv.Atoi(id)
	if errConv != nil {
		defer logger.LogError(errConv)
		res.Message = fmt.Sprintf("couldn't convert given id, error: %v", errConv)
		json.NewEncoder(w).Encode(res)
		return
	}
	Id := uint(idInt)

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.Task{}), sh.env, &dto.TaskRequestDTO{}, entities.Task{}, &entities.TaskHistory{})
	errService := service.DeleteItem(Id)
	if errService != nil {
		defer logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("Task deleted.")
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) AddNewClientOfferTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                                     // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add clientoffertask init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "new clientoffertask adding process started",
	}

	entity := entities.ClientOfferTask{}
	history := entities.ClientOfferTaskHistory{}
	dto := dto.ClientOfferTaskRequestDTO{}
	errJsonDecode := json.NewDecoder(r.Body).Decode(&dto)
	if errJsonDecode != nil {
		logger.LogError(errJsonDecode)
		res.Message = fmt.Sprintf("json decoding error: %v", errJsonDecode)
		json.NewEncoder(w).Encode(res)
		return
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entity), sh.env, &dto, entity, &history)
	id, errService := service.CreateNewItem(&dto)
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("New clientoffertask created with ID: %d", id)
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) UpdateClientOfferTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                                     // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add clientoffertask init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "clientoffertask updating process started",
	}

	entity := entities.ClientOfferTask{}
	history := entities.ClientOfferTaskHistory{}
	dto := dto.ClientOfferTaskRequestDTO{}
	errJsonDecode := json.NewDecoder(r.Body).Decode(&dto)
	if errJsonDecode != nil {
		logger.LogError(errJsonDecode)
		res.Message = fmt.Sprintf("json decoding error: %v", errJsonDecode)
		json.NewEncoder(w).Encode(res)
		return
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entity), sh.env, &dto, entity, &history)
	errService := service.UpdateItem(&dto)
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("ClientOfferTask updated.")
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) DeleteClientOfferTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                                        // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "delete clientoffertask init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "clientoffertask deleting process started",
	}

	vars := mux.Vars(r)
	id, _ := vars["id"]
	idInt, errConv := strconv.Atoi(id)
	if errConv != nil {
		defer logger.LogError(errConv)
		res.Message = fmt.Sprintf("couldn't convert given id, error: %v", errConv)
		json.NewEncoder(w).Encode(res)
		return
	}
	Id := uint(idInt)

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.ClientOfferTask{}), sh.env, &dto.ClientOfferTaskRequestDTO{}, entities.ClientOfferTask{}, &entities.ClientOfferTaskHistory{})
	errService := service.DeleteItem(Id)
	if errService != nil {
		defer logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("ClientOfferTask deleted.")
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) AddNewProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                             // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add project init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "new project adding process started",
	}

	entity := entities.Project{}
	history := entities.ProjectHistory{}
	dto := dto.ProjectRequestDTO{}
	errJsonDecode := json.NewDecoder(r.Body).Decode(&dto)
	if errJsonDecode != nil {
		logger.LogError(errJsonDecode)
		res.Message = fmt.Sprintf("json decoding error: %v", errJsonDecode)
		json.NewEncoder(w).Encode(res)
		return
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entity), sh.env, &dto, entity, &history)
	id, errService := service.CreateNewItem(&dto)
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("New project created with ID: %d", id)
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) UpdateProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                             // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add project init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "project updating process started",
	}

	entity := entities.Project{}
	history := entities.ProjectHistory{}
	dto := dto.ProjectRequestDTO{}
	errJsonDecode := json.NewDecoder(r.Body).Decode(&dto)
	if errJsonDecode != nil {
		logger.LogError(errJsonDecode)
		res.Message = fmt.Sprintf("json decoding error: %v", errJsonDecode)
		json.NewEncoder(w).Encode(res)
		return
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entity), sh.env, &dto, entity, &history)
	errService := service.UpdateItem(&dto)
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("Project updated.")
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) DeleteProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                                // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "delete project init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "project deleting process started",
	}

	vars := mux.Vars(r)
	id, _ := vars["id"]
	idInt, errConv := strconv.Atoi(id)
	if errConv != nil {
		defer logger.LogError(errConv)
		res.Message = fmt.Sprintf("couldn't convert given id, error: %v", errConv)
		json.NewEncoder(w).Encode(res)
		return
	}
	Id := uint(idInt)

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.Project{}), sh.env, &dto.ProjectRequestDTO{}, entities.Project{}, &entities.ProjectHistory{})
	errService := service.DeleteItem(Id)
	if errService != nil {
		defer logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("Project deleted.")
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) AddNewContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                             // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add contact init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "new contact adding process started",
	}

	entity := entities.Contact{}
	history := entities.ContactHistory{}
	dto := dto.ContactRequestDTO{}
	errJsonDecode := json.NewDecoder(r.Body).Decode(&dto)
	if errJsonDecode != nil {
		logger.LogError(errJsonDecode)
		res.Message = fmt.Sprintf("json decoding error: %v", errJsonDecode)
		json.NewEncoder(w).Encode(res)
		return
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entity), sh.env, &dto, entity, &history)
	id, errService := service.CreateNewItem(&dto)
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("New contact created with ID: %d", id)
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) UpdateContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                             // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add contact init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "contact updating process started",
	}

	entity := entities.Contact{}
	history := entities.ContactHistory{}
	dto := dto.ContactRequestDTO{}
	errJsonDecode := json.NewDecoder(r.Body).Decode(&dto)
	if errJsonDecode != nil {
		logger.LogError(errJsonDecode)
		res.Message = fmt.Sprintf("json decoding error: %v", errJsonDecode)
		json.NewEncoder(w).Encode(res)
		return
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entity), sh.env, &dto, entity, &history)
	errService := service.UpdateItem(&dto)
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("Contact updated.")
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) DeleteContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                                // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "delete contact init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "contact deleting process started",
	}

	vars := mux.Vars(r)
	id, _ := vars["id"]
	idInt, errConv := strconv.Atoi(id)
	if errConv != nil {
		defer logger.LogError(errConv)
		res.Message = fmt.Sprintf("couldn't convert given id, error: %v", errConv)
		json.NewEncoder(w).Encode(res)
		return
	}
	Id := uint(idInt)

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.Contact{}), sh.env, &dto.ContactRequestDTO{}, entities.Contact{}, &entities.ContactHistory{})
	errService := service.DeleteItem(Id)
	if errService != nil {
		defer logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("Contact deleted.")
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) AddNewSupplierPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                                   // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add supplierprice init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "new supplierprice adding process started",
	}

	entity := entities.SupplierPrice{}
	history := entities.SupplierPriceHistory{}
	dto := dto.SupplierPriceRequestDTO{}
	errJsonDecode := json.NewDecoder(r.Body).Decode(&dto)
	if errJsonDecode != nil {
		logger.LogError(errJsonDecode)
		res.Message = fmt.Sprintf("json decoding error: %v", errJsonDecode)
		json.NewEncoder(w).Encode(res)
		return
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entity), sh.env, &dto, entity, &history)
	id, errService := service.CreateNewItem(&dto)
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("New supplierprice created with ID: %d", id)
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) UpdateSupplierPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                                   // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add supplierprice init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "supplierprice updating process started",
	}

	entity := entities.SupplierPrice{}
	history := entities.SupplierPriceHistory{}
	dto := dto.SupplierPriceRequestDTO{}
	errJsonDecode := json.NewDecoder(r.Body).Decode(&dto)
	if errJsonDecode != nil {
		logger.LogError(errJsonDecode)
		res.Message = fmt.Sprintf("json decoding error: %v", errJsonDecode)
		json.NewEncoder(w).Encode(res)
		return
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entity), sh.env, &dto, entity, &history)
	errService := service.UpdateItem(&dto)
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("SupplierPrice updated.")
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) DeleteSupplierPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                                      // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "delete supplierprice init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "supplierprice deleting process started",
	}

	vars := mux.Vars(r)
	id, _ := vars["id"]
	idInt, errConv := strconv.Atoi(id)
	if errConv != nil {
		defer logger.LogError(errConv)
		res.Message = fmt.Sprintf("couldn't convert given id, error: %v", errConv)
		json.NewEncoder(w).Encode(res)
		return
	}
	Id := uint(idInt)

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.SupplierPrice{}), sh.env, &dto.SupplierPriceRequestDTO{}, entities.SupplierPrice{}, &entities.SupplierPriceHistory{})
	errService := service.DeleteItem(Id)
	if errService != nil {
		defer logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("SupplierPrice deleted.")
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) AddNewTaskConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                                // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add taskconfig init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "new taskconfig adding process started",
	}

	entity := entities.TaskConfig{}
	history := entities.TaskConfigHistory{}
	dto := dto.TaskConfigRequestDTO{}
	errJsonDecode := json.NewDecoder(r.Body).Decode(&dto)
	if errJsonDecode != nil {
		logger.LogError(errJsonDecode)
		res.Message = fmt.Sprintf("json decoding error: %v", errJsonDecode)
		json.NewEncoder(w).Encode(res)
		return
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entity), sh.env, &dto, entity, &history)
	id, errService := service.CreateNewItem(&dto)
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("New taskconfig created with ID: %d", id)
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) UpdateTaskConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                                // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add taskconfig init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "taskconfig updating process started",
	}

	entity := entities.TaskConfig{}
	history := entities.TaskConfigHistory{}
	dto := dto.TaskConfigRequestDTO{}
	errJsonDecode := json.NewDecoder(r.Body).Decode(&dto)
	if errJsonDecode != nil {
		logger.LogError(errJsonDecode)
		res.Message = fmt.Sprintf("json decoding error: %v", errJsonDecode)
		json.NewEncoder(w).Encode(res)
		return
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entity), sh.env, &dto, entity, &history)
	errService := service.UpdateItem(&dto)
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("TaskConfig updated.")
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) DeleteTaskConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                                   // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "delete taskconfig init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "taskconfig deleting process started",
	}

	vars := mux.Vars(r)
	id, _ := vars["id"]
	idInt, errConv := strconv.Atoi(id)
	if errConv != nil {
		defer logger.LogError(errConv)
		res.Message = fmt.Sprintf("couldn't convert given id, error: %v", errConv)
		json.NewEncoder(w).Encode(res)
		return
	}
	Id := uint(idInt)

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.TaskConfig{}), sh.env, &dto.TaskConfigRequestDTO{}, entities.TaskConfig{}, &entities.TaskConfigHistory{})
	errService := service.DeleteItem(Id)
	if errService != nil {
		defer logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("TaskConfig deleted.")
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) AddNewQuestion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                              // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add question init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "new question adding process started",
	}

	entity := entities.Question{}
	history := entities.QuestionHistory{}
	dto := dto.QuestionRequestDTO{}
	errJsonDecode := json.NewDecoder(r.Body).Decode(&dto)
	if errJsonDecode != nil {
		logger.LogError(errJsonDecode)
		res.Message = fmt.Sprintf("json decoding error: %v", errJsonDecode)
		json.NewEncoder(w).Encode(res)
		return
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entity), sh.env, &dto, entity, &history)
	id, errService := service.CreateNewItem(&dto)
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("New question created with ID: %d", id)
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) UpdateQuestion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                              // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add question init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "question updating process started",
	}

	entity := entities.Question{}
	history := entities.QuestionHistory{}
	dto := dto.QuestionRequestDTO{}
	errJsonDecode := json.NewDecoder(r.Body).Decode(&dto)
	if errJsonDecode != nil {
		logger.LogError(errJsonDecode)
		res.Message = fmt.Sprintf("json decoding error: %v", errJsonDecode)
		json.NewEncoder(w).Encode(res)
		return
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entity), sh.env, &dto, entity, &history)
	errService := service.UpdateItem(&dto)
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("Question updated.")
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                                 // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "delete question init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "question deleting process started",
	}

	vars := mux.Vars(r)
	id, _ := vars["id"]
	idInt, errConv := strconv.Atoi(id)
	if errConv != nil {
		defer logger.LogError(errConv)
		res.Message = fmt.Sprintf("couldn't convert given id, error: %v", errConv)
		json.NewEncoder(w).Encode(res)
		return
	}
	Id := uint(idInt)

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.Question{}), sh.env, &dto.QuestionRequestDTO{}, entities.Question{}, &entities.QuestionHistory{})
	errService := service.DeleteItem(Id)
	if errService != nil {
		defer logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("Question deleted.")
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) AddNewTaskOffered(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                                 // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add taskoffered init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "new taskoffered adding process started",
	}

	entity := entities.TaskOffered{}
	history := entities.TaskOfferedHistory{}
	dto := dto.TaskOfferedRequestDTO{}
	errJsonDecode := json.NewDecoder(r.Body).Decode(&dto)
	if errJsonDecode != nil {
		logger.LogError(errJsonDecode)
		res.Message = fmt.Sprintf("json decoding error: %v", errJsonDecode)
		json.NewEncoder(w).Encode(res)
		return
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entity), sh.env, &dto, entity, &history)
	id, errService := service.CreateNewItem(&dto)
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("New taskoffered created with ID: %d", id)
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) UpdateTaskOffered(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                                 // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add taskoffered init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "taskoffered updating process started",
	}

	entity := entities.TaskOffered{}
	history := entities.TaskOfferedHistory{}
	dto := dto.TaskOfferedRequestDTO{}
	errJsonDecode := json.NewDecoder(r.Body).Decode(&dto)
	if errJsonDecode != nil {
		logger.LogError(errJsonDecode)
		res.Message = fmt.Sprintf("json decoding error: %v", errJsonDecode)
		json.NewEncoder(w).Encode(res)
		return
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entity), sh.env, &dto, entity, &history)
	errService := service.UpdateItem(&dto)
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("TaskOffered updated.")
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) DeleteTaskOffered(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                                    // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "delete taskoffered init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "taskoffered deleting process started",
	}

	vars := mux.Vars(r)
	id, _ := vars["id"]
	idInt, errConv := strconv.Atoi(id)
	if errConv != nil {
		defer logger.LogError(errConv)
		res.Message = fmt.Sprintf("couldn't convert given id, error: %v", errConv)
		json.NewEncoder(w).Encode(res)
		return
	}
	Id := uint(idInt)

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.TaskOffered{}), sh.env, &dto.TaskOfferedRequestDTO{}, entities.TaskOffered{}, &entities.TaskOfferedHistory{})
	errService := service.DeleteItem(Id)
	if errService != nil {
		defer logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("TaskOffered deleted.")
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) AddNewBatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                           // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add batch init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "new batch adding process started",
	}

	entity := entities.Batch{}
	history := entities.BatchHistory{}
	dto := dto.BatchRequestDTO{}
	errJsonDecode := json.NewDecoder(r.Body).Decode(&dto)
	if errJsonDecode != nil {
		logger.LogError(errJsonDecode)
		res.Message = fmt.Sprintf("json decoding error: %v", errJsonDecode)
		json.NewEncoder(w).Encode(res)
		return
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entity), sh.env, &dto, entity, &history)
	id, errService := service.CreateNewItem(&dto)
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("New batch created with ID: %d", id)
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) UpdateBatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                           // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add batch init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "batch updating process started",
	}

	entity := entities.Batch{}
	history := entities.BatchHistory{}
	dto := dto.BatchRequestDTO{}
	errJsonDecode := json.NewDecoder(r.Body).Decode(&dto)
	if errJsonDecode != nil {
		logger.LogError(errJsonDecode)
		res.Message = fmt.Sprintf("json decoding error: %v", errJsonDecode)
		json.NewEncoder(w).Encode(res)
		return
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entity), sh.env, &dto, entity, &history)
	errService := service.UpdateItem(&dto)
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("Batch updated.")
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) DeleteBatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                              // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "delete batch init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "batch deleting process started",
	}

	vars := mux.Vars(r)
	id, _ := vars["id"]
	idInt, errConv := strconv.Atoi(id)
	if errConv != nil {
		defer logger.LogError(errConv)
		res.Message = fmt.Sprintf("couldn't convert given id, error: %v", errConv)
		json.NewEncoder(w).Encode(res)
		return
	}
	Id := uint(idInt)

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.Batch{}), sh.env, &dto.BatchRequestDTO{}, entities.Batch{}, &entities.BatchHistory{})
	errService := service.DeleteItem(Id)
	if errService != nil {
		defer logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("Batch deleted.")
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) AddNewCustomerPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                                   // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add customerprice init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "new customerprice adding process started",
	}

	entity := entities.CustomerPrice{}
	history := entities.CustomerPriceHistory{}
	dto := dto.CustomerPriceRequestDTO{}
	errJsonDecode := json.NewDecoder(r.Body).Decode(&dto)
	if errJsonDecode != nil {
		logger.LogError(errJsonDecode)
		res.Message = fmt.Sprintf("json decoding error: %v", errJsonDecode)
		json.NewEncoder(w).Encode(res)
		return
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entity), sh.env, &dto, entity, &history)
	id, errService := service.CreateNewItem(&dto)
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("New customerprice created with ID: %d", id)
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) UpdateCustomerPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                                   // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add customerprice init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "customerprice updating process started",
	}

	entity := entities.CustomerPrice{}
	history := entities.CustomerPriceHistory{}
	dto := dto.CustomerPriceRequestDTO{}
	errJsonDecode := json.NewDecoder(r.Body).Decode(&dto)
	if errJsonDecode != nil {
		logger.LogError(errJsonDecode)
		res.Message = fmt.Sprintf("json decoding error: %v", errJsonDecode)
		json.NewEncoder(w).Encode(res)
		return
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entity), sh.env, &dto, entity, &history)
	errService := service.UpdateItem(&dto)
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("CustomerPrice updated.")
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) DeleteCustomerPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                                      // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "delete customerprice init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "customerprice deleting process started",
	}

	vars := mux.Vars(r)
	id, _ := vars["id"]
	idInt, errConv := strconv.Atoi(id)
	if errConv != nil {
		defer logger.LogError(errConv)
		res.Message = fmt.Sprintf("couldn't convert given id, error: %v", errConv)
		json.NewEncoder(w).Encode(res)
		return
	}
	Id := uint(idInt)

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.CustomerPrice{}), sh.env, &dto.CustomerPriceRequestDTO{}, entities.CustomerPrice{}, &entities.CustomerPriceHistory{})
	errService := service.DeleteItem(Id)
	if errService != nil {
		defer logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("CustomerPrice deleted.")
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) AddNewClientOffer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                                 // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add clientoffer init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "new clientoffer adding process started",
	}

	entity := entities.ClientOffer{}
	history := entities.ClientOfferHistory{}
	dto := dto.ClientOfferRequestDTO{}
	errJsonDecode := json.NewDecoder(r.Body).Decode(&dto)
	if errJsonDecode != nil {
		logger.LogError(errJsonDecode)
		res.Message = fmt.Sprintf("json decoding error: %v", errJsonDecode)
		json.NewEncoder(w).Encode(res)
		return
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entity), sh.env, &dto, entity, &history)
	id, errService := service.CreateNewItem(&dto)
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("New clientoffer created with ID: %d", id)
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) UpdateClientOffer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                                 // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add clientoffer init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "clientoffer updating process started",
	}

	entity := entities.ClientOffer{}
	history := entities.ClientOfferHistory{}
	dto := dto.ClientOfferRequestDTO{}
	errJsonDecode := json.NewDecoder(r.Body).Decode(&dto)
	if errJsonDecode != nil {
		logger.LogError(errJsonDecode)
		res.Message = fmt.Sprintf("json decoding error: %v", errJsonDecode)
		json.NewEncoder(w).Encode(res)
		return
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entity), sh.env, &dto, entity, &history)
	errService := service.UpdateItem(&dto)
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("ClientOffer updated.")
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) DeleteClientOffer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                                    // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "delete clientoffer init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "clientoffer deleting process started",
	}

	vars := mux.Vars(r)
	id, _ := vars["id"]
	idInt, errConv := strconv.Atoi(id)
	if errConv != nil {
		defer logger.LogError(errConv)
		res.Message = fmt.Sprintf("couldn't convert given id, error: %v", errConv)
		json.NewEncoder(w).Encode(res)
		return
	}
	Id := uint(idInt)

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.ClientOffer{}), sh.env, &dto.ClientOfferRequestDTO{}, entities.ClientOffer{}, &entities.ClientOfferHistory{})
	errService := service.DeleteItem(Id)
	if errService != nil {
		defer logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("ClientOffer deleted.")
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) AddNewUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add user init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "new user adding process started",
	}

	entity := entities.User{}
	history := entities.UserHistory{}
	dto := dto.UserRequestDTO{}
	errJsonDecode := json.NewDecoder(r.Body).Decode(&dto)
	if errJsonDecode != nil {
		logger.LogError(errJsonDecode)
		res.Message = fmt.Sprintf("json decoding error: %v", errJsonDecode)
		json.NewEncoder(w).Encode(res)
		return
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entity), sh.env, &dto, entity, &history)
	id, errService := service.CreateNewItem(&dto)
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("New user created with ID: %d", id)
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add user init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "user updating process started",
	}

	entity := entities.User{}
	history := entities.UserHistory{}
	dto := dto.UserRequestDTO{}
	errJsonDecode := json.NewDecoder(r.Body).Decode(&dto)
	if errJsonDecode != nil {
		logger.LogError(errJsonDecode)
		res.Message = fmt.Sprintf("json decoding error: %v", errJsonDecode)
		json.NewEncoder(w).Encode(res)
		return
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entity), sh.env, &dto, entity, &history)
	errService := service.UpdateItem(&dto)
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("User updated.")
	json.NewEncoder(w).Encode(res)

}

func (sh ServerHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                             // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "delete user init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "user deleting process started",
	}

	vars := mux.Vars(r)
	id, _ := vars["id"]
	idInt, errConv := strconv.Atoi(id)
	if errConv != nil {
		defer logger.LogError(errConv)
		res.Message = fmt.Sprintf("couldn't convert given id, error: %v", errConv)
		json.NewEncoder(w).Encode(res)
		return
	}
	Id := uint(idInt)

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.User{}), sh.env, &dto.UserRequestDTO{}, entities.User{}, &entities.UserHistory{})
	errService := service.DeleteItem(Id)
	if errService != nil {
		defer logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Message = fmt.Sprintf("User deleted.")
	json.NewEncoder(w).Encode(res)

}
