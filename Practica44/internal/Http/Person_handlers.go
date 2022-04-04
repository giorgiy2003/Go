package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

type person struct {
	Id        int    `json:"id" xml:"id" form:"id" query:"id"`
	Email     string `json:"email" xml:"email" form:"email" query:"email"`
	Phone     string `json:"phone" xml:"phone" form:"phone" query:"phone"`
	FirstName string `json:"firstName" xml:"firstName" form:"firstName" query:"firstName"`
	LastName  string `json:"lastName" xml:"lastName" form:"lastName" query:"lastName"`
}

func (p person) formatStyle() string {
	return fmt.Sprintf("%d. Person name is %s. His lastname is %s. His phone is %s. Also email: %s", p.Id, p.FirstName, p.LastName, p.Phone, p.Email)
}

func home_page(w http.ResponseWriter, r *http.Request)  {
	database, _ := sql.Open("sqlite3", "./persons.db") //Для открытия соединения с базой данных используем функцию sql.Open()
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS persons (Id INTEGER PRIMARY KEY,email TEXT,phone TEXT,firstName TEXT,lastName TEXT)")
	statement.Exec()
	//statement, _ = database.Prepare("INSERT INTO persons (email,phone,firstName,lastName) VALUES (?,?,?,?)")
	//statement.Exec("Masha@email", "+79545454767", "masha", "veselova")
	rows, _ := database.Query("SELECT Id,email,phone,firstName,lastName FROM persons")
	persons := []person{}
	for rows.Next() {
		p := person{}
		rows.Scan(&p.Id, &p.Email, &p.Phone, &p.FirstName, &p.LastName)
		persons = append(persons, p)
	}

	for _, p := range persons {
		fmt.Fprintln(w, p.formatStyle())
	}
}

//GET /person/{id} – возвращает одну модель Person.
func  FindPersonByID(c echo.Context)  error {
	var err error
	p := person{}
	db:= &gorm.DB{} 
	err = db.Debug().Model(person{}).Where("id = ?", p.Id).Take(&p).Error
	if err != nil {
		return db.Error
	}
	if gorm.IsRecordNotFoundError(err) {
		return db.Error
	}
	return db.Error
}


//DELETE /person/{id} – удаляет модель Person
func DeletePerson(c echo.Context)  error {
	p := person{}
	db:= &gorm.DB{} 
	db = db.Debug().Model(&person{}).Where("id = ?", p.Id).Take(&person{}).Delete(&person{})

	if db.Error != nil {
		return db.Error
	}
	return db.Error
}

//PUT /person/{id} – обновляет модель Person
func UpdateAUser(c echo.Context)  error {
	var err error 
	p := person{}
	db:= &gorm.DB{} 
	db = db.Debug().Model(&person{}).Where("id = ?", p.Id).Take(&person{}).UpdateColumns(
	map[string]interface{}{
		"Id":        p.Id,
		"Email":     p.Email,
		"FirstName": p.FirstName,
		"LastName":  p.LastName,
		"update_at": time.Now(),
	},
)
if db.Error != nil {
	return db.Error
}
// Отоброжение обновленного пользователя
err = db.Debug().Model(&person{}).Where("id = ?", p.Id).Take(&p).Error
if err != nil {
	return  err
}
return nil
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.ListenAndServe(":8090", nil)
}

//e.GET("/show", show)
func show(c echo.Context) error {
	// Get Email and Phone and FirstName and LastName from the query string
	Email := c.QueryParam("Email")
	Phone := c.QueryParam("Phone")
	FirstName := c.QueryParam("FirstName")
	LastName := c.QueryParam("LastName")
	return c.String(http.StatusOK, "Email:"+Email+", Phone:"+Phone+", FirstName:"+FirstName+", LastName:"+LastName)
}

// e.POST("/save", save)
func save(c echo.Context) error {
	// Get Email and Phone and FirstName and LastName
	Email := c.FormValue("Email")
	Phone := c.FormValue("Phone")
	FirstName := c.FormValue("FirstName")
	LastName := c.FormValue("LastName")
	return c.String(http.StatusOK, "Email:"+Email+", Phone:"+Phone+", FirstName:"+FirstName+", LastName:"+LastName)
}

//GET /person/ - возвращает список моделей Person о которых у нас есть инфа.
func FindAllPersons(c echo.Context) error {
	var err error
	db:= &gorm.DB{} 
	persons := []person{}
	err = db.Debug().Model(&person{}).Limit(100).Find(&persons).Error
	if err != nil {
		return  err
	}
	return err
}