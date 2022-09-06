package environment

import (
	dbc "dbconn"
	"entities"
	"fmt"
	logging "logging"
)

type Environment interface {
	EmptyFunc()
	EnvFactory(dbc.DbHandler)
	GiveDbHandler() dbc.DbHandler
	GiveLogger() logging.Logger
	GiveEnvVariableValue(string) string
}

type Env struct {
	DbHandler              dbc.DbHandler
	Logger                 logging.Logger
	MsGraphTenantID        string
	MsGraphAppID           string
	MsGraphClientSecret    string
	MsGraphDriveID         string
	MsGraphProjectFolderID string
	MMTSHuSysMailPw        string
}

type proxyDbHandler struct {
	dbh dbc.DbHandler
}

func (env *Env) EnvFactory(dbh dbc.DbHandler) {
	env.DbHandler = dbh
	env.DbHandler.SetDbConn()

	// implementing exact logger type
	var gl entities.UPMLogger

	env.Logger = &gl
	env.Logger.SetLoggerDbHandler(env.DbHandler)
	fmt.Println(fmt.Sprintf("env logger: ", env.Logger.GetLoggerDbHandler()))
	env.Logger.InitLoggerPerFunc(env.Logger.GetCurrentFuncName(), "EnvFactory on")
	env.Logger.Log()

	// setting logger's database handler as the same
	fmt.Println("EnvFactory ok.")

}
func (env *Env) GiveDbHandler() dbc.DbHandler {
	return env.DbHandler
}

func (env *Env) GiveLogger() logging.Logger {
	return env.Logger
}

func (env *Env) GiveEnvVariableValue(variable string) string {
	switch variable {
	case "MsGraphTenantID":
		return env.MsGraphTenantID
	case "MsGraphAppID":
		return env.MsGraphAppID
	case "MsGraphClientSecret":
		return env.MsGraphClientSecret
	case "MsGraphDriveID":
		return env.MsGraphDriveID
	case "MsGraphProjectFolderID":
		return env.MsGraphProjectFolderID
	case "MMTSHuSysMailPw":
		return env.MMTSHuSysMailPw
	default:
		return ""
	}
}

func (env Env) EmptyFunc() {}
