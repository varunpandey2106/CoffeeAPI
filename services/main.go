package services

import (
	"time"
	"database/sql"
)


var db * sql.DB
const dbTimeout= time.Second * 3
