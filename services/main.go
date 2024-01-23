package services

import (
	"time"
)


var db * sql.DB
const dbTimeout= time.Second * 3
