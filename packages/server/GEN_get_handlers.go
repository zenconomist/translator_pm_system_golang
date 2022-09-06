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


func (sh ServerHandler) GetAllTaskStateChangeComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add taskstatechangecomment init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "getting taskstatechangecomment",
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.TaskStateChangeComment{}), sh.env, &dto.TaskStateChangeCommentResponseDTO{}, entities.TaskStateChangeComment{}, &entities.TaskStateChangeCommentHistory{})
	taskstatechangecomments, errService := service.GetItems()
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(taskstatechangecomments)

}

func (sh ServerHandler) GetTaskStateChangeCommentByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add taskstatechangecomment init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "getting taskstatechangecomment",
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

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.TaskStateChangeComment{}), sh.env, &dto.TaskStateChangeCommentResponseDTO{}, entities.TaskStateChangeComment{}, &entities.TaskStateChangeCommentHistory{})
	taskstatechangecomment, errService := service.GetItemByID(Id)
	if errService != nil {
		defer logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(taskstatechangecomment)

}


func (sh ServerHandler) GetAllBatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add batch init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "getting batch",
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.Batch{}), sh.env, &dto.BatchResponseDTO{}, entities.Batch{}, &entities.BatchHistory{})
	batchs, errService := service.GetItems()
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(batchs)

}

func (sh ServerHandler) GetBatchByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add batch init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "getting batch",
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

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.Batch{}), sh.env, &dto.BatchResponseDTO{}, entities.Batch{}, &entities.BatchHistory{})
	batch, errService := service.GetItemByID(Id)
	if errService != nil {
		defer logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(batch)

}


func (sh ServerHandler) GetAllContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add contact init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "getting contact",
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.Contact{}), sh.env, &dto.ContactResponseDTO{}, entities.Contact{}, &entities.ContactHistory{})
	contacts, errService := service.GetItems()
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(contacts)

}

func (sh ServerHandler) GetContactByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add contact init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "getting contact",
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

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.Contact{}), sh.env, &dto.ContactResponseDTO{}, entities.Contact{}, &entities.ContactHistory{})
	contact, errService := service.GetItemByID(Id)
	if errService != nil {
		defer logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(contact)

}


func (sh ServerHandler) GetAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add user init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "getting user",
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.User{}), sh.env, &dto.UserResponseDTO{}, entities.User{}, &entities.UserHistory{})
	users, errService := service.GetItems()
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(users)

}

func (sh ServerHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add user init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "getting user",
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

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.User{}), sh.env, &dto.UserResponseDTO{}, entities.User{}, &entities.UserHistory{})
	user, errService := service.GetItemByID(Id)
	if errService != nil {
		defer logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(user)

}


func (sh ServerHandler) GetAllCustomerPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add customerprice init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "getting customerprice",
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.CustomerPrice{}), sh.env, &dto.CustomerPriceResponseDTO{}, entities.CustomerPrice{}, &entities.CustomerPriceHistory{})
	customerprices, errService := service.GetItems()
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(customerprices)

}

func (sh ServerHandler) GetCustomerPriceByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add customerprice init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "getting customerprice",
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

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.CustomerPrice{}), sh.env, &dto.CustomerPriceResponseDTO{}, entities.CustomerPrice{}, &entities.CustomerPriceHistory{})
	customerprice, errService := service.GetItemByID(Id)
	if errService != nil {
		defer logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(customerprice)

}


func (sh ServerHandler) GetAllSupplierPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add supplierprice init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "getting supplierprice",
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.SupplierPrice{}), sh.env, &dto.SupplierPriceResponseDTO{}, entities.SupplierPrice{}, &entities.SupplierPriceHistory{})
	supplierprices, errService := service.GetItems()
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(supplierprices)

}

func (sh ServerHandler) GetSupplierPriceByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add supplierprice init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "getting supplierprice",
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

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.SupplierPrice{}), sh.env, &dto.SupplierPriceResponseDTO{}, entities.SupplierPrice{}, &entities.SupplierPriceHistory{})
	supplierprice, errService := service.GetItemByID(Id)
	if errService != nil {
		defer logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(supplierprice)

}


func (sh ServerHandler) GetAllDefaultCustomerPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add defaultcustomerprice init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "getting defaultcustomerprice",
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.DefaultCustomerPrice{}), sh.env, &dto.DefaultCustomerPriceResponseDTO{}, entities.DefaultCustomerPrice{}, &entities.DefaultCustomerPriceHistory{})
	defaultcustomerprices, errService := service.GetItems()
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(defaultcustomerprices)

}

