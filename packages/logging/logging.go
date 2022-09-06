package logging

import (
	dbc "dbconn"
	"fmt"
	"runtime"
	"time"

	"gorm.io/gorm"
)

type Logger interface {
	SetLoggerDbHandler(dbc.DbHandler)
	GetLoggerDbHandler() dbc.DbHandler
	InitLoggerPerFunc(string, string)
	SetLogStartTime()
	GetCurrentFuncName() string
	SetCurrentFuncName(string)
	SetInfo(string)
	Log()
	LogError(error)
}

type Model struct {
	ID        uint           `gorm:"primaryKey;autoIncrement"` // works both with postgres and SQL Server
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type UPMLogger struct {
	Model
	FuncName       string    `json:"func_name"`
	Info           string    `json:"info"`
	UserID         uint      `json:"user_id"`
	EventStartTime time.Time `json:"event_start_time"`
	EventEndTime   time.Time `json:"event_end_time"`
	CriticalError  bool      `json:"critical_error"`
	ErrorMessage   string    `json:"error_message"`
	dbHandler      dbc.DbHandler
}

func (ul *UPMLogger) InitLoggerPerFunc(funcName, info string) {
	ul.SetLogStartTime()
	ul.SetCurrentFuncName(funcName)
	ul.SetInfo(info)
}

func (ul *UPMLogger) GetCurrentFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	str := runtime.FuncForPC(pc).Name()
	return str
}

func (ul *UPMLogger) SetCurrentFuncName(currentFuncName string) {
	ul.FuncName = currentFuncName
}

func (ul *UPMLogger) SetLogStartTime() {
	ul.EventStartTime = time.Now()
}

func (ul *UPMLogger) SetInfo(info string) {
	ul.Info = info
}

func (ul *UPMLogger) GetLogger() *UPMLogger {
	return ul
}

func (ul *UPMLogger) Log() {
	defer recoverFromPanic()
	ul.EventEndTime = time.Now()
	fmt.Println("logging at function: " + ul.FuncName + " info: " + ul.Info + " log start time: " + ul.EventStartTime.String() + " ,log end time: " + ul.EventEndTime.String())

	// database implementation
	errSaveLogData := ul.saveLogData()
	if errSaveLogData != nil {
		panic(errSaveLogData.Error())
	}
}

func (ul *UPMLogger) LogError(err error) {
	defer recoverFromPanic()
	ul.EventEndTime = time.Now()
	fmt.Println("logging at function: " + ul.FuncName + " error: " + err.Error())

	// database implementation
	errSaveLogData := ul.saveLogData()
	if errSaveLogData != nil {
		panic(errSaveLogData.Error())
	}
}

func (ul *UPMLogger) SetLoggerDbHandler(dbh dbc.DbHandler) {
	ul.dbHandler = dbh
}

func (ul *UPMLogger) GetLoggerDbHandler() dbc.DbHandler {
	return ul.dbHandler
}

func recoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered ", r)
	}
}

func (ul *UPMLogger) saveLogData() error {
	db := ul.dbHandler.PassConnection()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Omit("id").Create(&ul).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
