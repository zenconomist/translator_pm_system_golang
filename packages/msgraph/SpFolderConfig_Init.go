package msgraph

import (
	dbc "dbconn"
	"entities"
	"fmt"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

const file = "./codegenerator/UPM_v0.2_DataModel_v1.xlsx"

func InitSpFolderConfig(dbh dbc.DbHandler) {
	// read in Excel file
	f, err := excelize.OpenFile(file)
	if err != nil {
		fmt.Println(err)
	}
	db := dbh.PassConnection()
	db.Exec("TRUNCATE TABLE share_point_folder_configs;")

	var spc entities.SharePointFolderConfig
	rowsSPC := f.GetRows("SpFolderStructure")
	for j, row := range rowsSPC {
		if j == 0 {
			// excel header to be skipped
			continue
		}
		configID, errConv1 := strconv.Atoi(row[0])
		if errConv1 != nil {
			fmt.Println(errConv1)
			return
		}
		spc.ConfigID = configID
		spc.FolderType = row[1]
		hierarchy, errConv2 := strconv.Atoi(row[2])
		if errConv2 != nil {
			fmt.Println(errConv2)
			return
		}
		spc.Hierarchy = hierarchy
		spc.FolderName = row[3]
		parentid, errConv3 := strconv.Atoi(row[4])
		if errConv3 != nil {
			fmt.Println(errConv3)
			return
		}
		spc.ParentFolderConfigID = parentid
		res := db.Omit("id").Create(&spc)
		if res.Error != nil {
			fmt.Println(res.Error)
			return
		}
	}
}
