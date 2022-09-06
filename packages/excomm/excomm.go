package excomm

import (
	"entities"
	"environment"
	"net/smtp"
)

// excomm stands for: external communications

// -----------IComms -> the communicator handler------------ \\

type IComms interface {
	SendOnAllChannels() error
}

type Comms struct {
	Communicators []Communicator
	Env           environment.Environment
}

func NewComms(communicators []Communicator) *Comms {
	return &Comms{
		Communicators: communicators,
	}
}

func (c *Comms) SendOnAllChannels() error {
	for _, communicator := range c.Communicators {
		if err := communicator.SendMessage(); err != nil {
			return err
		}
	}
	return nil
}

// -----------Mail Communicator------------ \\

type Communicator interface {
	SendMessage() error
}

type Mail struct {
	Env        environment.Environment
	SentByFunc string
	To         string   `json:"to,omitempty"`
	ToRecips   []string `json:"torecips"`
	From       string   `json:"from"`
	Message    Msg      `json:"message"` // 0: Recipients aka To, 1: Subject, 2: Body
}

type Msg struct {
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func NewEMail(env environment.Environment, sentByFunc string, to string, message Msg) Mail {
	return Mail{
		Env:        env,
		SentByFunc: sentByFunc,
		To:         to,
		Message:    message,
	}
}

func NewMessage(subject, body string) Msg {
	return Msg{
		Subject: subject,
		Body:    body,
	}
}

func (m Mail) SendMessage() error {

	m.From = ""
	a := smtp.PlainAuth("", "", m.Env.GiveEnvVariableValue("MMTSHuSysMailPw"), "")
	m.ToRecips = append(m.ToRecips, m.To)
	m.ToRecips = append(m.ToRecips, "")
	msg := []byte("To: " + m.To + "\r\n" +
		"Subject: " + m.Message.Subject + "\r\n" +
		"MIME-version: 1.0; \nContent-Type: text/html; charset=\"UTF-8\";\n\n" +
		"\r\n" +
		m.Message.Body + "\r\n")
	err := smtp.SendMail("", a, m.From, m.ToRecips, msg)
	if err != nil {

		return err
	}
	mailLog := entities.EmailSendingLog{
		SentByFunc: m.SentByFunc,
		SentTo:     m.To,
		Subject:    m.Message.Subject,
		Body:       m.Message.Body,
	}
	repo := entities.NewRepository(m.Env.GiveDbHandler().PassConnection(), mailLog)
	_, errCreate := repo.Create(mailLog)
	if errCreate != nil {
		return errCreate
	}

	return nil
}
