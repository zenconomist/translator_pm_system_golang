package entities_test

import (
	dbc "dbconn"
	"entities"
	envir "environment"

	"fmt"

	"testing"
)

var env envir.Env
var dbh dbc.TestPostgreDb // test/prod is here to decide

func TestFindAllSpFoldersWithSqlFile(t *testing.T) {
	env.EnvFactory(&dbh)
	cr := entities.NewCustomRepo(env.DbHandler.PassConnection(), entities.SharePointFolder{})
	spfs, err := cr.FindAllSpFoldersToCreate()
	if err != nil {
		t.Errorf("couldn't run find all sp folders: %v", err)
	}
	fmt.Println(spfs)
}

/*
func TestTurnStateTypeToQueryStr(t *testing.T) {
	queryStr := turnStateTypeToQueryStr("taskstate")
	if queryStr != "task_state_id" {
		t.Errorf("func not working")
	}
	fmt.Println(queryStr)
}

func turnStateTypeToQueryStr(stateType string) string {
	// todo: extend for all states
	switch strings.ToLower(stateType) {
	case "taskstate", "task":
		return "task_state_id"
	case "tasktimestate", "tasktime":
		return "task_time_state_id"
	case "projectstate", "project":
		return "project_state_id"
	case "projecttimestate", "projecttime":
		return "project_time_id"
	case "batchstate", "batch":
		return "batch_id"
	case "batchtimestate", "batchtime":
		return "batch_time_state_id"
	case "clientofferstate", "clientoffer":
		return "client_offer_state_id"
	default:
		return ""
	}

}

// =====================FIRM================== \\


func TestCreateFirm(t *testing.T) {
	var env envir.Env
	var dbh dbc.TestSqlServerDb // test/prod is here to decide
	env.EnvFactory(&dbh)

	id := uuid.New().String()
	entity := entities.Firm{
		Name: "test" + id,
		FirmAddress: entities.Address{
			CountryCode: "HU",
			PostCode:    "1111",
			City:        "Budapest",
			Address:     "Ház u. 1",
		},
		MainEmail: "test@test" + id + ".com",
	}
	repo := entities.NewRepository(dbh.PassConnection(), entity)
	ID, err := repo.Create(entity)
	if err != nil {
		t.Errorf("not working %v", err)
	}
	fmt.Println("entity created with ID: ", ID)

}
// */
// =====================CUSTOMER================== \\
/*
func TestFindAllCustomer(t *testing.T) {
	env.EnvFactory(&dbh)

	var entity entities.Customer
	repo := entities.NewRepository(dbh.PassConnection(), entity)
	entities, err := repo.FindAll()
	if err != nil {
		t.Errorf("not working %v", err)
	}
	fmt.Println(entities)

}

func TestFindByIDCustomer(t *testing.T) {
	env.EnvFactory(&dbh)

	var entity entities.Customer
	repo := entities.NewRepository(dbh.PassConnection(), entity)
	entity, err := repo.FindByID(uint(4))
	if err != nil {
		t.Errorf("not working %v", err)
	}
	fmt.Println(entity)

}


func TestFindAllWithIDsCustomer(t *testing.T) {
	env.EnvFactory(&dbh)

	var entity entities.Customer
	repo := entities.NewRepository(dbh.PassConnection(), entity)
	ids := []uint{3, 4, 5}
	entities_, err := repo.FindAllWithIDs(ids)
	if err != nil {
		t.Errorf("not working %v", err)
	}
	fmt.Println(entities_)

}

// */

