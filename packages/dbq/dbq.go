package dbq

import (
	"entities"
	enviro "environment"
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

const file = "./codegenerator/UPM_v0.2_DataModel_v1.xlsx"

func EmptyDbq() {}

func Migrate(env enviro.Environment) {
	db := env.GiveDbHandler().PassConnection()
	db.AutoMigrate(&entities.UPMLogger{}) // ok
	// db.AutoMigrate(&entities.Address{})  // embedded
	db.AutoMigrate(&entities.Batch{})
	db.AutoMigrate(&entities.BillingLog{})
	db.AutoMigrate(&entities.ClientOffer{})
	db.AutoMigrate(&entities.ClientOfferTask{})
	db.AutoMigrate(&entities.CustomerPrice{})
	db.AutoMigrate(&entities.Customer{})
	// db.AutoMigrate(&entities.Email{}) // not relevant in this way to persist
	db.AutoMigrate(&entities.Firm{}) // ok
	// db.AutoMigrate(&entities.FirmAddress{}) // embedded
	db.AutoMigrate(&entities.Permission{}) // ok
	db.AutoMigrate(&entities.SupplierPrice{})
	db.AutoMigrate(&entities.TaskConfig{})
	// db.AutoMigrate(&entities.TaskCustomerProps{}) // embedded
	// db.AutoMigrate(&entities.TaskSupplierProps{}) // embedded
	db.AutoMigrate(&entities.Role{}) // ok
	db.AutoMigrate(&entities.User{})

	db.AutoMigrate(&entities.TaskOffered{})
	db.AutoMigrate(&entities.Contact{})
	db.AutoMigrate(&entities.SharePointFolder{})
	db.AutoMigrate(&entities.SharePointFolderConfig{})
	db.AutoMigrate(&entities.EmailSendingLog{})
	db.AutoMigrate(&entities.TaskStateConfigHead{})
	db.AutoMigrate(&entities.TaskState{})
	db.AutoMigrate(&entities.TaskStateChanges{})
	db.AutoMigrate(&entities.TaskTimeStateConfigHead{})
	db.AutoMigrate(&entities.TaskTimeState{})
	db.AutoMigrate(&entities.TaskTimeStateChanges{})
	db.AutoMigrate(&entities.EmailNotifications{})
	db.AutoMigrate(&entities.ActualStateChanges{})

	db.AutoMigrate(&entities.Task{})

	db.AutoMigrate(&entities.Project{})

	// MigrateSpecial(env)
	fmt.Println("migrated all")
}

func MigrateHistory(env enviro.Environment) {
	db := env.GiveDbHandler().PassConnection()
	db.AutoMigrate(&entities.UPMLoggerHistory{}) // ok
	db.AutoMigrate(&entities.BatchHistory{})
	db.AutoMigrate(&entities.BillingLogHistory{})
	db.AutoMigrate(&entities.ClientOfferHistory{})
	db.AutoMigrate(&entities.ClientOfferTaskHistory{})
	db.AutoMigrate(&entities.CustomerHistory{})
	db.AutoMigrate(&entities.FirmHistory{}) // ok
	db.AutoMigrate(&entities.TaskConfigHistory{})
	db.AutoMigrate(&entities.UserHistory{})
	db.AutoMigrate(&entities.TaskOfferedHistory{})
	db.AutoMigrate(&entities.ContactHistory{})
	db.AutoMigrate(&entities.SharePointFolderHistory{})
	db.AutoMigrate(&entities.SharePointFolderConfigHistory{})
	db.AutoMigrate(&entities.TaskHistory{})
	db.AutoMigrate(&entities.ProjectHistory{})
	db.AutoMigrate(&entities.ActualStateChangesHistory{})

	fmt.Println("migrated all history")
}

func MigrateSpecial(env enviro.Environment) {
	db := env.GiveDbHandler().PassConnection()
	// db.AutoMigrate(&entities.CustomerPrice{})
	db.AutoMigrate(&entities.Task{})
	db.AutoMigrate(&entities.Project{})

}

func UpdateActualStatesAndChanges(env enviro.Environment) error {
	var asc entities.ActualStateChanges
	db := env.GiveDbHandler().PassConnection()
	db.Exec("TRUNCATE TABLE actual_state_changes CASCADE;")

	f, err := excelize.OpenFile(file)
	if err != nil {
		fmt.Println(err)
	}

	sheet := "TaskState_StateChanges"
	rowsDT := f.GetRows(sheet) // get sheet named after a distinct datatype
	for j, row := range rowsDT {
		if j == 0 {
			// excel header to be skipped
			continue
		}

		asc.From = row[0]
		asc.To = row[1]
		asc.IsAllowed = convertFromStringToBool(row[2])
		asc.StateChangeInfo = row[3]
		asc.NeedsComment = convertFromStringToBool(row[4])
		asc.PMPerform = convertFromStringToBool(row[5])
		asc.PreparerPerfrom = convertFromStringToBool(row[6])
		asc.ReviewerPerform = convertFromStringToBool(row[7])
		asc.SupplierPerform = convertFromStringToBool(row[8])
		asc.PMNotified = convertFromStringToBool(row[9])
		asc.PreparerNotified = convertFromStringToBool(row[10])
		asc.ReviewerNotified = convertFromStringToBool(row[11])
		asc.SupplierNotified = convertFromStringToBool(row[12])
		asc.ConditionOrExplanation = row[13]

		result := db.Omit("id").Create(&asc)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

func ExecStateChangeUpdateScript(env enviro.Environment) error {
	db := env.GiveDbHandler().PassConnection()
	result := db.Exec("call usp_task_state_change_load();")
	if result.Error != nil {
		return result.Error
	}
	testresult := db.Exec("call usp_task_state_change_testing();")
	if testresult.Error != nil {
		return testresult.Error
	}

	return nil
}

func convertFromStringToBool(a string) bool {
	if a == "True" {
		return true
	} else {
		return false
	}
}
