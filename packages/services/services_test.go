package services_test

import (
	dbc "dbconn"
	"environment"
	msgr "msgraph"
	"services"
	"testing"
)

var env environment.Env
var dbh dbc.TestPostgreDb

func TestMockMsGraphClient(t *testing.T) {
	env.EnvFactory(&dbh)
	var tenantid = "5e8275ce-0b52-4bf2-9de2-d8f152c7867f"   //os.Getenv("MSGRAPH_TENANTID")
	var appid = "bdfc7797-5d28-4124-a155-3d7b9eba5c8d"      //os.Getenv("MSGRAPH_APPID")
	var clientsecret = "VtTIL~9PJCcd5H-dks66Fy1LI~D.w126lS" //os.Getenv("MSGRAPH_CLIENTSECRET")
	// var driveID = "b!KghC8XlmgUWVFI2kpGEGQgq9LOV2XBpKtfWQvdvMB2CTiXyN7H7_Rqurz9mBpom-" //os.Getenv("MSGRAPH_DRIVEID")

	// add MsGraph folders
	gcl, errgcl := msgr.NewGraphClient(tenantid, appid, clientsecret)
	if errgcl != nil {
		t.Errorf("unable to establish msgraph connection: %v", errgcl)
	}
	mgs := services.NewMsGraphService(&env, gcl)
	if errCreateFolders := mgs.CreateProjectFolders(); errCreateFolders != nil {
		t.Errorf("unable to create folders: %v", errCreateFolders)
	}
}

// /*

