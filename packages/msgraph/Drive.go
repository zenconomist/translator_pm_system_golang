package msgraph

type Drives struct {
	Context     string       `json:"@odata.context"`
	Drive       []Drive      `json:"value"`
	graphClient *GraphClient // the graphClient that called the group
}

type Drive struct {
	CreatedTime    string     `json:"createdDateTime"`
	Description    string     `json:"description"`
	LastModified   string     `json:"lastModifiedDateTime"`
	ID             string     `json:"id"`
	DriveType      string     `json:"driveType"` // like "documentLibrary"
	Name           string     `json:"name"`
	CreatedBy      CreatedBy  `json:"createdBy"`
	LastModifiedBy LModBy     `json:"lastModifiedBy"`
	Owner          DriveOwner `json:"owner"`
	WebURL         string     `json:"webUrl"`
}

type LModBy struct {
	LmodUser LModUSR `json:"user"`
}

type LModUSR struct {
	Email       string `json:"email"`
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
}

type CreatedBy struct {
	CbUser CbUser `json:"user"`
}

type CbUser struct {
	DisplayName string `json:"displayName"`
}

type DriveOwner struct {
	Group DriveGroup `json:"group"`
	Quota DriveQuota `json:"quota"`
	// User User `json:"user"`
}

type DriveQuota struct {
	Deleted   int64  `json:"deleted"`
	Remaining int64  `json:"remaining"`
	State     string `json:"state"`
	Total     int64  `json:"total"`
	Used      int64  `json:"used"`
}

type DriveGroup struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
}

type DriveUser struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
}

func (d *Drives) SetGraphClient(g *GraphClient) {
	d.graphClient = g
}
