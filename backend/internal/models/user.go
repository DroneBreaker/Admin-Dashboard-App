package models

import "time"

type User struct {
	id        int       `json:"id" xml:"id"`
	FirstName string    `json:"firstName" xml:"firstName"`
	LastName  string    `json:"lasttName" xml:"lastName"`
	Username  string    `json:"username" xml:"username"`
	Email     string    `json:"email" xml:"email"`
	Password  string    `json:"password" xml:"password"`
	CreatedAt time.Time `json:"createdAt" xml:"createdAt"`
}
