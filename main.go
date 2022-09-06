package main

import (
	cg "codegenerator"
	dbc "dbconn"
	"dbq"

	"server"

	envir "environment"
)

func main() {
	var env envir.Env
	var dbh dbc.TestPostgreDb // test/prod is here to decide
	env.EnvFactory(&dbh)
	env.SetEnvVariables()

	// --------------code generation---------------\\
	cg.EmptyCodeGen()
	cg.ImportConfigExcelDataTypes_httpToDb(&dbh)
	cg.ImportConfigExcelDistinctTypesToDb(&dbh)
	cg.ImportConfigExcelDataTypes_DTOToDb(&dbh)
	// cg.GenerateEntitiesData(&dbh) //ok
	// cg.GenerateEntitiesHistoryData(&dbh)
	// cg.GenerateRepos(&dbh)        //ok
	cg.GenerateDtoData(&dbh)
	// cg.GenerateDtoFiles_New(&dbh)
	cg.GenerateCustomDtoFiles(&dbh) // when custom implementation starts, this should'nt be called.
	// cg.GenerateServiceFiles(&dbh)
	// cg.GenerateHandlerFiles(&dbh)

	types := []string{"Question", "TaskStateChangeComment"}
	// types := []string{"Customer"}
	// cg.GenerateSpecificEntitiesData(&dbh, types)
	// cg.GenerateSpecificEntitiesHistoryData(&dbh, types)
	// cg.GenerateSpecificRepos(&dbh, types)
	// cg.GenerateSpecificDtoData(&dbh, types)
	// cg.GenerateSpecificCustomDtoFiles(&dbh, types)
	// cg.GenerateSpecificServiceFiles(&dbh, types)
	// cg.GenerateSpecificHandlerFiles(&dbh, types)

	// Generate all datatypes into one file
	cg.GenerateEntitiesDataInOneFile(&dbh)
	cg.GenerateEntitiesHistoryDataInOneFile(&dbh)
	cg.GenerateHandlerFilesInOneFile(&dbh)

	// --------------database actions---------------\\
	dbq.EmptyDbq()
	dbq.Migrate(&env)
	dbq.MigrateHistory(&env)
	dbq.MigrateSpecial(&env)

	// --------------state updates---------------\\
	// if errStateChanges := stateUpdates(&env); errStateChanges != nil {
	// 	panic(errStateChanges)
	// }

	// -------------msgraph config init--------------\\
	// msgraph.InitSpFolderConfig(&dbh)

	// --------------server init---------------\\
	server.EmptyHttpServer(&env)

	httpServer := server.InitiateServer(&env)
	fmt.Println("http server serving")
	log.Fatal(httpServer.ListenAndServe())

	// log.Fatal(httpServer.ListenAndServeTLS("/upm/tls/fullchain.pem", "/upm/tls/privkey.pem"))

}

func stateUpdates(env envir.Environment) error {
	// inserting actual statechanges from excel into db
	if errStateChangeInsert := dbq.UpdateActualStatesAndChanges(env); errStateChangeInsert != nil {
		return errStateChangeInsert
	}
	// execute task_state_sp.sql script
	if execTaskStateScript := dbq.ExecStateChangeUpdateScript(env); execTaskStateScript != nil {
		return execTaskStateScript
	}

	return nil
}

/*

	Fontos:
		SOLID principles:
			S: Only one purpose of a function or object
			O: Open/Closed principle -> Open for extension but closed for modification
			L: Liskov Substitution principle -> nem annyira érdekes
			I: Interface Segregation -> keep interfaces as simple as possible
			D: Dependency Inversion principle -> on higher levels things should be very loosely coupled
				and through interfaces and composed together.

		Other important Design Patterns:
			- Builder -> be able to create objects step-by-step
			- Factory -> creating big objecs
			- Proxy -> validations

*/
