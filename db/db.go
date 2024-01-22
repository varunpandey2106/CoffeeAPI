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
	d, err:= sql.Open("pgx", dsn)
	if err !=nil{
		return nil, err
	}

	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetConnMaxIdleTime(maxIdleDbConn)
	d.SetConnMaxLifetime(maxDbLifeime )

	
}