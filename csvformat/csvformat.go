package main

import (
	"database/sql"
	"fmt"
	"io"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	Name   string
	Age    int
	Rollno int
}

func openConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:Passwo@123@tcp(localhost:3306)/student")
	if err != nil {
		fmt.Println(err)
	}
	return db, nil
}

func recieveData(db *sql.DB) (s []student) {
	rows, _ := db.Query("SELECT * FROM student")
	defer rows.Close()
	for rows.Next() {
		var student1 student
		rows.Scan(&student1)
		s = append(s, student1)
	}
	return s
}

func createFile(filename string) *os.File {
	f, _ := os.Create(filename)
	return f
}

func writeFile(s student, w io.Writer) error {
	a := fmt.Sprintf("%s, %d, %d \n", s.Name, s.Age, s.Rollno)
	_, err := w.Write([]byte(a))
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func recieved(s []student, f *os.File) {
	for _, c := range s {
		writeFile(c, f)
	}
}

func main() {
	db, _ := openConnection()
	data := recieveData(db)
	f := createFile("file.txt")
	recieved(data, f)

}