// */
/*
func TestGetSqlStringFromFile(t *testing.T) {
	SQLbytes, err := ioutil.ReadFile("/db/prod/customqueries/spfolders_to_create.sql")
	if err != nil {
		t.Errorf("couldn't read file: %v", err)
	}
	fmt.Println(string(SQLbytes))

}




func TestDeleteAllTasks(t *testing.T) {
	env.EnvFactory(&dbh)
	service := services.NewService(entities.NewRepository(env.DbHandler.PassConnection(), entities.Task{}), &env, &dto.TaskRequestDTO{}, entities.Task{}, &entities.TaskHistory{})
	tasks, errGet := service.GetItems()
	if errGet != nil {
		t.Errorf("couldn't get project dto-s: %v", errGet)
	}
	for _, task := range tasks {
		if errDelete := service.DeleteItem(task.ID); errDelete != nil {
			t.Errorf("couldn't delete items: %v", errDelete)
		}
	}

}


func TestDeleteAllProjects(t *testing.T) {
	env.EnvFactory(&dbh)
	service := services.NewService(entities.NewRepository(env.DbHandler.PassConnection(), entities.Project{}), &env, &dto.ProjectRequestDTO{}, entities.Project{}, &entities.ProjectHistory{})
	projectDtos, errGet := service.GetItems()
	if errGet != nil {
		t.Errorf("couldn't get project dto-s: %v", errGet)
	}
	for _, p := range projectDtos {
		if errDelete := service.DeleteItem(p.ID); errDelete != nil {
			t.Errorf("couldn't delete items: %v", errDelete)
		}
	}

}
*/
/*
func TestGetAllFoldersToCreate(t *testing.T) {
	env.EnvFactory(&dbh)
	var tenantid = "5e8275ce-0b52-4bf2-9de2-d8f152c7867f"   //os.Getenv("MSGRAPH_TENANTID")
	var appid = "bdfc7797-5d28-4124-a155-3d7b9eba5c8d"      //os.Getenv("MSGRAPH_APPID")
	var clientsecret = "VtTIL~9PJCcd5H-dks66Fy1LI~D.w126lS" //os.Getenv("MSGRAPH_CLIENTSECRET")
	// var driveID = "b!KghC8XlmgUWVFI2kpGEGQgq9LOV2XBpKtfWQvdvMB2CTiXyN7H7_Rqurz9mBpom-" //os.Getenv("MSGRAPH_DRIVEID")
	// add MsGraph folders
	gcl, errgcl := msgr.NewGraphClient(tenantid, appid, clientsecret)
	if errgcl != nil {
		t.Errorf("unable to establish msgraph connection: %v", errgcl)
	}
	mgs := services.NewMsGraphService(&env, gcl)
	// select all folders that have to be created
	cr := entities.NewCustomRepo(mgs.Env.GiveDbHandler().PassConnection(), entities.SharePointFolder{})
	spfs, errFindSpFolders := cr.FindAllSpFoldersToCreate()
	if errFindSpFolders != nil {
		fmt.Errorf("cannot query folders: %v", errFindSpFolders)
	}
	// call CreateFldr for all of these folders
	for _, f := range spfs {
		fmt.Println(f)
	}
}


func TestGetProjectWithTasks(t *testing.T) {
	s := time.Now()
	env.EnvFactory(&dbh)
	cs := services.NewCustomService(entities.NewCustomRepo(env.GiveDbHandler().PassConnection(), entities.Project{}))
	projectDtos, err := cs.GetAllProjectsAndTasks()
	if err != nil {
		t.Errorf("service error: %v", err)
	}
	fmt.Println(projectDtos)
	fmt.Printf("projects and tasks queried: @ %v", time.Since(s))
}


func TestNewService(t *testing.T) {
	var env envir.Env
	var dbh dbc.TestSqlServerDb // test/prod is here to decide
	env.EnvFactory(&dbh)

	id := uuid.New().String()
	dto := dto.CustomerRequestDTO{
		Name: "test" + id,
		Address: dto.AddressRequestDTO{
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
		CurrencyName:        "HUF",
		DefaultFirm:         1,
		InvoiceLangName:     "en",
		PaymentDueDays:      34,
	}
	var entity entities.Customer
	var history entities.CustomerHistory
	repo := entities.NewRepository(env.GiveDbHandler().PassConnection(), entity)
	service := services.NewService(repo, &env, &dto, entity, &history)
	service.EmptyService()
	service.CreateNewItem(&dto)
}

func TestUpdateCustomerService(t *testing.T) {
	var env envir.Env
	var dbh dbc.TestSqlServerDb // test/prod is here to decide
	env.EnvFactory(&dbh)

	var entity entities.Customer
	id := uuid.New().String()
	dbh.PassConnection().Last(&entity)
	dto := dto.CustomerRequestDTO{
		ID:   entity.Model.ID,
		Name: "test" + id,
		Address: dto.AddressRequestDTO{
			// CustomerID:  1,
			CountryCode: "HU",
			PostCode:    "1222",
			City:        "Budapest",
			Address:     "Szög u. 2",
		},
		BillPID:             22222222222,
		Email:               "test@test" + id + ".com",
		TaxCode:             "2",
		Iban:                "2",
		Swift:               "2",
		AccountNumber:       "2",
		Phone:               "2",
		GeneralLedgerNumber: "féwlkmf",
		TaxType:             "nem katás",
		CustGeneralInfo:     "this is a test",
		CurrencyName:        "EUR",
		DefaultFirm:         1,
		InvoiceLangName:     "hu",
		PaymentDueDays:      76,
	}
	var history entities.CustomerHistory
	repo := entities.NewRepository(env.GiveDbHandler().PassConnection(), entity)
	service := services.NewService(repo, &env, &dto, entity, &history)
	service.EmptyService()
	errUpdate := service.UpdateItem(&dto)
	if errUpdate != nil {
		t.Errorf(errUpdate.Error())
	}
}


func TestDeleteCustomerService(t *testing.T) {
	var env envir.Env
	var dbh dbc.TestSqlServerDb // test/prod is here to decide
	env.EnvFactory(&dbh)

	var entity entities.Customer
	dbh.PassConnection().Last(&entity)
	var dto dto.CustomerRequestDTO
	var history entities.CustomerHistory
	repo := entities.NewRepository(env.GiveDbHandler().PassConnection(), entity)
	service := services.NewService(repo, &env, &dto, entity, &history)
	service.EmptyService()
	errDel := service.DeleteItem(entity.Model.ID)
	if errDel != nil {
		t.Errorf(errDel.Error())
	}
}

// ======================PROJECT====================\\

func TestNewProjectService(t *testing.T) {
	var env envir.Env
	var dbh dbc.TestSqlServerDb // test/prod is here to decide
	env.EnvFactory(&dbh)

}


*/
// ======================TASKS====================\\
/*
func TestNewTaskService(t *testing.T) {
	var env envir.Env
	var dbh dbc.TestSqlServerDb // test/prod is here to decide
	env.EnvFactory(&dbh)

	// id := uuid.New().String()
	taskState := globalconstants.TaskOpen
	state := globalconstants.TaskState{State: taskState}
	state.SetTaskStateName()
	dto := dto.TaskUpsertRequestDTO{
		ProjectID:        1,
		BatchID:          0,
		OrderWithinBatch: 1,
		ProjectManager:   1,
		TaskState:        state.GiveTaskStateName(),
		TaskStateID:      state.GiveTaskStateUint(),
	}
	var entity entities.Task
	var history entities.TaskHistory
	repo := entities.NewRepository(env.GiveDbHandler().PassConnection(), entity)
	service := services.NewService(repo, &env, &dto, entity, &history)
	service.EmptyService()
	service.CreateNewItem(&dto)
}
func TestUpdateTaskService(t *testing.T) {
	var env envir.Env
	var dbh dbc.TestSqlServerDb // test/prod is here to decide
	env.EnvFactory(&dbh)

	var entity entities.Task
	dbh.PassConnection().Last(&entity)
	taskState := entities.TaskStPending
	state := entity.TaskState{State: taskState}
	state.SetTaskStateName()
	dto := dto.TaskUpsertRequestDTO{
		ID:               entity.Model.ID,
		ProjectID:        1,
		BatchID:          0,
		OrderWithinBatch: 1,
		ProjectManager:   1,
		TaskState:        state.GiveTaskStateName(),
		TaskStateID:      state.GiveTaskStateUint(),
	}
	var history entities.TaskHistory
	repo := entities.NewRepository(env.GiveDbHandler().PassConnection(), entity)
	service := services.NewService(repo, &env, &dto, entity, &history)
	service.EmptyService()
	errUpdate := service.UpdateItem(&dto)
	if errUpdate != nil {
		t.Errorf(errUpdate.Error())
	}
}
*/
/*
func TestDeleteTaskService(t *testing.T) {
	var env envir.Env
	var dbh dbc.TestSqlServerDb // test/prod is here to decide
	env.EnvFactory(&dbh)

	var entity entities.Customer
	dbh.PassConnection().Last(&entity)
	var dto dto.CustomerUpsertRequestDTO
	var history entities.CustomerHistory
	repo := entities.NewRepository(env.GiveDbHandler().PassConnection(), entity)
	service := services.NewService(repo, &env, &dto, entity, &history)
	service.EmptyService()
	errDel := service.DeleteItem(entity.Model.ID)
	if errDel != nil {
		t.Errorf(errDel.Error())
	}
}



*/
