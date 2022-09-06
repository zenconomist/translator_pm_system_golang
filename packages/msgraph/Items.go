package msgraph

type Items struct {
	Context     string       `json:"@odata.context"`
	Items       []Item       `json:"value"`
	graphClient *GraphClient // the graphClient that called the group
}

type Item struct {
	Context         string           `json:"@odata.context,omitempty"`
	CreatedTime     string           `json:"createdDateTime"`
	ETag            string           `json:"eTag"`
	ID              string           `json:"id"`
	LastModified    string           `json:"lastModifiedDateTime"`
	Name            string           `json:"name"`
	WebURL          string           `json:"webUrl"`
	CTag            string           `json:"cTag"`
	Size            int64            `json:"size"`
	CreatedBy       ItemsCreatedBy   `json:"createdBy"`
	LastModifiedBy  ItemsLModBy      `json:"lastModifiedBy"`
	ParentReference ItemsParentRef   `json:"parentReference"`
	FileSystemInfo  ItemsFileSysInfo `json:"fileSystemInfo"`
	Folder          ItemsFolder      `json:"folder,omitempty"`
}

type ItemsCreatedBy struct {
	Application ItemsApp `json:"application"`
}

type ItemsApp struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
}

type ItemsLModBy struct {
	Application ItemsApp `json:"application"`
}

type ItemsParentRef struct {
	DriveID   string `json:"driveId"`
	DriveType string `json:"driveType"`
	ID        string `json:"id"`
	Path      string `json:"path"`
}

type ItemsFileSysInfo struct {
	CreatedTime  string      `json:"createdDateTime"`
	LastModified string      `json:"lastModifiedDateTime"`
	Folder       ItemsFolder `json:"folder,omitempty"`
}

type ItemsFolder struct {
	ChildCount int64 `json:"childCount"`
}
