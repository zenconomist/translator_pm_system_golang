package msgraph

// "encoding/json"
// "fmt"

type Mail struct {
	Message     Msg          `json:"message"`
	graphClient *GraphClient // the graphClient that called the group
}

type Msg struct {
	Subject      string        `json:"subject"`
	Body         MBody         `json:"body"`
	ToRecipients []MailAddress `json:"toRecipients"`
}

type MBody struct {
	ContentType string `json:"contentType"`
	Content     string `json:"content"`
}

type MailAddress struct {
	MailAddr Addr `json:"emailAddress"`
}

type Addr struct {
	Address string `json:"address"`
}

// func (m Mail) String() string {
// 	return fmt.Sprintf("Mail(Subject: ")
// }

func (m *Mail) SetGraphClient(mC *GraphClient) {
	m.graphClient = mC
}
