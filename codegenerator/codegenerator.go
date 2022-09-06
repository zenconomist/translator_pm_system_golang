package codegenerator

import (
	"entities"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	dbc "dbconn"
	// "os"

	"github.com/360EntSecGroup-Skylar/excelize"
	"gorm.io/gorm"
)

type IConfig interface {
	MapToConfig() Config
}

type Config struct {
	Type       string
	Field      string
	GolangType string
	GormTag    string
	JsonTag    string
	// PackageName string
	Enum    string
	Comment string
}

type UpmConfig struct {
	entities.Model
	Type       string
	Field      string
	GolangType string
	GormTag    string
	JsonTag    string
	// PackageName string
	Enum           string
	HistoryGormTag string
	Comment        string
}

func (uc *UpmConfig) MapToConfig() Config {
	var c Config
	c.Type = uc.Type
	c.Field = uc.Field
	c.GolangType = uc.GolangType
	c.GormTag = uc.GormTag
	c.JsonTag = uc.JsonTag
	c.Enum = uc.Enum
	c.Comment = uc.Comment
	return c
}

type DtoConfig struct {
	entities.Model
	Entity     string
	Type       string
	Field      string
	GolangType string
	// GormTag    string
	JsonTag string
	// PackageName string
	Enum              string
	RequestOrResponse string
	Comment           string
}

func (dc *DtoConfig) MapToConfig() Config {
	var c Config
	c.Type = dc.Type
	c.Field = dc.Field
	c.GolangType = dc.GolangType
	// c.GormTag = dc.GormTag
	c.JsonTag = dc.JsonTag
	c.Enum = dc.Enum
	c.Comment = dc.Comment
	return c
}

type DistinctType struct {
	Types           string
	Automigrate     bool
	ServiceGenerate bool
	DtoGenerate     bool
	HandlerGenerate bool
	HistoryGenerate bool
	HandlerGet      bool
	HandlerCrud     bool
	DbNames         string
}

type ToGenerate int

const (
	Entities ToGenerate = iota
	Repos
	Service
	Dto
	Handler
	AutoMigrate
	HandlerGet
	HandlerCrud
)

const file = "./codegenerator/UPM_v0.2_DataModel_v1.xlsx"

func EmptyCodeGen() {}

// ====================GENERATE INTO ONE FILE==================\\

func GenerateEntitiesDataInOneFile(dbh dbc.DbHandler) {
	configMap, err := GetConfigMaps(dbh.PassConnection())
	if err != nil {
		fmt.Println("couldn't get config maps")
		return
	}
	sp := " "      // one space
	eol := " \r\n" //end of line
	tag := "`"
	apostr := "\""
	distinctTypesMap, errDTypes := GetDistinctTypeMap(dbh.PassConnection())
	if errDTypes != nil {
		fmt.Println("couldn't get distinct type map")
		return
	}

	fullFileStr := strings.Builder{}
	fullFileStr.WriteString(getTemplateStrings("entity_template_all"))

	// ----------------ENTITIES---------------\\
	for k, v := range configMap {
		fullTextString := getTemplateStrings("entity_template_datatype")

		fields := strings.Builder{}
		for _, field := range v {
			fields.WriteString(field.Field)
			fields.WriteString(sp)
			if field.Enum == "y" {
				fields.WriteString("globalconstants.")
			}
			fields.WriteString(field.GolangType)
			fields.WriteString(sp)
			if field.JsonTag != "<<empty>>" && field.GormTag != "<<empty>>" {
				fields.WriteString(tag)
				fields.WriteString("json:")
				fields.WriteString(apostr)
				fields.WriteString(field.JsonTag)
				fields.WriteString(apostr)
				fields.WriteString(sp)
				fields.WriteString("gorm:")
				fields.WriteString(apostr)
				fields.WriteString(field.GormTag)
				fields.WriteString(apostr)
				fields.WriteString(tag)
			} else {
				if field.JsonTag != "<<empty>>" {
					fields.WriteString(tag)
					fields.WriteString("json:")
					fields.WriteString(apostr)
					fields.WriteString(field.JsonTag)
					fields.WriteString(apostr)
					fields.WriteString(tag)
				}
				if field.GormTag != "<<empty>>" {
					fields.WriteString(tag)
					fields.WriteString("gorm:")
					fields.WriteString(apostr)
					fields.WriteString(field.GormTag)
					fields.WriteString(apostr)
					fields.WriteString(tag)
				}

			}
			fields.WriteString(eol)
		}
		fullTextString = strings.Replace(fullTextString, "<<entity>>", k, -1)
		fullTextString = strings.Replace(fullTextString, "<<entity.struct.fields>>", fields.String(), -1)
		if !distinctTypesMap[k].Automigrate {
			fullTextString = strings.ReplaceAll(fullTextString, "Model", "")
		} else {
			giveID := getTemplateStrings("entity_giveid_template")
			giveID = strings.ReplaceAll(giveID, "<<entity>>", k)
			fullTextString = fullTextString + giveID
		}

		// fmt.Println(fields.String())
		// fmt.Println(fullTextString)
		fullFileStr.WriteString(fullTextString)

	}

	errCreateFile := ioutil.WriteFile("./packages/entities/GEN_entities_all.go", []byte(fullFileStr.String()), 0777)
	if errCreateFile != nil {
		fmt.Println("couldn't create file", errCreateFile)
	}

	fmt.Println("entity files generated")
}

