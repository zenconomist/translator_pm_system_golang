package msgraph

type LinkBase struct {
	ID    string   `json:"id"`
	Roles []string `json:"roles"`
	Link  Link     `json:"link"`
}

type Link struct {
	Type   string      `json:"type"`
	Scope  string      `json:"scope"`
	WebURL string      `json:"webUrl"`
	Appl   Application `json"application"`
}

type Application struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
}
