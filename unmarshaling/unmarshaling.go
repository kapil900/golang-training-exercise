package main

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