func GenerateEntitiesHistoryDataInOneFile(dbh dbc.DbHandler) {
	configMap, err := GetConfigMaps(dbh.PassConnection())
	if err != nil {
		fmt.Println("couldn't get config maps")
		return
	}
	sp := " "      // one space
	eol := " \r\n" //end of line
	tag := "`"
	apostr := "\""
	distinctTypesMap, errDTypes := GetDistinctTypeMap(dbh.PassConnection())
	if errDTypes != nil {
		fmt.Println("couldn't get distinct type map")
		return
	}

	fullFileStr := strings.Builder{}
	fullFileStr.WriteString(getTemplateStrings("entity_history_template_all"))

	// ----------------ENTITIES HISTORY---------------\\
	for k, v := range configMap {
		if !distinctTypesMap[k].HistoryGenerate {
			continue
		}
		fullTextString := getTemplateStrings("entity_history_datatype_template")

		fields := strings.Builder{}
		for _, field := range v {
			if field.Comment == "interface" {
				// interfaces cannot be mapped when creating the history, and are irrelevant for persistance
				// since their code is going to be persisted (uint value from the constant enums)
				continue
			}
			fields.WriteString(field.Field)
			fields.WriteString(sp)
			if field.Enum == "y" {
				fields.WriteString("globalconstants.")
			}
			fields.WriteString(field.GolangType)
			fields.WriteString(sp)
			if field.JsonTag != "<<empty>>" && field.HistoryGormTag != "<<empty>>" {
				fields.WriteString(tag)
				fields.WriteString("json:")
				fields.WriteString(apostr)
				fields.WriteString(field.JsonTag)
				fields.WriteString(apostr)
				fields.WriteString(sp)
				fields.WriteString("gorm:")
				fields.WriteString(apostr)
				fields.WriteString(field.HistoryGormTag)
				fields.WriteString(apostr)
				fields.WriteString(tag)
			} else {
				if field.JsonTag != "<<empty>>" {
					fields.WriteString(tag)
					fields.WriteString("json:")
					fields.WriteString(apostr)
					fields.WriteString(field.JsonTag)
					fields.WriteString(apostr)
					fields.WriteString(tag)
				}
				if field.HistoryGormTag != "<<empty>>" {
					fields.WriteString(tag)
					fields.WriteString("gorm:")
					fields.WriteString(apostr)
					fields.WriteString(field.HistoryGormTag)
					fields.WriteString(apostr)
					fields.WriteString(tag)
				}

			}
			fields.WriteString(eol)
		}
		fullTextString = strings.Replace(fullTextString, "<<entity>>", k, -1)
		fullTextString = strings.Replace(fullTextString, "<<entity.struct.fields>>", fields.String(), -1)

		fullFileStr.WriteString(fullTextString)

	}
	errCreateFile := ioutil.WriteFile("./packages/entities/GEN_entities_history_all.go", []byte(fullFileStr.String()), 0777)
	if errCreateFile != nil {
		fmt.Println("couldn't create file", errCreateFile)
	}
	fmt.Println("entity files generated")
}

func GenerateHandlerFilesInOneFile(dbh dbc.DbHandler) {
	generateFile(dbh, "handler_get", "server", "get_handlers", HandlerGet)
	generateFile(dbh, "handler_crud", "server", "crud_handlers", HandlerCrud)
	fmt.Println("handlers file generated")
}

func generateFile(dbh dbc.DbHandler, inputfn, outputfldr, outputsuffix string, tg ToGenerate) {
	fullFileStr := strings.Builder{}
	realConfigMap := make(map[string][]Config)
	configMap, err := GetConfigMaps(dbh.PassConnection())
	if err != nil {
		fmt.Println("couldn't get config maps")
		return
	}
	// fullConfigDtoMap, errDto := GetDtoConfigMaps(dbh.PassConnection())
	// if errDto != nil {
	// 	fmt.Println("couldn't get config maps")
	// 	return
	// }
	switch tg {
	case Entities, Repos, Service, Handler, HandlerCrud, HandlerGet:
		for k, v := range configMap {
			for _, j := range v {
				realConfigMap[k] = append(realConfigMap[k], j.MapToConfig())
			}
		}
		// case Dto:
		// 	for _, types := range fullConfigDtoMap {
		// 		for _, configDtoMap := range types {
		// 			for k, v := range configDtoMap {
		// 				for _, j := range v {
		// 					realConfigMap[k] = append(realConfigMap[k], j.MapToConfig())
		// 				}
		// 			}

		// 		}
		// 	}
	}
	distinctTypesMap, errDTypes := GetDistinctTypeMap(dbh.PassConnection())
	if errDTypes != nil {
		fmt.Println("couldn't get distinct type map")
		return
	}

	switch tg {
	case HandlerGet, HandlerCrud:
		fullFileStr.WriteString(getTemplateStrings("handler_header_template"))
	}
	// fullFileStr.WriteString(getTemplateStrings(inputfn + "_template"))
	fmt.Println(realConfigMap)
	for k := range realConfigMap {
		fmt.Println(k)
		// to skip certain generations which are not needed
		switch tg {
		case Entities:
		case Repos:
		case Service:
			if !distinctTypesMap[k].ServiceGenerate {
				continue
			}
		case Dto:
			if !distinctTypesMap[k].DtoGenerate {
				continue
			}
		case Handler:
			if !distinctTypesMap[k].HandlerGenerate {
				continue
			}
		case HandlerGet:
			if !distinctTypesMap[k].HandlerGet {
				fmt.Println("handler get skipped for ", k)
				continue
			}
		case HandlerCrud:
			if !distinctTypesMap[k].HandlerCrud {
				fmt.Println("handler crud skipped for ", k)
				continue
			}
		}
		entity := k
		entityLowerCase := strings.ToLower(k)
		fullTextString := getTemplateStrings(inputfn + "_template")

		fullTextString = strings.ReplaceAll(fullTextString, "<<entity>>", entity)
		fullTextString = strings.ReplaceAll(fullTextString, "<<entity.lowercase>>", entityLowerCase)
		if !distinctTypesMap[k].Automigrate {
			fullTextString = strings.ReplaceAll(fullTextString, "Model", "")
		} else {
			if inputfn != "dto_custom" && inputfn != "handler" && inputfn != "handler_crud" && inputfn != "handler_get" {
				giveID := getTemplateStrings("entity_giveid_template")
				fullTextString = fullTextString + giveID
			}
		}
		fmt.Println(fullTextString)
		fullFileStr.WriteString(fullTextString)

	}

	fileData := []byte(fullFileStr.String())
	errCreateFile := ioutil.WriteFile("./packages/"+outputfldr+"/GEN_"+outputsuffix+".go", fileData, 0777)
	if errCreateFile != nil {
		fmt.Println("couldn't create file", errCreateFile)
	}
	fmt.Println("file generated")
}

