package services

import (
	"environment"
	msgr "msgraph"
)

// ----------------MsGraph folders----------------\\

type AddMsGraphFolders struct {
	Env environment.Environment
}

func InitAddMsGraphFolders(env environment.Environment) *AddMsGraphFolders {
	return &AddMsGraphFolders{
		Env: env,
	}
}

func (amf *AddMsGraphFolders) Execute() error {
	// add MsGraph folders
	gcl, errgcl := msgr.NewGraphClient(amf.Env.GiveEnvVariableValue("MsGraphTenantID"), amf.Env.GiveEnvVariableValue("MsGraphAppID"), amf.Env.GiveEnvVariableValue("MsGraphClientSecret"))
	if errgcl != nil {
		return errgcl
	}
	mgs := NewMsGraphService(amf.Env, gcl)
	if errCreateFolders := mgs.CreateProjectFolders(); errCreateFolders != nil {
		return errCreateFolders
	}
	return nil

}