/*
func TestCreateCustomer(t *testing.T) {
	env.EnvFactory(&dbh)
	db := env.DbHandler.PassConnection()

	id := uuid.New().String()
	entity := entities.Customer{
		Name: "test" + id,
		Address: entities.Address{
			// CustomerID:  1,
			CountryCode: "HU",
			PostCode:    "1111",
			City:        "Budapest",
			Address:     "Ház u. 1",
		},
		BillPID:             111111111111111,
		Email:               "test@test" + id + ".com",
		TaxCode:             "1111111111",
		Iban:                "11111111111111111111111111111",
		Swift:               "111111111111111111111111111111",
		AccountNumber:       "1111111111-1111111-1111111",
		Phone:               "1111111111111",
		GeneralLedgerNumber: "dawdawdawfaegse",
		TaxType:             "ev.",
		CustGeneralInfo:     "is this a test? It has to be... Otherwise I can't go on.",
		Currency:            1,
		FirmID:              1,
		InvoiceLang:         1,
		PaymentDueDays:      34,
	}
	repo := entities.NewRepository(db, entity)
	newID, err := repo.Create(entity)
	if err != nil {
		t.Errorf("not working %v", err)
	}
	fmt.Println("entity created.")

	// historicization
	var history entities.CustomerHistory
	hrepo := entities.NewHistoryRepo(db, &history, entities.NewHistoryMapper(&history, entity), entity)
	hist, errMap := hrepo.HistoryMapper.MapToHistoryForCreation(entity)
	if errMap != nil {
		t.Errorf("not working %v", errMap)
	}
	hist.SetModelToCreated(newID)
	hrepo.Create(&hist)
}

/*
func TestUpdateCustomer(t *testing.T) {
	env.EnvFactory(&dbh)

	var entity entities.Customer
	db := env.DbHandler.PassConnection()
	result := db.Last(&entity)
	if result.Error != nil {
		t.Errorf("couldn't query added customer, %v", result.Error)
	}
	repo := entities.NewRepository(dbh.PassConnection(), entity)
	db.Last(&entity)

	entity.AccountNumber = "N/A"
	entity.GeneralLedgerNumber = "N/A"

	err := repo.Update(entity)
	if err != nil {
		t.Errorf("not working %v", err)
	}

	fmt.Println("entity updated.")

	// historicization
	var history entities.CustomerHistory
	hrepo := entities.NewHistoryRepo(db, &history, entities.NewHistoryMapper(&history, entity), entity)
	hist, errMap := hrepo.HistoryMapper.MapToHistoryForCreation(entity)
	if errMap != nil {
		t.Errorf("not working %v", errMap)
	}
	hist.SetModelToUpdated(entity.GiveID())
	errUpdate := hrepo.Update(hist, entity)
	if errUpdate != nil {
		t.Errorf("update failed %v", errUpdate)
	}
	hist.SetModelToCreated(entity.GiveID())
	hrepo.Create(&hist)
}

func TestDeleteCustomer(t *testing.T) {
	env.EnvFactory(&dbh)

	var entity entities.Customer
	db := env.DbHandler.PassConnection()
	result := db.Last(&entity)
	if result.Error != nil {
		t.Errorf("couldn't query added customer, %v", result.Error)
	}
	repo := entities.NewRepository(dbh.PassConnection(), entity)
	err := repo.Delete(entity)
	if err != nil {
		t.Errorf("not working %v", err)
	}
	fmt.Println("entity deleted.")

	// historicization
	var history entities.CustomerHistory
	hrepo := entities.NewHistoryRepo(db, &history, entities.NewHistoryMapper(&history, entity), entity)
	hist, errMap := hrepo.HistoryMapper.MapToHistoryForCreation(entity)
	if errMap != nil {
		t.Errorf("not working %v", errMap)
	}
	hist.SetModelToDeleted(entity.GiveID())
	errUpdate := hrepo.Update(hist, entity)
	if errUpdate != nil {
		t.Errorf("delete history failed %v", errUpdate)
	}
}


// =====================USER================== \\

func TestFindAllUser(t *testing.T) {
	env.EnvFactory(&dbh)

	var entity entities.User
	repo := entities.NewRepository(dbh.PassConnection(), entity)
	entities, err := repo.FindAll()
	if err != nil {
		t.Errorf("not working %v", err)
	}
	fmt.Println(entities)

}

func TestFindByIDUser(t *testing.T) {
	env.EnvFactory(&dbh)

	var entity entities.User
	repo := entities.NewRepository(dbh.PassConnection(), entity)
	entity, err := repo.FindByID(uint(3))
	if err != nil {
		t.Errorf("not working %v", err)
	}
	fmt.Println(entity)

}

func TestFindAllWithIDsUser(t *testing.T) {
	env.EnvFactory(&dbh)

	var entity entities.User
	repo := entities.NewRepository(dbh.PassConnection(), entity)
	entity, err := repo.FindByID(uint(3))
	if err != nil {
		t.Errorf("not working %v", err)
	}
	fmt.Println(entity)

}

func TestCreateUser(t *testing.T) {
	env.EnvFactory(&dbh)

	id := uuid.New().String()
	entity := entities.User{
		Name:     "test",
		UserName: "test" + id,
		Email:    "test@test" + id + ".com",
		Roles: []entities.Role{entities.Role{
			RoleName: "admin",
		},
		},
		Active: true,
		Car:    true,
		Address: entities.Address{
			// CustomerID:  1,
			CountryCode: "HU",
			PostCode:    "2222",
			City:        "Valahol",
			Address:     "Nemtom u. 2/b.",
		},
		BillingCurrency: 2,
		Prices: []entities.SupplierPrice{
			entities.SupplierPrice{
				TaskTypeID: 0,
				SourceLang: "HU",
				TargetLang: "EN",
				Price:      0.61,
			}},
		Specialities: "orvosi, műszaki",
		Languages:    "HU->EN, EN->HU",
	}
	repo := entities.NewRepository(dbh.PassConnection(), entity)
	ID, err := repo.Create(entity)
	if err != nil {
		t.Errorf("not working %v", err)
	}
	fmt.Println("entity created with ID: ", ID)

}
*/
/*
func TestUpdateUser(t *testing.T) {
	env.EnvFactory(&dbh)

	var entity entities.User
	db := env.DbHandler.PassConnection()
	result := db.Last(&entity)
	if result.Error != nil {
		t.Errorf("couldn't query added customer, %v", result.Error)
	}
	repo := entities.NewRepository(db, entity)
	entity.BillingCurrency = 0
	entity.Specialities = "N/A"
	err := repo.Update(entity)
	if err != nil {
		t.Errorf("not working %v", err)
	}
	fmt.Println("entity updated.")

}

func TestDeleteUser(t *testing.T) {
	env.EnvFactory(&dbh)

	var entity entities.User
	db := env.DbHandler.PassConnection()
	result := db.Last(&entity)
	if result.Error != nil {
		t.Errorf("couldn't query added customer, %v", result.Error)
	}
	repo := entities.NewRepository(db, entity)
	err := repo.Delete(entity)
	if err != nil {
		t.Errorf("not working %v", err)
	}
	fmt.Println("entity deleted.")

}


func TestCheckIfStatesAreValid(t *testing.T) {
	testStates1 := []int{1, 2, 3, 4}
	testStates2 := []int{1, 2, 3, 4, 16}

	var ts globalconstants.TaskState
	errCheck1 := ts.CheckIfStatesAreValid(testStates1)
	if errCheck1 != nil {
		t.Errorf("not valid states, error: %v", errCheck1)
	}
	errCheck2 := ts.CheckIfStatesAreValid(testStates2)
	if errCheck2 == nil {
		t.Errorf("invalid states validated, error: %v", errCheck2)
	}
}

// */

/*
func TestTaskStateSetting(t *testing.T) {
	defer recoverFromPanic()
	env.EnvFactory(&dbh)
	var task entities.Task
	env.GiveDbHandler().PassConnection().Last(&task)
	fmt.Println("here starts the state-setting")
	task.TaskState = &entities.TaskStateOpen{
		Task: &task,
		Tsc:  entities.TaskStOpen,
	}
	var err error
	err = task.TaskState.SetTaskStateToOpen()
	if err == nil {
		t.Errorf("couldn't set task to open: %v", err)
	}
	err = task.TaskState.SetTaskStateToOffered()
	if err != nil {
		t.Errorf("couldn't set task to offered: %v", err)
	} else {
		if task.TaskState.GiveName() != entities.TaskStOffered.String() {
			t.Errorf("couldn't set task to offered: %v", err)
		}
	}

	fmt.Println("task's current state: ", task.TaskState.GiveName())
}
*/

func recoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println("recovery, ", r)
	}
}
