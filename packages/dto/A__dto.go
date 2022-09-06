package dto

import (
	"encoding/json"
	"entities"
	"environment"
	"logging"

	"gorm.io/gorm"
)

type DTO[E entities.Entity] interface {
	MapToEntity(*gorm.DB, *E) (*E, error)
	MapFromEntity(*gorm.DB, E) error
	CustomDTOTransformations(*gorm.DB) error
	CustomValidations(*gorm.DB, logging.Logger) error
	customMappingToEntity(*gorm.DB) error
	customMappingFromEntity(E, *gorm.DB) error
	GiveID() (uint, error)
	// MapFromEntityToDtoByVal(*gorm.DB, E) (IDTO, error)
}

type IDTO interface {
	GiveID() (uint, error)
}

// --DTOHANDLER - currently not in use--\\

type IDTOHandler[E entities.Entity, D DTO[E]] interface {
	MapToEntity(*E) (*E, error)
	MapFromEntity(E) (D, error)
}

type DTOHandler[E entities.Entity, D DTO[E]] struct {
	Dto    *D
	Entity *E
	Env    environment.Environment
}

func NewDtoHandler[E entities.Entity, D DTO[E]](entity *E, dto *D, env environment.Environment) *DTOHandler[E, D] {
	return &DTOHandler[E, D]{
		Env:    env,
		Dto:    dto,
		Entity: entity,
	}
}

func (dh *DTOHandler[E, D]) MapToEntity() (E, error) {
	data := *dh.Dto
	errorCustomMapping := data.customMappingToEntity(dh.Env.GiveDbHandler().PassConnection())
	if errorCustomMapping != nil {
		return *dh.Entity, errorCustomMapping
	}
	jsonDtoBytes, errJsonMarshal := json.Marshal(data)
	if errJsonMarshal != nil {
		return *dh.Entity, errJsonMarshal
	}
	errJsonUnmarshal := json.Unmarshal(jsonDtoBytes, *dh.Entity)
	if errJsonUnmarshal != nil {
		return *dh.Entity, errJsonUnmarshal
	}
	return *dh.Entity, nil
}

func (dh *DTOHandler[E, D]) MapFromEntity() (D, error) {
	var dto D
	jsonDataBytes, errJsonMarshal := json.Marshal(dh.Entity)
	if errJsonMarshal != nil {
		return dto, errJsonMarshal
	}
	errJsonUnmarshal := json.Unmarshal(jsonDataBytes, dh.Dto)
	if errJsonUnmarshal != nil {
		return dto, errJsonUnmarshal
	}
	dto = *dh.Dto
	errorCustomMapping := dto.customMappingFromEntity(*dh.Entity, dh.Env.GiveDbHandler().PassConnection())
	if errorCustomMapping != nil {
		return dto, errorCustomMapping
	}
	return dto, nil
}

// ----------------Multi handler----------- \\

type IDTOHandlerMulti[E entities.Entity, D DTO[E]] interface {
	MapFromEntities() error
}

type DTOHandlerMulti[E entities.Entity, D DTO[E]] struct {
	Dtos     []D
	Entities []E
	Env      environment.Environment
}

func NewDtoHandlerMulti[E entities.Entity, D DTO[E]](entities []E, dtos []D, env environment.Environment) *DTOHandlerMulti[E, D] {
	return &DTOHandlerMulti[E, D]{
		Dtos:     dtos,
		Entities: entities,
		Env:      env,
	}
}

func (dh *DTOHandlerMulti[E, D]) MapFromEntities() error {
	for _, entity := range dh.Entities {
		var ptrDto D
		jsonDataBytes, errJsonMarshal := json.Marshal(entity)
		if errJsonMarshal != nil {
			return errJsonMarshal
		}
		errJsonUnmarshal := json.Unmarshal(jsonDataBytes, &ptrDto)
		if errJsonUnmarshal != nil {
			return errJsonUnmarshal
		}
		for _, d := range dh.Dtos {
			errorCustomMapping := d.customMappingFromEntity(entity, dh.Env.GiveDbHandler().PassConnection())
			if errorCustomMapping != nil {
				return errorCustomMapping
			}
		}
		dh.Dtos = append(dh.Dtos, ptrDto)
	}
	return nil

}
