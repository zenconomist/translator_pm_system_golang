module main

go 1.18

replace (
	codegenerator => ./codegenerator
	dbconn => ./packages/dbconn
	dbq => ./packages/dbq
	dto => ./packages/dto
	entities => ./packages/entities
	environment => ./packages/environment
	globalconstants => ./packages/globalconstants
	logging => ./packages/logging
	msgraph => ./packages/msgraph
	readfiles => ./packages/readfiles
	server => ./packages/server
	services => ./packages/services
)

require dbconn v0.0.0

require dbq v0.0.0-00010101000000-000000000000

require (
	codegenerator v0.0.0
	entities v0.0.0
	environment v0.0.0
	globalconstants v0.0.0 // indirect
	logging v0.0.0 // indirect
	msgraph v0.0.0 // indirect
	server v0.0.0
)

require (
	dto v0.0.0 // indirect
	github.com/360EntSecGroup-Skylar/excelize v1.4.1 // indirect
	github.com/denisenkom/go-mssqldb v0.12.0 // indirect
	github.com/felixge/httpsnoop v1.0.1 // indirect
	github.com/golang-sql/civil v0.0.0-20190719163853-cb61b32ac6fe // indirect
	github.com/golang-sql/sqlexp v0.0.0-20170517235910-f1bb20e5a188 // indirect
	github.com/gorilla/handlers v1.5.1 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.12.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.11.0 // indirect
	github.com/jackc/pgx/v4 v4.16.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/text v0.3.7 // indirect
	gorm.io/driver/postgres v1.3.7 // indirect
	gorm.io/driver/sqlserver v1.3.2 // indirect
	gorm.io/gorm v1.23.7 // indirect
	readfiles v0.0.0-00010101000000-000000000000 // indirect
	services v0.0.0 // indirect
)
