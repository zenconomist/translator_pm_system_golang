package services

import (
	"dto"
	"entities"
	"environment"
	"fmt"
	"reflect"
	"strings"

	"gorm.io/gorm"
)

/*
	Main responsibilities of services:
		1. Error logging from repositories/entities
        2. Perform Validations
		3. Perform Dto mappings
*/

// --SERVICE WITH DTO -> non-generic Dto mapping-- \\

type IService[R entities.IRepository[T], E environment.Environment, D dto.DTO[T], T entities.Entity] interface {
	GiveDb() *gorm.DB
	GetItems() ([]D, error)
	GetItemsWithIDs() ([]D, error)
	GetItemByID(uint) (D, error)
	CreateNewItem(D) (uint, error)
	UpdateItem(D) error
	DeleteItem(uint) error
}

type Service[R entities.IRepository[T], E environment.Environment, D dto.DTO[T], T entities.Entity, H entities.EntityHistory] struct {
	Repo    R
	Entity  T
	History H
	env     E
	Dto     D
	NewID   uint
}

func NewService[R entities.IRepository[T], E environment.Environment, D dto.DTO[T], T entities.Entity, H entities.EntityHistory](repo R, env E, dto D, entity T, history H) *Service[R, E, D, T, H] {
	return &Service[R, E, D, T, H]{
		Repo: repo,
		env:  env,
		Dto:  dto,
	}

}

func (service *Service[R, E, D, T, H]) EmptyService() {
	fmt.Println("service up and running")
}

func (service *Service[R, E, D, T, H]) GiveDb() *gorm.DB {
	return service.env.GiveDbHandler().PassConnection()
}

func (service *Service[R, E, D, T, H]) GetItems() ([]D, error) {
	var dtos []D
	items, err := service.Repo.FindAll()
	if err != nil {
		defer service.env.GiveLogger().LogError(err)
		return dtos, err
	}

	dtoHandlerMulti := dto.NewDtoHandlerMulti(items, dtos, service.env)
	if errMap := dtoHandlerMulti.MapFromEntities(); errMap != nil {
		return dtoHandlerMulti.Dtos, errMap
	}
	return dtoHandlerMulti.Dtos, nil
}

func (service *Service[R, E, D, T, H]) GetItemsWithIDs(ids []uint) ([]D, error) {
	var dtos []D
	var dtoByVal D
	var errMapping error
	items, err := service.Repo.FindAllWithIDs(ids)
	if err != nil {
		defer service.env.GiveLogger().LogError(err)
		return dtos, err
	}

	for _, item := range items {
		dtoHandler := dto.NewDtoHandler(&item, &service.Dto, service.env)
		dtoByVal, errMapping = dtoHandler.MapFromEntity()
		if errMapping != nil {
			defer service.env.GiveLogger().LogError(err)
			return dtos, err
		}
		dtos = append(dtos, dtoByVal)
	}
	return dtos, nil
}

func (service *Service[R, E, D, T, H]) GetItemByID(ID uint) (D, error) {
	dto := service.Dto
	item, err := service.Repo.FindByID(ID)
	if err != nil {
		defer service.env.GiveLogger().LogError(err)
		return dto, err
	}
	errMapping := dto.MapFromEntity(service.GiveDb(), item)
	if errMapping != nil {
		defer service.env.GiveLogger().LogError(err)
		return dto, err
	}
	return dto, nil
}

func (service *Service[R, E, D, T, H]) CreateNewItem(data D) (uint, error) {
	service.env.GiveLogger().InitLoggerPerFunc(service.env.GiveLogger().GetCurrentFuncName(), "create new item")
	defer service.env.GiveLogger().Log()

	var defaultID uint
	// default transformations on DTO-s
	errCustDTOTransform := data.CustomDTOTransformations(service.GiveDb())
	if errCustDTOTransform != nil {
		defer service.env.GiveLogger().LogError(errCustDTOTransform)
		return defaultID, errCustDTOTransform
	}
	// validations
	errCustomValidations := data.CustomValidations(service.GiveDb(), service.env.GiveLogger())
	if errCustomValidations != nil {
		defer service.env.GiveLogger().LogError(errCustomValidations)
		return defaultID, errCustomValidations
	}
	// net mapping
	var ent T
	entity, errMap := data.MapToEntity(service.GiveDb(), &ent)
	if errMap != nil {
		defer service.env.GiveLogger().LogError(errMap)
		return defaultID, errMap
	}
	ID, err := service.Repo.Create(*entity)
	if err != nil {
		defer service.env.GiveLogger().LogError(err)
		return ID, err
	}
	service.NewID = ID
	done := make(chan bool)
	// task historicization is handled through an action object with sql script.
	if reflect.TypeOf(*entity).String() != "entities.Task" {
		defer service.createHistory(*entity, ID)
	}
	go service.ExecuteActions(service.env.GiveLogger().GetCurrentFuncName(), done)

	fmt.Println("end of default service call")
	return ID, err
}

