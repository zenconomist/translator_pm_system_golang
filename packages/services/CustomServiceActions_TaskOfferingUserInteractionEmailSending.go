package services

import (
	"entities"
	"environment"
	"fmt"
)

type TaskOfferingUserInteractionEmailSending struct {
	To  *entities.TaskOffered
	Env environment.Environment
}

func InitTaskOfferingUserInteractionEmailSending(to *entities.TaskOffered, env environment.Environment) *TaskOfferingUserInteractionEmailSending {
	return &TaskOfferingUserInteractionEmailSending{
		To:  to,
		Env: env,
	}
}

func (toes *TaskOfferingUserInteractionEmailSending) Execute() error {
	// sends email to pm that task offer was accepted

	// sends email to the Supplier that he/she accepted the task
	// send email to all other suppliers to whom this task was offered
	fmt.Println("send all the mails")
	return nil
}
