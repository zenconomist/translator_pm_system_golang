package services

import (
	dt "dto"
	"entities"
	"environment"
	"reflect"
)

type IAction interface {
	Execute() error
}

type Actions struct {
	Actions []IAction
}

func (actions *Actions) AddNewMember(action IAction) {
	actions.Actions = append(actions.Actions, action)
}

type EmptyAction struct{}

func (a *EmptyAction) Execute() error {
	return nil
}

func GiveActions[E environment.Environment, T entities.Entity, D dt.DTO[T]](action string, entity T, dto D, env E, newID uint) []IAction {
	var actions Actions
	taskRelated := false // easier to write then to call from all separate entities.
	switch action {
	case "services.(*Service[...]).CreateNewItem":
		switch reflect.TypeOf(entity).String() {
		case "entities.TaskOffered":
			// TODO: finalize e-mail sending (not hardcoded e-mail)
			actions.AddNewMember(InitNewTaskOfferingEmailSending(entity, dto, env, newID))
			actions.AddNewMember(InitNewTaskOfferingSettingTaskState(entity, dto, env, newID))
			taskRelated = true
		case "entities.Project":
			actions.AddNewMember(InitNewDefaultProjectTasks(entity, dto, env, newID))
			// TODO: CheckProjectStatesAction implementation
			actions.AddNewMember(InitCheckProjectStatesAction(entity, dto, env.GiveDbHandler().PassConnection()))
			actions.AddNewMember(InitAddMsGraphFolders(env))
			taskRelated = true
		case "entities.Task":
			actions.AddNewMember(InitCheckProjectStatesAction(entity, dto, env.GiveDbHandler().PassConnection()))
			actions.AddNewMember(InitAddMsGraphFolders(env))
			taskRelated = true
		default:
			// none
		}
	case "services.(*Service[...]).UpdateItem":
		switch reflect.TypeOf(entity).String() {
		default:
			// none
		}
	case "services.(*Service[...]).DeleteItem":
		switch reflect.TypeOf(entity).String() {
		case "entities.Project":
			actions.AddNewMember(InitDeleteAssociatedTasks(entity.GiveID(), env))
			taskRelated = true
		case "entities.TaskOffered":
			// TODO: Email sending implementation
			// actions.AddNewMember(InitTaskOfferingUserInteractionEmailSending(entity, env))
			// actions.AddNewMember(InitNewAcceptOrDeclineTaskOffering(entity.GiveID(), dto, env))
			taskRelated = true

		}
	default:
		actions.AddNewMember(&EmptyAction{})
	}
	if taskRelated {
		actions.AddNewMember(InitNewTaskHistoricization(env))
	}
	return actions.Actions
}
