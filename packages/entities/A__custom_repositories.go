package entities

import (
	"readfiles"

	"gorm.io/gorm"
)

type ICustomRepo interface {
	FindAllAssociatedTask(Project) ([]Task, error)
}

type CustomRepo struct {
	Db       *gorm.DB
	Entity   IEntity
	Entities []IEntity
}

func NewCustomRepo(db *gorm.DB, entity IEntity) *CustomRepo {
	return &CustomRepo{
		Db:     db,
		Entity: entity,
	}
}

func (cr *CustomRepo) FindAllAssociatedTask(project Project) ([]Task, error) {
	var tasks []Task
	cr.Db.Model(&project).Association("Tasks").Find(&tasks)
	return tasks, nil

}

func (cr *CustomRepo) FindAllSpFoldersToCreate() ([]SharePointFolder, error) {
	var spfs []SharePointFolder
	SQLbytes, errReadSql := readfiles.GetFileContents("spfolders_to_create.sql")
	if errReadSql != nil {
		return spfs, errReadSql
	}
	result := cr.Db.Raw(string(SQLbytes)).Scan(&spfs)
	if result.Error != nil {
		return spfs, result.Error
	}
	return spfs, nil
}

func (cr *CustomRepo) GetParentItem(projectID uint, configID int) (SharePointFolder, error) {
	var spf SharePointFolder
	result := cr.Db.Model(&SharePointFolder{}).Where("project_id = ? and config_id = ?", projectID, configID).Scan(&spf)
	if result.Error != nil {
		return spf, result.Error
	}
	return spf, nil
}

func (cr *CustomRepo) GetTaskOfferToValidate(taskID, supplierID uint) bool {
	taskOfferExists := false
	var taskOffer TaskOffered
	result := cr.Db.Where("task_id = ? and supplier_id = ?", taskID, supplierID).First(&taskOffer)
	if result.RowsAffected > 0 {
		taskOfferExists = true
	}
	return taskOfferExists
}

func (cr *CustomRepo) TaskHistoricization() error {
	SQLbytes, errReadSql := readfiles.GetFileContents("tasks_historicization.sql")
	if errReadSql != nil {
		return errReadSql
	}
	result := cr.Db.Exec(string(SQLbytes))
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (cr *CustomRepo) FindOfferedTask(id uint) (*Task, error) {
	var task Task

	// cr.Db.Where()
	// invalidTsc := TaskStInvalid
	// offeredStateID, errState := invalidTsc.StringToState("Offered")
	// if errState != nil {
	// 	return &task, fmt.Errorf("task offering state's id couldn't be retreived")
	// }
	// result := cr.Db.Where("task_state_id = ? and id = ?", uint(offeredStateID), id).First(&task)
	// if result.Error != nil {
	// 	return &task, result.Error
	// }
	return &task, nil
}

func (cr *CustomRepo) FindAnyOtherTaskOfferings(taskID uint) bool {
	var tos []TaskOffered
	result := cr.Db.Where("task_id = ?", taskID).Find(&tos)
	if result.Error != nil {
		return false
	}
	return len(tos) > 0
}
