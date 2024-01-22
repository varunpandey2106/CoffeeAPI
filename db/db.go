package db

import (
	"database/sql"
	"time"
)

type DB struct{
	DB * sql.DB
}

var dbConn= &DB{}

const maxOpenDbConn = 10
const maxIdleDbConn =5 
const maxDbLifeime= 5* time.Minute

func ConnectPostgres(dsn string) (*DB, error){
	
}