// ====================SPECIFIC FILES TO GENERATE==================\\

func GenerateSpecificHandlerFiles(dbh dbc.DbHandler, types []string) {
	tg := Handler
	generateSpecificFiles(dbh, "handler", "server", "handlers", tg, types)
	fmt.Println("specific handler files generated")
}

func GenerateSpecificServiceFiles(dbh dbc.DbHandler, types []string) {
	tg := Service
	generateSpecificFiles(dbh, "service", "services", "service", tg, types)
	fmt.Println("specific service files generated")
}

func GenerateSpecificCustomServiceFiles(dbh dbc.DbHandler, types []string) {
	tg := Service
	generateSpecificFiles(dbh, "service_custom", "services", "customservice", tg, types)
	fmt.Println("specific custom service files generated")
}

func GenerateSpecificRepos(dbh dbc.DbHandler, types []string) {
	tg := Repos
	generateSpecificFiles(dbh, "repo", "repositories", "repo", tg, types)
	fmt.Println("specific repo files generated")
}

func GenerateSpecificCustomDtoFiles(dbh dbc.DbHandler, types []string) {
	tg := Dto
	generateSpecificFiles(dbh, "dto_custom", "dto", "dto_custom", tg, types)
	fmt.Println("specific custom dto files generated")
}

func GenerateSpecificDtoData(dbh dbc.DbHandler, types []string) {
	configMap, err := GetDtoConfigMaps(dbh.PassConnection())
	if err != nil {
		fmt.Println("couldn't get config maps")
		return
	}
	distinctTypesMap, errDTypes := GetDistinctTypeMap(dbh.PassConnection())
	if errDTypes != nil {
		fmt.Println("couldn't get distinct type map")
		return
	}

	sp := " "      // one space
	eol := " \r\n" //end of line
	tag := "`"
	apostr := "\""

	// ----------------ENTITIES---------------\\
	for k, v := range configMap {
		//skip if not in parameter slice
		if !contains(types, k) {
			continue
		}
		if !distinctTypesMap[k].DtoGenerate {
			continue
		}
		fullTextString := getTemplateStrings("dto_template")
		fields := strings.Builder{}
		for _, field := range v {
			// for _, lines := range dtoTypes {

			// for _, field := range dtoTypes {
			fields.WriteString(field.Field)
			fields.WriteString(sp)
			// enums and state objects have to be converted into strings

			if field.Enum == "y" {
				fields.WriteString("globalconstants.")
			}
			fields.WriteString(field.GolangType)
			fields.WriteString(sp)
			// only json tags have to be added to dto-s
			if field.JsonTag != "<<empty>>" {
				fields.WriteString(tag)
				fields.WriteString("json:")
				fields.WriteString(apostr)
				fields.WriteString(field.JsonTag)
				fields.WriteString(apostr)
				fields.WriteString(tag)
			}

			fields.WriteString(eol)
			// }
			// }

		}

		fullTextString = strings.Replace(fullTextString, "<<entity>>", k, -1)
		fullTextString = strings.Replace(fullTextString, "<<dto.struct.fields>>", fields.String(), -1)

		// fmt.Println(fields.String())
		// fmt.Println(fullTextString)
		fileData := []byte(fullTextString)
		errCreateFile := ioutil.WriteFile("./packages/dto/"+strings.ToLower(k)+"_dto.go", fileData, 0777)
		if errCreateFile != nil {
			fmt.Println("couldn't create file", errCreateFile)
		}

	}
	fmt.Println("specific dto files generated")
}

func GenerateSpecificEntitiesData(dbh dbc.DbHandler, types []string) {
	configMap, err := GetConfigMaps(dbh.PassConnection())
	if err != nil {
		fmt.Println("couldn't get config maps")
		return
	}
	distinctTypesMap, errDTypes := GetDistinctTypeMap(dbh.PassConnection())
	if errDTypes != nil {
		fmt.Println("couldn't get distinct type map")
		return
	}
	sp := " "      // one space
	eol := " \r\n" //end of line
	tag := "`"
	apostr := "\""

	// ----------------ENTITIES---------------\\
	for k, v := range configMap {
		//skip if not in parameter slice
		if !contains(types, k) {
			continue
		}
		fullTextString := getTemplateStrings("entity_template")

		fields := strings.Builder{}
		for _, field := range v {
			fields.WriteString(field.Field)
			fields.WriteString(sp)
			if field.Enum == "y" {
				fields.WriteString("globalconstants.")
			}
			fields.WriteString(field.GolangType)
			fields.WriteString(sp)
			if field.JsonTag != "<<empty>>" && field.GormTag != "<<empty>>" {
				fields.WriteString(tag)
				fields.WriteString("json:")
				fields.WriteString(apostr)
				fields.WriteString(field.JsonTag)
				fields.WriteString(apostr)
				fields.WriteString(sp)
				fields.WriteString("gorm:")
				fields.WriteString(apostr)
				fields.WriteString(field.GormTag)
				fields.WriteString(apostr)
				fields.WriteString(tag)
			} else {
				if field.JsonTag != "<<empty>>" {
					fields.WriteString(tag)
					fields.WriteString("json:")
					fields.WriteString(apostr)
					fields.WriteString(field.JsonTag)
					fields.WriteString(apostr)
					fields.WriteString(tag)
				}
				if field.GormTag != "<<empty>>" {
					fields.WriteString(tag)
					fields.WriteString("gorm:")
					fields.WriteString(apostr)
					fields.WriteString(field.GormTag)
					fields.WriteString(apostr)
					fields.WriteString(tag)
				}

			}
			fields.WriteString(eol)
		}
		fullTextString = strings.Replace(fullTextString, "<<entity>>", k, -1)
		fullTextString = strings.Replace(fullTextString, "<<entity.struct.fields>>", fields.String(), -1)
		if !distinctTypesMap[k].Automigrate {
			fullTextString = strings.ReplaceAll(fullTextString, "Model", "")
		} else {
			giveID := getTemplateStrings("entity_giveid_template")
			giveID = strings.ReplaceAll(giveID, "<<entity>>", k)
			fullTextString = fullTextString + giveID
		}
		// fmt.Println(fields.String())
		// fmt.Println(fullTextString)
		fileData := []byte(fullTextString)
		errCreateFile := ioutil.WriteFile("./packages/entities/E_"+strings.ToLower(k)+"_entity.go", fileData, 0777)
		if errCreateFile != nil {
			fmt.Println("couldn't create file", errCreateFile)
		}

	}
	fmt.Println("specific entity files generated")
}

