package entities

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"gorm.io/gorm"
	// "golang.org/x/exp/constraints"
)

type Model struct {
	ID        uint           `gorm:"primaryKey;autoIncrement"` // works both with postgres and SQL Server
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type IEntity interface {
	GiveID() uint
}

type Entity interface {
	Address | Batch | BillingLog | ClientOffer | ClientOfferTask | ClientOfferTaskCustomerProps | Customer | CustomerPrice | Firm | Permission | Project | Role | SupplierPrice | Task | TaskConfig | TaskCustomerProps | TaskSupplierProps | TaskOffered | UPMLogger | User | DefaultCustomerPrice | DefaultSupplierPrice | Contact | EmailSendingLog
	GiveID() uint
}

type IRepository[E Entity] interface {
	FindAll() ([]E, error)
	FindAllWithIDs([]uint) ([]E, error)
	FindByID(uint) (E, error)
	FindLast() (E, error)
	Create(E) (uint, error)
	Update(E) error
	Delete(E) error
}

type Repository[E Entity] struct {
	Entity E
	db     *gorm.DB
}

func NewRepository[E Entity](db *gorm.DB, entity E) *Repository[E] {
	return &Repository[E]{
		db:     db,
		Entity: entity,
	}
}

func (repo Repository[E]) FindAll() ([]E, error) {
	var entities []E
	var entity E

	var result *gorm.DB
	queryString, params := repo.queryBuilder(entity)
	if queryString == "" {
		result = repo.db.Find(&entities)
	} else {
		result = repo.db.Where(queryString, params).Find(&entities)
	}
	if result.RowsAffected == 0 {
		return entities, fmt.Errorf("no results")
	}
	return entities, result.Error
}

func (repo Repository[E]) FindAllWithIDs(ids []uint) ([]E, error) {
	var entities []E
	var entity E

	queryString, params := repo.queryBuilder(entity)
	if queryString == "" {
		if result := repo.db.Find(&entities, ids); result.Error != nil {
			return entities, result.Error
		}
	} else {
		if result := repo.db.Where(queryString, params).Find(&entities, ids); result != nil {
			return entities, result.Error
		}
	}
	return entities, nil
}

func (repo Repository[E]) FindByID(id uint) (E, error) {
	var entity E

	var result *gorm.DB
	queryString, params := repo.queryBuilder(entity)
	if queryString == "" {
		result = repo.db.First(&entity, id)
	} else {
		result = repo.db.Where(queryString, params).First(&entity, id)
	}
	return entity, result.Error
}

func (repo Repository[E]) FindLast() (E, error) {
	result := repo.db.Last(&repo.Entity)
	if result.Error != nil {
		return repo.Entity, result.Error
	}
	return repo.Entity, nil
}

func (repo Repository[E]) Create(entity E) (uint, error) {
	result := repo.db.Omit("id").Create(&entity)
	return entity.GiveID(), result.Error
}

func (repo Repository[E]) Update(entity E) error {
	result := repo.db.Save(&entity)
	return result.Error
}

func (repo Repository[E]) Delete(entity E) error {
	result := repo.db.Delete(&entity, entity.GiveID())
	return result.Error
}

func (repo Repository[E]) queryBuilder(entity E) (string, []interface{}) {
	metaValue := reflect.ValueOf(entity)
	field := metaValue.FieldByName("Active")
	queryString := strings.Builder{}
	var params []interface{}
	if field.IsValid() {
		queryString.WriteString("active = ? ")
		params = append(params, true)
	}
	return queryString.String(), params
}
