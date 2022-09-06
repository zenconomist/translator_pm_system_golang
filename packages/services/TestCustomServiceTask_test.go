package services_test

// "services"

/*
func TestSetTaskState(t *testing.T) {
	var env environment.Environment
	var dbh dbc.TestSqlServerDb
	env.EnvFactory(&dbh)
	var task entities.Task
	dbh.PassConnection().Last(&task)

	rws := entities.NewRepoWithStates(dbh.PassConnection(), task, task.TaskState)
	tss := services.NewTaskStateService(task, task.TaskState, rws, dbh.PassConnection())
	if err := tss.SetTaskState("Open", &task); err != nil {
		t.Errorf("couldn't set to new state")
	}

}


func TestFindByStates(t *testing.T) {
	var env envir.Env
	var dbh dbc.TestSqlServerDb // test/prod is here to decide
	env.EnvFactory(&dbh)

	var task entities.Task
	var tasks []entities.Task
	fmt.Println("before repo instantiate")
	rws := entities.NewRepoWithStates(dbh.PassConnection(), task, task.TaskState)
	fmt.Println("after repo instantiate")
	states := make(map[string][]int)
	stateNums := []int{1, 3}
	states["taskstate"] = stateNums

	result := dbh.PassConnection().Where("task_state_id IN ?", stateNums).Find(&tasks)
	if result.Error != nil {
		fmt.Errorf("couldn't query data normally")
	}
	fmt.Println(tasks)

	fmt.Println("before find")
	tasks, err := rws.FindByStates(states)
	if err != nil {
		t.Errorf("couldn't find tasks, %v", err)
	}
	fmt.Println(tasks)
}
*/
