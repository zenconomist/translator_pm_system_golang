package services

import (
	"dto"
	"entities"
	"environment"
	"fmt"
	"main/packages/excomm"
	"reflect"
	"strconv"
)

// ----------------TaskOfferingEmailSending----------------\\

func InitNewTaskOfferingEmailSending[E entities.Entity, D dto.DTO[E]](entity E, dto D, env environment.Environment, newID uint) *TaskOfferingEmailSending {
	var to TaskOfferingEmailSending
	if reflect.TypeOf(entity).String() != "entities.TaskOffered" {
		defer env.GiveLogger().LogError(fmt.Errorf("by InitNewTaskOffering the given entity is not a TaskOffered entity"))
		return &to
	}
	to.Env = env
	to.TaskOfferID = newID
	return &to
}

type TaskOfferingEmailSending struct {
	Env         environment.Environment
	TaskOfferID uint
	TaskID      uint
	SupplierID  uint
}

func (to *TaskOfferingEmailSending) Execute() error {
	// send email to the supplier who has been offered
	repo := entities.NewRepository(to.Env.GiveDbHandler().PassConnection(), entities.TaskOffered{})
	toEntity, errFind := repo.FindByID(to.TaskOfferID)
	if errFind != nil {
		return errFind
	}
	to.TaskID = toEntity.TaskID
	to.SupplierID = toEntity.SupplierID
	mail := excomm.NewEMail(to.Env, to.Env.GiveLogger().GetCurrentFuncName(), "zsolt.kreisz@mmts.hu", excomm.NewMessage("Task offered to you with id: "+strconv.Itoa(int(to.TaskID)), "test"))
	var communicators []excomm.Communicator
	communicators = append(communicators, mail)
	comms := excomm.NewComms(communicators)
	if errSend := comms.SendOnAllChannels(); errSend != nil {
		return errSend
	}
	return nil
}

// ----------------TaskOfferingSettingTaskState----------------\\

func InitNewTaskOfferingSettingTaskState[E entities.Entity, D dto.DTO[E]](entity E, dto D, env environment.Environment, newID uint) *TaskOfferingSettingTaskState {
	var to TaskOfferingSettingTaskState
	if reflect.TypeOf(entity).String() != "entities.TaskOffered" {
		defer env.GiveLogger().LogError(fmt.Errorf("by InitNewTaskOffering the given entity is not a TaskOffered entity"))
		return &to
	}
	to.Env = env
	to.TaskOfferID = newID
	return &to
}

type TaskOfferingSettingTaskState struct {
	Env         environment.Environment
	TaskOfferID uint
	TaskID      uint
}

// purpose is to set the task state to Offered
func (to *TaskOfferingSettingTaskState) Execute() error {
	repo := entities.NewRepository(to.Env.GiveDbHandler().PassConnection(), entities.TaskOffered{})
	toEntity, errFind := repo.FindByID(to.TaskOfferID)
	if errFind != nil {
		return errFind
	}
	to.TaskID = toEntity.TaskID
	var initTask entities.Task
	taskRepo := entities.NewRepository(to.Env.GiveDbHandler().PassConnection(), initTask)
	task, errFindTask := taskRepo.FindByID(to.TaskID)
	if errFindTask != nil {
		return errFindTask
	}
	tss := NewTaskStateService(&task, to.Env)
	tss.SetState("Offered")

	return nil
}