func (sh ServerHandler) GetDefaultCustomerPriceByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add defaultcustomerprice init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "getting defaultcustomerprice",
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

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.DefaultCustomerPrice{}), sh.env, &dto.DefaultCustomerPriceResponseDTO{}, entities.DefaultCustomerPrice{}, &entities.DefaultCustomerPriceHistory{})
	defaultcustomerprice, errService := service.GetItemByID(Id)
	if errService != nil {
		defer logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(defaultcustomerprice)

}


func (sh ServerHandler) GetAllUPMLogger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add upmlogger init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "getting upmlogger",
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.UPMLogger{}), sh.env, &dto.UPMLoggerResponseDTO{}, entities.UPMLogger{}, &entities.UPMLoggerHistory{})
	upmloggers, errService := service.GetItems()
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(upmloggers)

}

func (sh ServerHandler) GetUPMLoggerByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add upmlogger init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "getting upmlogger",
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

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.UPMLogger{}), sh.env, &dto.UPMLoggerResponseDTO{}, entities.UPMLogger{}, &entities.UPMLoggerHistory{})
	upmlogger, errService := service.GetItemByID(Id)
	if errService != nil {
		defer logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(upmlogger)

}


func (sh ServerHandler) GetAllClientOffer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add clientoffer init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "getting clientoffer",
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.ClientOffer{}), sh.env, &dto.ClientOfferResponseDTO{}, entities.ClientOffer{}, &entities.ClientOfferHistory{})
	clientoffers, errService := service.GetItems()
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(clientoffers)

}

func (sh ServerHandler) GetClientOfferByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add clientoffer init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "getting clientoffer",
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

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.ClientOffer{}), sh.env, &dto.ClientOfferResponseDTO{}, entities.ClientOffer{}, &entities.ClientOfferHistory{})
	clientoffer, errService := service.GetItemByID(Id)
	if errService != nil {
		defer logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(clientoffer)

}


func (sh ServerHandler) GetAllDefaultSupplierPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add defaultsupplierprice init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "getting defaultsupplierprice",
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.DefaultSupplierPrice{}), sh.env, &dto.DefaultSupplierPriceResponseDTO{}, entities.DefaultSupplierPrice{}, &entities.DefaultSupplierPriceHistory{})
	defaultsupplierprices, errService := service.GetItems()
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(defaultsupplierprices)

}

func (sh ServerHandler) GetDefaultSupplierPriceByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add defaultsupplierprice init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "getting defaultsupplierprice",
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

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.DefaultSupplierPrice{}), sh.env, &dto.DefaultSupplierPriceResponseDTO{}, entities.DefaultSupplierPrice{}, &entities.DefaultSupplierPriceHistory{})
	defaultsupplierprice, errService := service.GetItemByID(Id)
	if errService != nil {
		defer logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(defaultsupplierprice)

}


func (sh ServerHandler) GetAllQuestion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add question init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "getting question",
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.Question{}), sh.env, &dto.QuestionResponseDTO{}, entities.Question{}, &entities.QuestionHistory{})
	questions, errService := service.GetItems()
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(questions)

}

func (sh ServerHandler) GetQuestionByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add question init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "getting question",
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

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.Question{}), sh.env, &dto.QuestionResponseDTO{}, entities.Question{}, &entities.QuestionHistory{})
	question, errService := service.GetItemByID(Id)
	if errService != nil {
		defer logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(question)

}


func (sh ServerHandler) GetAllFirm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add firm init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "getting firm",
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.Firm{}), sh.env, &dto.FirmResponseDTO{}, entities.Firm{}, &entities.FirmHistory{})
	firms, errService := service.GetItems()
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(firms)

}

func (sh ServerHandler) GetFirmByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add firm init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "getting firm",
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

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.Firm{}), sh.env, &dto.FirmResponseDTO{}, entities.Firm{}, &entities.FirmHistory{})
	firm, errService := service.GetItemByID(Id)
	if errService != nil {
		defer logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(firm)

}


func (sh ServerHandler) GetAllTaskOffered(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add taskoffered init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "getting taskoffered",
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.TaskOffered{}), sh.env, &dto.TaskOfferedResponseDTO{}, entities.TaskOffered{}, &entities.TaskOfferedHistory{})
	taskoffereds, errService := service.GetItems()
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(taskoffereds)

}

