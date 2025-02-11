package models

import "time"

type User struct {
	ID        int       `json:"id" xml:"id"`
	FirstName string    `json:"firstName" xml:"firstName"`
	LastName  string    `json:"lastName" xml:"lastName"`
	Username  string    `json:"username" xml:"username"`
	Email     string    `json:"email" xml:"email"`
	Password  string    `json:"-" xml:"-"`
	CreatedAt time.Time `json:"createdAt" xml:"createdAt"`
}