func GenerateSpecificEntitiesHistoryData(dbh dbc.DbHandler, types []string) {
	configMap, err := GetConfigMaps(dbh.PassConnection())
	if err != nil {
		fmt.Println("couldn't get config maps")
		return
	}
	sp := " "      // one space
	eol := " \r\n" //end of line
	tag := "`"
	apostr := "\""
	distinctTypesMap, errDTypes := GetDistinctTypeMap(dbh.PassConnection())
	if errDTypes != nil {
		fmt.Println("couldn't get distinct type map")
		return
	}
	// ----------------ENTITIES HISTORY---------------\\
	for k, v := range configMap {
		//skip if not in parameter slice
		if !contains(types, k) {
			continue
		}
		if !distinctTypesMap[k].HistoryGenerate {
			continue
		}
		fullTextString := getTemplateStrings("entity_history_template")

		fields := strings.Builder{}
		for _, field := range v {
			fields.WriteString(field.Field)
			fields.WriteString(sp)
			if field.Enum == "y" {
				fields.WriteString("globalconstants.")
			}
			fields.WriteString(field.GolangType)
			fields.WriteString(sp)
			if field.JsonTag != "<<empty>>" && field.HistoryGormTag != "<<empty>>" {
				fields.WriteString(tag)
				fields.WriteString("json:")
				fields.WriteString(apostr)
				fields.WriteString(field.JsonTag)
				fields.WriteString(apostr)
				fields.WriteString(sp)
				fields.WriteString("gorm:")
				fields.WriteString(apostr)
				fields.WriteString(field.HistoryGormTag)
				fields.WriteString(apostr)
				fields.WriteString(tag)
			} else {
				if field.JsonTag != "<<empty>>" {
					fields.WriteString(tag)
					fields.WriteString("json:")
					fields.WriteString(apostr)
					fields.WriteString(field.JsonTag)
					fields.WriteString(apostr)
					fields.WriteString(tag)
				}
				if field.HistoryGormTag != "<<empty>>" {
					fields.WriteString(tag)
					fields.WriteString("gorm:")
					fields.WriteString(apostr)
					fields.WriteString(field.HistoryGormTag)
					fields.WriteString(apostr)
					fields.WriteString(tag)
				}

			}
			fields.WriteString(eol)
		}
		fullTextString = strings.Replace(fullTextString, "<<entity>>", k, -1)
		fullTextString = strings.Replace(fullTextString, "<<entity.struct.fields>>", fields.String(), -1)

		fileData := []byte(fullTextString)
		errCreateFile := ioutil.WriteFile("./packages/entities/H_"+strings.ToLower(k)+"_entity_history.go", fileData, 0777)
		if errCreateFile != nil {
			fmt.Println("couldn't create file", errCreateFile)
		}

	}
	fmt.Println("entity history files generated")
}

func generateSpecificFiles(dbh dbc.DbHandler, inputfn, outputfldr, outputsuffix string, tg ToGenerate, types []string) {
	realConfigMap := make(map[string][]Config)
	configMap, err := GetConfigMaps(dbh.PassConnection())
	if err != nil {
		fmt.Println("couldn't get config maps")
		return
	}
	// configDtoMap, errDto := GetDtoConfigMaps(dbh.PassConnection())
	// if errDto != nil {
	// 	fmt.Println("couldn't get config maps")
	// 	return
	// }
	switch tg {
	case Entities, Repos, Service, Handler:
		for k, v := range configMap {
			for _, j := range v {
				realConfigMap[k] = append(realConfigMap[k], j.MapToConfig())
			}
		}
		// case Dto:
		// 	for k, v := range configDtoMap {
		// 		for _, j := range v {
		// 			realConfigMap[k] = append(realConfigMap[k], j.MapToConfig())
		// 		}
		// 	}
	}
	distinctTypesMap, errDTypes := GetDistinctTypeMap(dbh.PassConnection())
	if errDTypes != nil {
		fmt.Println("couldn't get distinct type map")
		return
	}
	for k := range realConfigMap {
		//skip if not in parameter slice
		if !contains(types, k) {
			continue
		}
		// to skip certain generations which are not needed
		switch tg {
		case Entities:
		case Repos:
		case Service:
			if !distinctTypesMap[k].ServiceGenerate {
				continue
			}
		case Dto:
			if !distinctTypesMap[k].DtoGenerate {
				continue
			}
		case Handler:
			if !distinctTypesMap[k].HandlerGenerate {
				continue
			}
		}
		entity := k
		entityLowerCase := strings.ToLower(k)
		fullTextString := getTemplateStrings(inputfn + "_template")

		fullTextString = strings.ReplaceAll(fullTextString, "<<entity>>", entity)
		fullTextString = strings.ReplaceAll(fullTextString, "<<entity.lowercase>>", entityLowerCase)
		fileData := []byte(fullTextString)
		errCreateFile := ioutil.WriteFile("./packages/"+outputfldr+"/"+strings.ToLower(k)+"_"+outputsuffix+".go", fileData, 0777)
		if errCreateFile != nil {
			fmt.Println("couldn't create file", errCreateFile)
		}

	}
	fmt.Println("files generated")
}

