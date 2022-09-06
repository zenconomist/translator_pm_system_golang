package msgraph



type Folder struct {
	Name string `json:"name"`
	Folder Fldr `json:"folder"`
}

type Fldr struct {
	// FolderHierarchy string `json:""`
}