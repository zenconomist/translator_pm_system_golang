package dbconn

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type DbHandler interface {
	SetDbConn()
	PassConnection() *gorm.DB
}

type SqlServerDb struct {
	Db *gorm.DB
}

type TestSqlServerDb struct {
	Db *gorm.DB
}

type PostgreDb struct {
	Db *gorm.DB
}

type TestPostgreDb struct {
	Db *gorm.DB
}

func (dbType *SqlServerDb) SetDbConn() {
	var err error
	dsn := "server=; user id = ;password=;port=;database="
	dbType.Db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		// DisableForeignKeyConstraintWhenMigrating: true,
		// NamingStrategy: schema.NamingStrategy{
		// 	TablePrefix: "gdp_", // as for GoDesignPatterns
		// },
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("connection to the db is set")
}

func (dbType *TestSqlServerDb) SetDbConn() {
	var err error
	dsn := "server=; user id = ;password=;port=;database="
	dbType.Db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		// DisableForeignKeyConstraintWhenMigrating: true,
		// NamingStrategy: schema.NamingStrategy{
		// 	TablePrefix: "gdp_", // as for GoDesignPatterns
		// },
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("connection to the db is set")
}

func (dbType *PostgreDb) SetDbConn() {
	// var err error
	dsn := "host= user= password= dbname= port= sslmode="
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	dbType.Db = db
	fmt.Println("connection to the db is set")
}

func (dbType *TestPostgreDb) SetDbConn() {
	// var err error
	dsn := "host= user= password= dbname= port= sslmode="
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	dbType.Db = db
	fmt.Println("connection to the db is set")
}

func (dbType SqlServerDb) PassConnection() *gorm.DB {
	return dbType.Db
}

func (dbType PostgreDb) PassConnection() *gorm.DB {
	return dbType.Db
}

func (dbType TestPostgreDb) PassConnection() *gorm.DB {
	return dbType.Db
}

func (dbType TestSqlServerDb) PassConnection() *gorm.DB {
	return dbType.Db
}