// ====================ALL FILES TO GENERATE==================\\

func GenerateHandlerFiles(dbh dbc.DbHandler) {
	tg := Handler
	generateFiles(dbh, "handler", "server", "handlers", tg)
	fmt.Println("handler files generated")
}

func GenerateCustomDtoFiles(dbh dbc.DbHandler) {
	tg := Dto
	generateFiles(dbh, "dto_custom", "dto", "dto_custom", tg)
}

func GenerateServiceFiles(dbh dbc.DbHandler) {
	tg := Service
	generateFiles(dbh, "service", "services", "service", tg)
}

func GenerateCustomServiceFiles(dbh dbc.DbHandler) {
	tg := Service
	generateFiles(dbh, "service_custom", "services", "customservice", tg)
}

func GenerateRepos(dbh dbc.DbHandler) {
	tg := Repos
	generateFiles(dbh, "repo", "repositories", "repo", tg)
}

func generateFiles(dbh dbc.DbHandler, inputfn, outputfldr, outputsuffix string, tg ToGenerate) {
	realConfigMap := make(map[string][]Config)
	configMap, err := GetConfigMaps(dbh.PassConnection())
	if err != nil {
		fmt.Println("couldn't get config maps")
		return
	}
	configDtoMap, errDto := GetDtoConfigMaps(dbh.PassConnection())
	if errDto != nil {
		fmt.Println("couldn't get config maps")
		return
	}
	switch tg {
	case Entities, Repos, Service, Handler:
		for k, v := range configMap {
			for _, j := range v {
				realConfigMap[k] = append(realConfigMap[k], j.MapToConfig())
			}
		}
	case Dto:
		for k, v := range configDtoMap {
			for _, j := range v {
				realConfigMap[k] = append(realConfigMap[k], j.MapToConfig())
			}
		}
	}
	distinctTypesMap, errDTypes := GetDistinctTypeMap(dbh.PassConnection())
	if errDTypes != nil {
		fmt.Println("couldn't get distinct type map")
		return
	}
	for k := range realConfigMap {
		// to skip certain generations which are not needed
		switch tg {
		case Entities:
		case Repos:
		case Service:
			if !distinctTypesMap[k].ServiceGenerate {
				continue
			}
		case Dto:
			if !distinctTypesMap[k].DtoGenerate {
				continue
			}
		case Handler:
			if !distinctTypesMap[k].HandlerGenerate {
				continue
			}
		}
		entity := k
		entityLowerCase := strings.ToLower(k)
		fullTextString := getTemplateStrings(inputfn + "_template")

		fullTextString = strings.ReplaceAll(fullTextString, "<<entity>>", entity)
		fullTextString = strings.ReplaceAll(fullTextString, "<<entity.lowercase>>", entityLowerCase)
		if !distinctTypesMap[k].Automigrate {
			fullTextString = strings.ReplaceAll(fullTextString, "Model", "")
		} else {
			if inputfn != "dto_custom" {
				giveID := getTemplateStrings("entity_giveid_template")
				fullTextString = fullTextString + giveID
			}
		}
		fileData := []byte(fullTextString)
		errCreateFile := ioutil.WriteFile("./packages/"+outputfldr+"/"+strings.ToLower(k)+"_"+outputsuffix+".go", fileData, 0777)
		if errCreateFile != nil {
			fmt.Println("couldn't create file", errCreateFile)
		}

	}
	fmt.Println("files generated")
}

func GenerateDtoFiles_New(dbh dbc.DbHandler) {

	// ==============INTERFACES=============== \\
	dtoInterfacesHeader := getTemplateStrings("dto_interfaces_header_template")
	dtoInterfaceBase := getTemplateStrings("dto_interfaces_template")
	distinctTypesMap, errDTypes := GetDistinctTypeMap(dbh.PassConnection())
	if errDTypes != nil {
		fmt.Println("couldn't get distinct type map")
		return
	}
	fullTextString := strings.Builder{}
	fullTextString.WriteString(dtoInterfacesHeader)
	for k, _ := range distinctTypesMap {
		fullTextString.WriteString(strings.ReplaceAll(dtoInterfaceBase, "<<entity>>", k))
	}

	fileData := []byte(fullTextString.String())
	errCreateFile := ioutil.WriteFile("./packages/dto/GEN_dto_interfaces.go", fileData, 0777)
	if errCreateFile != nil {
		fmt.Println("couldn't create file", errCreateFile)
	}

	// once generated successfully, custom implementations will take place, so no regen, or only with new filenames!
	/*
		// ==============CUSTOMDTO FILES=============== \\
		dtoCustomBase := getTemplateStrings("dto_custom_template")

		dtoEntityMap := getDtoEntityMaps()

		for dto, entity := range dtoEntityMap {
			fullCustomTextString := strings.Builder{}
			fullCustomTextString.WriteString("package dto \n")
			text := strings.ReplaceAll(dtoCustomBase, "<<entity>>", entity)
			text = strings.ReplaceAll(text, "<<dto>>", dto)
			fullCustomTextString.WriteString(text)
			fileData2 := []byte(fullCustomTextString.String())
			errCreateFile2 := ioutil.WriteFile("./packages/dto/GEN_"+dto+"_custom.go", fileData2, 0777) // when success, add GEN_ as filename prefix!
			if errCreateFile2 != nil {
				fmt.Println("couldn't create file", errCreateFile2)
			}
		}
	*/
}

