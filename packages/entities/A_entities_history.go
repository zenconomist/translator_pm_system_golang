package entities

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"gorm.io/gorm"
)

type HistoryModel struct {
	ID           uint           `json:"history_id" gorm:"primaryKey;autoIncrement"`
	OriginalID   uint           `json:"id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	FlgIsCurrent bool
	FlgIsDeleted bool
	DatFrom      time.Time
	DatTo        time.Time
}

type IHistoryMapper[H EntityHistory, E Entity] interface {
	MapToHistoryForCreation(E) (H, error)
}

type HistoryMapper[H EntityHistory, E Entity] struct {
	EntityHistory H
	Entity        E
}

func NewHistoryMapper[H EntityHistory, E Entity](history H, entity E) HistoryMapper[H, E] {
	return HistoryMapper[H, E]{
		EntityHistory: history,
		Entity:        entity,
	}
}

func (hm HistoryMapper[H, E]) MapToHistoryForCreation(entity E) (H, error) {
	var history H
	jsonDtoBytes, errJsonMarshal := json.Marshal(entity)
	if errJsonMarshal != nil {
		return history, errJsonMarshal
	}
	errJsonUnmarshal := json.Unmarshal(jsonDtoBytes, &history)
	if errJsonUnmarshal != nil {
		return history, errJsonUnmarshal
	}

	return history, nil
}

// -------------- REPO --------------- \\

type IEntityHistory interface {
	GiveID() uint
}

type EntityHistory interface {
	*BatchHistory | *ClientOfferHistory | *ClientOfferTaskHistory | *CustomerHistory | *FirmHistory | *ProjectHistory | *TaskHistory | *TaskConfigHistory | *TaskOfferedHistory | *UserHistory | *BillingLogHistory | *UPMLoggerHistory | *DefaultCustomerPriceHistory | *DefaultSupplierPriceHistory | *CustomerPriceHistory | *SupplierPriceHistory | *ContactHistory
	GiveID() uint
	SetModelToCreated(uint)
	SetModelToUpdated(uint)
	SetModelToDeleted(uint)
}

type IHistoryRepo[H EntityHistory] interface {
	FindAll() ([]H, error)
	FindAllWithIDs([]uint) ([]H, error)
	FindByID(uint) (H, error)
	Create(H) error
	Update(H) error
	Delete(H) error
}

type HistoryRepo[H EntityHistory, HM HistoryMapper[H, E], E Entity] struct {
	EntityHistory H
	HistoryMapper HM
	Entity        E
	db            *gorm.DB
}

func NewHistoryRepo[H EntityHistory, HM HistoryMapper[H, E], E Entity](db *gorm.DB, history H, historyMapper HM, entity E) *HistoryRepo[H, HM, E] {
	return &HistoryRepo[H, HM, E]{
		db:            db,
		EntityHistory: history,
		HistoryMapper: historyMapper,
		Entity:        entity,
	}
}

func (repo HistoryRepo[H, HM, E]) FindAll() ([]H, error) {
	var histories []H
	var history H
	queryString, params := repo.queryBuilder(history)
	result := repo.db.Where(queryString, params).Find(&histories)
	if result.RowsAffected == 0 {
		return histories, fmt.Errorf("no results")
	}
	return histories, result.Error
}

func (repo HistoryRepo[H, HM, E]) FindAllWithIDs(ids []uint) ([]H, error) {
	var histories []H
	var history H
	queryString, params := repo.queryBuilder(history)
	result := repo.db.Where(queryString, params).Find(&histories, ids)
	return histories, result.Error
}

func (repo HistoryRepo[H, HM, E]) FindByID(id uint) (H, error) {
	var history H
	queryString, params := repo.queryBuilder(history)
	result := repo.db.Where(queryString, params).First(&history, id)
	return history, result.Error
}

func (repo HistoryRepo[H, HM, E]) Create(history *H) error {
	result := repo.db.Omit("id").Create(&history)
	return result.Error
}

func (repo HistoryRepo[H, HM, E]) Update(history H, entity E) error {
	result := repo.db.Model(&history).Where("flg_is_current = ? AND flg_is_deleted = ? AND original_id = ?", true, false, entity.GiveID()).Select("flg_is_current", "flg_is_deleted", "dat_to").Updates(&history)
	return result.Error
}

func (repo HistoryRepo[H, HM, E]) Delete(history H) error {
	result := repo.db.Delete(&history, history.GiveID())
	return result.Error
}

func (repo HistoryRepo[H, HM, E]) queryBuilder(history H) (string, []interface{}) {
	// queryBuilder
	metaValue := reflect.ValueOf(history)
	field := metaValue.FieldByName("Active")
	queryString := strings.Builder{}
	queryString.WriteString("1=1")
	var params []interface{}
	if field.IsValid() {
		queryString.WriteString(" AND active = ? ")
		params = append(params, true)
	}
	return queryString.String(), params
}
