package main

import (
	_ "github.com/mattn/go-sqlite3"
)

type person struct {
	Id        int    `json:"id" xml:"id" form:"id" query:"id"`
	Email     string `json:"email" xml:"email" form:"email" query:"email"`
	Phone     string `json:"phone" xml:"phone" form:"phone" query:"phone"`
	FirstName string `json:"firstName" xml:"firstName" form:"firstName" query:"firstName"`
	LastName  string `json:"lastName" xml:"lastName" form:"lastName" query:"lastName"`
}