func GenerateDtoData(dbh dbc.DbHandler) {
	configMap, err := GetDtoConfigMaps(dbh.PassConnection())
	if err != nil {
		fmt.Println("couldn't get config maps")
		return
	}

	fullTextString := getTemplateStrings("dto_header_template")

	sp := " "      // one space
	eol := " \r\n" //end of line
	tag := "`"
	apostr := "\""

	for dto, cm := range configMap {
		// ----------------ENTITIES---------------\\
		textString := getTemplateStrings("dto_template")
		fields := strings.Builder{}
		for _, field := range cm {
			fields.WriteString(field.Field)
			fields.WriteString(sp)
			// enums and state objects have to be converted into strings

			if field.Enum == "y" {
				fields.WriteString("globalconstants.")
			}
			fields.WriteString(field.GolangType)
			fields.WriteString(sp)
			// only json tags have to be added to dto-s
			if field.JsonTag != "<<empty>>" {
				fields.WriteString(tag)
				fields.WriteString("json:")
				fields.WriteString(apostr)
				fields.WriteString(field.JsonTag)
				fields.WriteString(apostr)
				fields.WriteString(tag)
			}

			fields.WriteString(eol)

		}
		textString = strings.ReplaceAll(textString, "<<dto>>", dto)
		textString = strings.ReplaceAll(textString, "<<dto.struct.fields>>", fields.String())
		// fmt.Println(fields.String())
		// fmt.Println(fullTextString)
		// fmt.Println(" \n new dto \n ")
		fullTextString += textString
	}
	fileData := []byte(fullTextString)
	errCreateFile := ioutil.WriteFile("./packages/dto/GEN_dtos.go", fileData, 0777)
	if errCreateFile != nil {
		fmt.Println("couldn't create file", errCreateFile)
	}

	fmt.Println("dto files generated")

}

func GenerateEntitiesData(dbh dbc.DbHandler) {
	configMap, err := GetConfigMaps(dbh.PassConnection())
	if err != nil {
		fmt.Println("couldn't get config maps")
		return
	}
	sp := " "      // one space
	eol := " \r\n" //end of line
	tag := "`"
	apostr := "\""
	distinctTypesMap, errDTypes := GetDistinctTypeMap(dbh.PassConnection())
	if errDTypes != nil {
		fmt.Println("couldn't get distinct type map")
		return
	}

	// ----------------ENTITIES---------------\\
	for k, v := range configMap {
		fullTextString := getTemplateStrings("entity_template")

		fields := strings.Builder{}
		for _, field := range v {
			fields.WriteString(field.Field)
			fields.WriteString(sp)
			if field.Enum == "y" {
				fields.WriteString("globalconstants.")
			}
			fields.WriteString(field.GolangType)
			fields.WriteString(sp)
			if field.JsonTag != "<<empty>>" && field.GormTag != "<<empty>>" {
				fields.WriteString(tag)
				fields.WriteString("json:")
				fields.WriteString(apostr)
				fields.WriteString(field.JsonTag)
				fields.WriteString(apostr)
				fields.WriteString(sp)
				fields.WriteString("gorm:")
				fields.WriteString(apostr)
				fields.WriteString(field.GormTag)
				fields.WriteString(apostr)
				fields.WriteString(tag)
			} else {
				if field.JsonTag != "<<empty>>" {
					fields.WriteString(tag)
					fields.WriteString("json:")
					fields.WriteString(apostr)
					fields.WriteString(field.JsonTag)
					fields.WriteString(apostr)
					fields.WriteString(tag)
				}
				if field.GormTag != "<<empty>>" {
					fields.WriteString(tag)
					fields.WriteString("gorm:")
					fields.WriteString(apostr)
					fields.WriteString(field.GormTag)
					fields.WriteString(apostr)
					fields.WriteString(tag)
				}

			}
			fields.WriteString(eol)
		}
		fullTextString = strings.Replace(fullTextString, "<<entity>>", k, -1)
		fullTextString = strings.Replace(fullTextString, "<<entity.struct.fields>>", fields.String(), -1)
		if !distinctTypesMap[k].Automigrate {
			fullTextString = strings.ReplaceAll(fullTextString, "Model", "")
		} else {
			giveID := getTemplateStrings("entity_giveid_template")
			giveID = strings.ReplaceAll(giveID, "<<entity>>", k)
			fullTextString = fullTextString + giveID
		}

		// fmt.Println(fields.String())
		// fmt.Println(fullTextString)
		fileData := []byte(fullTextString)
		errCreateFile := ioutil.WriteFile("./packages/entities/E_"+strings.ToLower(k)+"_entity.go", fileData, 0777)
		if errCreateFile != nil {
			fmt.Println("couldn't create file", errCreateFile)
		}

	}
	fmt.Println("entity files generated")
}