func (sh ServerHandler) GetTaskOfferedByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add taskoffered init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "getting taskoffered",
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

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.TaskOffered{}), sh.env, &dto.TaskOfferedResponseDTO{}, entities.TaskOffered{}, &entities.TaskOfferedHistory{})
	taskoffered, errService := service.GetItemByID(Id)
	if errService != nil {
		defer logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(taskoffered)

}


func (sh ServerHandler) GetAllBillingLog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add billinglog init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "getting billinglog",
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.BillingLog{}), sh.env, &dto.BillingLogResponseDTO{}, entities.BillingLog{}, &entities.BillingLogHistory{})
	billinglogs, errService := service.GetItems()
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(billinglogs)

}

func (sh ServerHandler) GetBillingLogByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add billinglog init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "getting billinglog",
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

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.BillingLog{}), sh.env, &dto.BillingLogResponseDTO{}, entities.BillingLog{}, &entities.BillingLogHistory{})
	billinglog, errService := service.GetItemByID(Id)
	if errService != nil {
		defer logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(billinglog)

}


func (sh ServerHandler) GetAllTaskConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add taskconfig init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "getting taskconfig",
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.TaskConfig{}), sh.env, &dto.TaskConfigResponseDTO{}, entities.TaskConfig{}, &entities.TaskConfigHistory{})
	taskconfigs, errService := service.GetItems()
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(taskconfigs)

}

func (sh ServerHandler) GetTaskConfigByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add taskconfig init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "getting taskconfig",
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

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.TaskConfig{}), sh.env, &dto.TaskConfigResponseDTO{}, entities.TaskConfig{}, &entities.TaskConfigHistory{})
	taskconfig, errService := service.GetItemByID(Id)
	if errService != nil {
		defer logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(taskconfig)

}


func (sh ServerHandler) GetAllProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add project init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "getting project",
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.Project{}), sh.env, &dto.ProjectResponseDTO{}, entities.Project{}, &entities.ProjectHistory{})
	projects, errService := service.GetItems()
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(projects)

}

func (sh ServerHandler) GetProjectByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add project init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "getting project",
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

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.Project{}), sh.env, &dto.ProjectResponseDTO{}, entities.Project{}, &entities.ProjectHistory{})
	project, errService := service.GetItemByID(Id)
	if errService != nil {
		defer logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(project)

}


func (sh ServerHandler) GetAllTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add task init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "getting task",
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.Task{}), sh.env, &dto.TaskResponseDTO{}, entities.Task{}, &entities.TaskHistory{})
	tasks, errService := service.GetItems()
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(tasks)

}

func (sh ServerHandler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add task init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "getting task",
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

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.Task{}), sh.env, &dto.TaskResponseDTO{}, entities.Task{}, &entities.TaskHistory{})
	task, errService := service.GetItemByID(Id)
	if errService != nil {
		defer logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(task)

}


func (sh ServerHandler) GetAllCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add customer init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "getting customer",
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.Customer{}), sh.env, &dto.CustomerResponseDTO{}, entities.Customer{}, &entities.CustomerHistory{})
	customers, errService := service.GetItems()
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(customers)

}

func (sh ServerHandler) GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add customer init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "getting customer",
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

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.Customer{}), sh.env, &dto.CustomerResponseDTO{}, entities.Customer{}, &entities.CustomerHistory{})
	customer, errService := service.GetItemByID(Id)
	if errService != nil {
		defer logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(customer)

}


func (sh ServerHandler) GetAllClientOfferTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add clientoffertask init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "getting clientoffertask",
	}

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.ClientOfferTask{}), sh.env, &dto.ClientOfferTaskResponseDTO{}, entities.ClientOfferTask{}, &entities.ClientOfferTaskHistory{})
	clientoffertasks, errService := service.GetItems()
	if errService != nil {
		logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(clientoffertasks)

}

func (sh ServerHandler) GetClientOfferTaskByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logger := sh.env.GiveLogger()                                          // getting pointer of initialized logger
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "add clientoffertask init") // initializing logger data for handler
	defer logger.Log()

	defer sh.recoverFromPanic() //default panic->recover
	res := Response{
		Message: "getting clientoffertask",
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

	service := services.NewService(entities.NewRepository(sh.env.GiveDbHandler().PassConnection(), entities.ClientOfferTask{}), sh.env, &dto.ClientOfferTaskResponseDTO{}, entities.ClientOfferTask{}, &entities.ClientOfferTaskHistory{})
	clientoffertask, errService := service.GetItemByID(Id)
	if errService != nil {
		defer logger.LogError(errService)
		res.Message = fmt.Sprintf("service error: %v", errService)
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(clientoffertask)

}
