package model

import "time"

type User struct {
	ID      string
	Name    string
	Created time.Time
	Updated time.Time
}