func GenerateEntitiesHistoryData(dbh dbc.DbHandler) {
	configMap, err := GetConfigMaps(dbh.PassConnection())
	if err != nil {
		fmt.Println("couldn't get config maps")
		return
	}
	sp := " "      // one space
	eol := " \r\n" //end of line
	tag := "`"
	apostr := "\""
	distinctTypesMap, errDTypes := GetDistinctTypeMap(dbh.PassConnection())
	if errDTypes != nil {
		fmt.Println("couldn't get distinct type map")
		return
	}
	// ----------------ENTITIES HISTORY---------------\\
	for k, v := range configMap {
		if !distinctTypesMap[k].HistoryGenerate {
			continue
		}
		fullTextString := getTemplateStrings("entity_history_template")

		fields := strings.Builder{}
		for _, field := range v {
			fields.WriteString(field.Field)
			fields.WriteString(sp)
			if field.Enum == "y" {
				fields.WriteString("globalconstants.")
			}
			fields.WriteString(field.GolangType)
			fields.WriteString(sp)
			if field.JsonTag != "<<empty>>" && field.HistoryGormTag != "<<empty>>" {
				fields.WriteString(tag)
				fields.WriteString("json:")
				fields.WriteString(apostr)
				fields.WriteString(field.JsonTag)
				fields.WriteString(apostr)
				fields.WriteString(sp)
				fields.WriteString("gorm:")
				fields.WriteString(apostr)
				fields.WriteString(field.HistoryGormTag)
				fields.WriteString(apostr)
				fields.WriteString(tag)
			} else {
				if field.JsonTag != "<<empty>>" {
					fields.WriteString(tag)
					fields.WriteString("json:")
					fields.WriteString(apostr)
					fields.WriteString(field.JsonTag)
					fields.WriteString(apostr)
					fields.WriteString(tag)
				}
				if field.HistoryGormTag != "<<empty>>" {
					fields.WriteString(tag)
					fields.WriteString("gorm:")
					fields.WriteString(apostr)
					fields.WriteString(field.HistoryGormTag)
					fields.WriteString(apostr)
					fields.WriteString(tag)
				}

			}
			fields.WriteString(eol)
		}
		fullTextString = strings.Replace(fullTextString, "<<entity>>", k, -1)
		fullTextString = strings.Replace(fullTextString, "<<entity.struct.fields>>", fields.String(), -1)

		fileData := []byte(fullTextString)
		errCreateFile := ioutil.WriteFile("./packages/entities/H_"+strings.ToLower(k)+"_entity_history.go", fileData, 0777)
		if errCreateFile != nil {
			fmt.Println("couldn't create file", errCreateFile)
		}

	}
	fmt.Println("entity files generated")
}

func getTemplateStrings(filename string) string {
	outputBytes, err0 := ioutil.ReadFile("./codegenerator/CGTemplates/" + filename + ".txt")
	if err0 != nil {
		log.Print(err0)
	}

	output := string(outputBytes)

	return output
}

func GetConfigMaps(db *gorm.DB) (map[string][]UpmConfig, error) {
	configMap := make(map[string][]UpmConfig)
	var cL []UpmConfig

	result := db.Find(&cL)
	if result.Error != nil {
		return configMap, result.Error
	}
	for _, c := range cL {
		configMap[c.Type] = append(configMap[c.Type], c)
	}
	return configMap, nil
}

func GetDtoConfigMaps(db *gorm.DB) (map[string][]DtoConfig, error) {

	configMap := make(map[string][]DtoConfig)
	var cL []DtoConfig

	result := db.Find(&cL)
	if result.Error != nil {
		return configMap, result.Error
	}

	for _, c := range cL {
		configMap[c.Type] = append(configMap[c.Type], c)
	}

	return configMap, nil
}

func GetDistinctTypeMap(db *gorm.DB) (map[string]DistinctType, error) {
	dtMap := make(map[string]DistinctType)
	var dtL []DistinctType

	result := db.Find(&dtL)
	if result.Error != nil {
		return dtMap, result.Error
	}
	for _, d := range dtL {
		dtMap[d.Types] = d
	}
	return dtMap, nil

}

func ImportConfigExcelDataTypes_httpToDb(dbh dbc.DbHandler) {
	// read in Excel file
	f, err := excelize.OpenFile(file)
	if err != nil {
		fmt.Println(err)
	}

	db := dbh.PassConnection()

	// UPM DataTypes for http server
	var c UpmConfig //c: config
	db.Exec("DROP TABLE IF EXISTS upm_configs;")
	db.AutoMigrate(&c) // migrate config
	// db.Exec("TRUNCATE TABLE upm_configs;")

	rowsDT := f.GetRows("DataTypes_http") // get sheet named after a distinct datatype
	for j, row := range rowsDT {
		if j == 0 {
			// excel header to be skipped
			continue
		}
		c.Type = row[0]
		c.Field = row[1]
		c.GolangType = row[2]
		c.GormTag = row[3]
		c.JsonTag = row[4]
		// c.PackageName = row[5]
		c.Enum = row[5]
		c.HistoryGormTag = row[6]
		c.Comment = row[7]
		c.Model.ID = uint(j)

		_ = db.Omit("Id").Create(&c)
		// fmt.Println(c.Model.ID)

	}

}

func ImportConfigExcelDistinctTypesToDb(dbh dbc.DbHandler) {
	// read in Excel file
	f, err := excelize.OpenFile(file)
	if err != nil {
		fmt.Println(err)
	}

	db := dbh.PassConnection()

	// UPM DataTypes for http server
	var dt DistinctType //c: config
	db.Exec("DROP TABLE IF EXISTS distinct_types;")
	db.AutoMigrate(&dt) // migrate config
	// db.Exec("TRUNCATE TABLE upm_configs;")

	rowsDT := f.GetRows("DistinctTypes") // get sheet named after a distinct datatype
	for j, row := range rowsDT {
		if j == 0 {
			// excel header to be skipped
			continue
		}
		dt.Types = row[0]
		am := row[1]
		if am == "y" {
			dt.Automigrate = true
		} else {
			dt.Automigrate = false
		}

		sg := row[2]
		if sg == "y" {
			dt.ServiceGenerate = true
		} else {
			dt.ServiceGenerate = false
		}
		dtogen := row[3]
		if dtogen == "y" {
			dt.DtoGenerate = true
		} else {
			dt.DtoGenerate = false
		}
		hgen := row[4]
		if hgen == "y" {
			dt.HandlerGenerate = true
		} else {
			dt.HandlerGenerate = false
		}
		hist := row[5]
		if hist == "y" {
			dt.HistoryGenerate = true
		} else {
			dt.HistoryGenerate = false
		}
		hGet := row[6]
		if hGet == "y" {
			dt.HandlerGet = true
		} else {
			dt.HandlerGet = false
		}
		hCrud := row[7]
		if hCrud == "y" {
			dt.HandlerCrud = true
		} else {
			dt.HandlerCrud = false
		}
		dt.DbNames = row[8]

		_ = db.Omit("Id").Create(&dt)
		// fmt.Println(c.Model.ID)

	}

}