func (service *Service[R, E, D, T, H]) createHistory(entity T, ID uint) error {
	// historicization
	hrepo := entities.NewHistoryRepo(service.env.GiveDbHandler().PassConnection(), service.History, entities.NewHistoryMapper(service.History, entity), entity)
	hist, errMapToHistory := hrepo.HistoryMapper.MapToHistoryForCreation(entity)
	if errMapToHistory != nil {
		return errMapToHistory
	}
	hist.SetModelToCreated(ID)
	errInsert := hrepo.Create(&hist)
	if errInsert != nil {
		return errInsert
	}
	return nil
}

func (service *Service[R, E, D, T, H]) UpdateItem(data D) error {
	logger := service.env.GiveLogger()
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "update user")
	defer logger.Log()

	id, errGetItem := data.GiveID()
	if errGetItem != nil {
		logger.SetInfo("couldn't get item based on DTO's ID value.")
		// critical error
		defer logger.LogError(errGetItem)
		return errGetItem
	}
	entity, errFind := service.Repo.FindByID(id)
	if errFind != nil {
		defer logger.LogError(errFind)
		return errFind
	}
	// dto -> entity mapping implementation
	entityPtr, errMap := data.MapToEntity(service.GiveDb(), &entity)
	if errMap != nil {
		defer logger.LogError(errMap)
		return errMap
	}

	err := service.Repo.Update(*entityPtr)
	if err != nil {
		defer service.env.GiveLogger().LogError(err)
		return err
	}
	done := make(chan bool)
	defer service.updateHistory(entity)
	go service.ExecuteActions(service.env.GiveLogger().GetCurrentFuncName(), done)

	return nil
}

func (service *Service[R, E, D, T, H]) updateHistory(entity T) error {
	// historicization
	hrepo := entities.NewHistoryRepo(service.env.GiveDbHandler().PassConnection(), service.History, entities.NewHistoryMapper(service.History, entity), entity)
	hist, errMapToHistory := hrepo.HistoryMapper.MapToHistoryForCreation(entity)
	if errMapToHistory != nil {
		return errMapToHistory
	}
	hist.SetModelToUpdated(entity.GiveID())
	errUpdate := hrepo.Update(hist, entity)
	if errUpdate != nil {
		return errUpdate
	}
	hist.SetModelToCreated(entity.GiveID())
	errInsert := hrepo.Create(&hist)
	if errInsert != nil {
		return errInsert
	}
	return nil
}

func (service *Service[R, E, D, T, H]) DeleteItem(id uint) error {
	logger := service.env.GiveLogger()
	logger.InitLoggerPerFunc(logger.GetCurrentFuncName(), "create new user")
	defer logger.Log()

	entity, errFind := service.Repo.FindByID(id)
	if errFind != nil {
		defer service.env.GiveLogger().LogError(errFind)
		return errFind
	}

	err := service.Repo.Delete(entity)
	if err != nil {
		defer service.env.GiveLogger().LogError(err)
		return err
	}
	done := make(chan bool)
	defer service.deleteHistory(entity)
	go service.ExecuteActions(service.env.GiveLogger().GetCurrentFuncName(), done)

	return nil
}

func (service *Service[R, E, D, T, H]) deleteHistory(entity T) error {
	// historicization
	hrepo := entities.NewHistoryRepo(service.env.GiveDbHandler().PassConnection(), service.History, entities.NewHistoryMapper(service.History, entity), entity)
	hist, errMapToHistory := hrepo.HistoryMapper.MapToHistoryForCreation(entity)
	if errMapToHistory != nil {
		return errMapToHistory
	}
	hist.SetModelToDeleted(entity.GiveID())
	errUpdate := hrepo.Update(hist, entity)
	if errUpdate != nil {
		return errUpdate
	}
	return nil
}

func (service *Service[R, E, D, T, H]) ExecuteActions(funcname string, done chan bool) {
	logger := service.env.GiveLogger()
	actionObjects := strings.Builder{}
	actionObjects.WriteString("actionObjects: ")
	actions := GiveActions(funcname, service.Entity, service.Dto, service.env, service.NewID)

	go func() {
		for _, a := range actions {
			if errExec := a.Execute(); errExec != nil {
				defer service.env.GiveLogger().LogError(errExec)
			}
			actionObjects.WriteString(reflect.TypeOf(a).String())
			actionObjects.WriteString("; ")

		}
		done <- true
	}()
	logger.SetInfo("action for " + funcname + ", executed for entity: " + reflect.TypeOf(service.Entity).String() + actionObjects.String())
	logger.Log()

	<-done

}
