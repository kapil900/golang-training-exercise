package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
)

type student struct {
	Name   string
	Age    int
	Rollno int
}

type config struct {
	user     string
	password string
	host     string
	port     string
	database string
}
type connector struct {
	sql.DB
}

var data = config{
	user:     "root",
	password: "123@Passwo",
	host:     "localhost",
	port:     "3306",
	database: "student",
}

func Checknil(err error) {
	if err != nil {
		fmt.Println(err)
	}

}
func Connectionsetting(driver, dataSource string) *connector {
	db, err := sql.Open(driver, dataSource)
	if err != nil {
		fmt.Println(err)
	}
	return &connector{*db}
}

func getDataSource(data config) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", data.user, data.password, data.host, data.port, data.database)
}

//defer file.Close()

func WriteFile(student []student, file *os.File) {
	length, err := json.Marshal(student)
	Checknil(err)
	file.WriteString(string(length))
}

// func processData(s []student, con *connector) {
// 	for _, c := range s {
// 		con.WriteFile()
// 	}

// }

func CreateFile(nameoffile string) (filename *os.File) {
	filename, err := os.Create(nameoffile)
	Checknil(err)
	return
}

func recieveData(myDB *sql.DB) ([]student, error) {
	row, err := myDB.Query("Select * FROM student")
	Checknil(err)
	var student1 []student
	for row.Next() {
		var s student
		err = row.Scan(&s.Name, &s.Age, &s.Rollno)
		Checknil(err)
		student1 = append(student1, s)
	}
	return student1, nil
}

func main() {
	a := getDataSource(data)
	b := Connectionsetting("mysql", a)
	//c:=  processData(b)
	fmt.Println(b)

}