// DataTypes_DTO

func ImportConfigExcelDataTypes_DTOToDb(dbh dbc.DbHandler) {
	// read in Excel file
	f, err := excelize.OpenFile(file)
	if err != nil {
		fmt.Println(err)
	}

	db := dbh.PassConnection()

	// UPM DataTypes for http server
	var c DtoConfig //c: config
	db.Exec("DROP TABLE IF EXISTS dto_configs;")
	db.AutoMigrate(&c) // migrate config
	// db.Exec("TRUNCATE TABLE upm_configs;")

	rowsDT := f.GetRows("DataTypes_DTO") // get sheet named after a distinct datatype
	for j, row := range rowsDT {
		if j == 0 {
			// excel header to be skipped
			continue
		}
		c.Entity = row[0]
		c.Type = row[1]
		c.Field = row[2]
		c.GolangType = row[3]
		// c.GormTag = row[3]
		c.JsonTag = row[4]
		// c.PackageName = row[5]
		c.Enum = row[5]
		c.RequestOrResponse = row[6]
		c.Comment = row[7]
		c.Model.ID = uint(j)

		_ = db.Omit("Id").Create(&c)
		// fmt.Println(c.Model.ID)

	}

}

func contains(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func getDtoEntityMaps() map[string]string {
	var dtoEntityMap = make(map[string]string)
	dtoEntityMap["CustomerRequestDTO"] = "Customer"
	dtoEntityMap["FirmRequestDTO"] = "Firm"
	dtoEntityMap["UserRequestDTO"] = "User"
	dtoEntityMap["RoleRequestDTO"] = "Role"
	dtoEntityMap["PermissionRequestDTO"] = "Permission"
	dtoEntityMap["ProjectRequestDTO"] = "Project"
	dtoEntityMap["BatchRequestDTO"] = "Batch"
	dtoEntityMap["TaskRequestDTO"] = "Task"
	dtoEntityMap["TaskCustomerPropsRequestDTO"] = "TaskCustomerProps"
	dtoEntityMap["TaskSupplierPropsRequestDTO"] = "TaskSupplierProps"
	dtoEntityMap["TaskConfigRequestDTO"] = "TaskConfig"
	dtoEntityMap["TaskOfferedRequestDTO"] = "TaskOffered"
	dtoEntityMap["ClientOfferRequestDTO"] = "ClientOffer"
	dtoEntityMap["CustomerPriceRequestDTO"] = "CustomerPrice"
	dtoEntityMap["SupplierPriceRequestDTO"] = "SupplierPrice"
	dtoEntityMap["DefaultCustomerPriceRequestDTO"] = "DefaultCustomerPrice"
	dtoEntityMap["DefaultSupplierPriceRequestDTO"] = "DefaultSupplierPrice"
	dtoEntityMap["ClientOfferTaskRequestDTO"] = "ClientOfferTask"
	dtoEntityMap["ClientOfferTaskCustomerPropsRequestDTO"] = "ClientOfferTaskCustomerProps"
	dtoEntityMap["BillingLogResponseDTO"] = "BillingLog"
	dtoEntityMap["UPMLoggerResponseDTO"] = "UPMLogger"
	dtoEntityMap["CustomerResponseDTO"] = "Customer"
	dtoEntityMap["FirmResponseDTO"] = "Firm"
	dtoEntityMap["UserResponseDTO"] = "User"
	dtoEntityMap["RoleResponseDTO"] = "Role"
	dtoEntityMap["PermissionResponseDTO"] = "Permission"
	dtoEntityMap["ProjectResponseDTO"] = "Project"
	dtoEntityMap["BatchResponseDTO"] = "Batch"
	dtoEntityMap["TaskResponseDTO"] = "Task"
	dtoEntityMap["TaskCustomerPropsResponseDTO"] = "TaskCustomerProps"
	dtoEntityMap["TaskSupplierPropsResponseDTO"] = "TaskSupplierProps"
	dtoEntityMap["TaskConfigResponseDTO"] = "TaskConfig"
	dtoEntityMap["TaskOfferedResponseDTO"] = "TaskOffered"
	dtoEntityMap["ClientOfferResponseDTO"] = "ClientOffer"
	dtoEntityMap["CustomerPriceResponseDTO"] = "CustomerPrice"
	dtoEntityMap["SupplierPriceResponseDTO"] = "SupplierPrice"
	dtoEntityMap["DefaultCustomerPriceResponseDTO"] = "DefaultCustomerPrice"
	dtoEntityMap["DefaultSupplierPriceResponseDTO"] = "DefaultSupplierPrice"
	dtoEntityMap["ClientOfferTaskResponseDTO"] = "ClientOfferTask"
	dtoEntityMap["ClientOfferTaskCustomerPropsResponseDTO"] = "ClientOfferTaskCustomerProps"
	dtoEntityMap["AddressResponseDTO"] = "Address"
	return dtoEntityMap
}
