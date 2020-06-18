package db

import (
	_ "database/sql"
)


type Migration struct {
	Version     int
	Description string
	Script      string
}
