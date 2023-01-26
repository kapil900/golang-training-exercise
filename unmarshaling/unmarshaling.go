/*package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type person struct {
	Name   string
	Age    int
	Rollno int
	Phone  []string
}

func main() {

	jsonFile, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
	}
	bytevalue, _ := ioutil.ReadAll(jsonFile)
	var person1 []person

	error := json.Unmarshal([]byte(bytevalue), &person1)
	fmt.Println(person1, error)
	var p []person
	var s []person

	content, a := json.Marshal(person1)
	if a != nil {
		fmt.Println(a)
	}

	a = ioutil.WriteFile("userfile.json", content, 0644)
	for i := 0; i < len(person1); i++ {
		if person1[i].Age >= 10 {
			s = append(s, person1[i])
		} else {
			p = append(p, person1[i])
		}
	}

	primary, _ := json.Marshal(p)
	_ = ioutil.WriteFile("primary.json", primary, 0644)
	secondary, _ := json.Marshal(s)
	_ = ioutil.WriteFile("secondary.json", secondary, 0644)

}
*/

//Refactoring based on dependency injection

package refactor

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type student struct {
	Name   string
	Rollno int
	Age    int
	Phone  []string
}

func Checknil(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
func OpenFile(a string) *os.File {
	Opening, err := os.Open(a)
	Checknil(err)
	return Opening
}
func readStudent(r io.Reader) ([]student, error) {
	var student []student
	byteValue, err1 := ioutil.ReadAll(r) //the io.Read
	if err1 != nil {
		return nil, err1
	}
	err := json.Unmarshal(byteValue, &student)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func WriteFile(student []student, file *os.File) {
	length, err := json.Marshal(student)
	Checknil(err)
	file.WriteString(string(length))
	defer file.Close()
}

func CreateFile(nameoffile string) (filename *os.File) {
	filename, err := os.Create(nameoffile)
	Checknil(err)
	return
}
func FilterByAge(student []student) (primary []student, secondary []student) {
	for i := 0; i < len(student); i++ {
		if student[i].Age < 10 {
			primary = append(primary, student[i])
		} else {
			secondary = append(secondary, student[i])
		}
	}
	return
}

/*func main() {
	a := OpenFile("data.json")
	b := readStudent(a)
	c, d := FilterByAge(b)
	e := CreateFile("primary.json")
	f := CreateFile("secondary.json")
	WriteFile(c, e)
	WriteFile(d, f)

}*/
