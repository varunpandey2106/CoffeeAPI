package db

import (
	"database/sql"
<<<<<<< HEAD
	"time"
=======
>>>>>>> d7159c2d6a77617532346b74bc9023f116fb5cef
)

type DB struct{
	DB * sql.DB
<<<<<<< HEAD
}

var dbConn= &DB{}

const maxOpenDbConn = 10
const maxIdleDbConn =5 
const maxDbLifeime= 5* time.Minute

func ConnectPostgres(dsn string) (*DB, error){
	
=======
>>>>>>> d7159c2d6a77617532346b74bc9023f116fb5cef
}