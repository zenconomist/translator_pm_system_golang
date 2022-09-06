package services

import (
	"entities"
	"environment"
	"fmt"
	msgr "msgraph"
)

type IMsGraphService interface {
}

type MsGraphService struct {
	Env         environment.Environment
	GraphClient *msgr.GraphClient
}

func NewMsGraphService(env environment.Environment, gcl *msgr.GraphClient) *MsGraphService {
	return &MsGraphService{
		Env:         env,
		GraphClient: gcl,
	}
}

func (mgs *MsGraphService) CreateProjectFolders() error {
	// select all folders that have to be created
	cr := entities.NewCustomRepo(mgs.Env.GiveDbHandler().PassConnection(), entities.SharePointFolder{})
	spfs, errFindSpFolders := cr.FindAllSpFoldersToCreate()
	if errFindSpFolders != nil {
		return errFindSpFolders
	}
	for _, f := range spfs {
		var parentItemID_ string
		if f.FolderType == "pf" {
			parentItemID_ = mgs.Env.GiveEnvVariableValue("MsGraphProjectFolderID")
		} else {
			parentItem, errGetParentItem := cr.GetParentItem(f.ProjectID, f.ParentFolderConfigID)
			if errGetParentItem != nil {
				return errGetParentItem
			}
			parentItemID_ = parentItem.SPID
		}
		spf, errCreate := mgs.CreateFldr(f.FolderName, parentItemID_, f.FolderType, f.Hierarchy, f.ProjectID, f.TaskID, f.ConfigID, f.ParentFolderConfigID)
		if errCreate != nil {
			continue
			// return errCreate
		}
		mgs.Env.GiveLogger().SetInfo(fmt.Sprintf("folder created: %v", spf))
		mgs.Env.GiveLogger().Log()
	}

	return nil
}

func (mgs *MsGraphService) CreateFldr(fname string, pItemID string, ftype string, hierarchy int, projectID uint, taskID uint, configID int, parentFolderConfigID int) (entities.SharePointFolder, error) {
	fakeSpFolder := entities.SharePointFolder{}
	// create task folder within project
	driveID := mgs.Env.GiveEnvVariableValue("MsGraphDriveID")
	var f msgr.Folder
	f.Name = fname
	item, err0 := mgs.GraphClient.CreateProjectFolder(driveID, pItemID, f)
	if err0 != nil {
		return fakeSpFolder, err0
	}

	// add it to the db
	spFolder := entities.SharePointFolder{
		FolderType:           ftype,
		FolderName:           item.Name,
		ProjectID:            projectID,
		TaskID:               taskID,
		SPCreatedTime:        item.CreatedTime,
		SPID:                 item.ID,
		LastModified:         item.LastModified,
		WebURL:               item.WebURL,
		Size:                 int(item.Size),
		Context:              item.Context,
		Etag:                 item.ETag,
		Ctag:                 item.CTag,
		CreatedBy:            item.CreatedBy.Application.DisplayName,
		LastModifiedBy:       item.LastModifiedBy.Application.DisplayName,
		ParentDriveID:        item.ParentReference.DriveID,
		ParentDriveType:      item.ParentReference.DriveType,
		ParentPath:           item.ParentReference.Path,
		ChildCount:           int(item.Folder.ChildCount),
		Hierarchy:            hierarchy,
		ParentSPID:           pItemID,
		ConfigID:             configID,
		ParentFolderConfigID: parentFolderConfigID,
	}

	db := mgs.Env.GiveDbHandler().PassConnection()

	result := db.Omit("id").Create(&spFolder)
	if result.Error != nil {
		return spFolder, result.Error
	}

	lb1, err2 := mgs.GraphClient.CreateLinkToItem(driveID, item.ID, "edit")
	if err2 != nil {
		return spFolder, err2
	}
	resultUpdate1 := db.Model(&spFolder).Select("read_write_link").Updates(&entities.SharePointFolder{WebURL: lb1.Link.WebURL})
	if resultUpdate1.Error != nil {
		return spFolder, resultUpdate1.Error
	}

	lb2, err4 := mgs.GraphClient.CreateLinkToItem(driveID, item.ID, "view")
	if err2 != nil {
		return spFolder, err4
	}

	resultUpdate2 := db.Model(&spFolder).Select("read_link").Updates(&entities.SharePointFolder{WebURL: lb2.Link.WebURL})
	if resultUpdate2.Error != nil {
		return spFolder, resultUpdate2.Error
	}

	return spFolder, nil
}
