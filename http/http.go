package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	Id     int
	Age    int
	Name   string
	Adress string
}

func DatabaseReader(myDb *sql.DB) ([]student, error) {

	rows, err := myDb.Query("SELECT * FROM student")
	if err != nil {
		return nil, err
	}
	var studnet []student
	for rows.Next() {
		var s student
		err = rows.Scan(&s.Id, &s.Age, &s.Name, &s.Adress)
		if err != nil {
			//fmt.Println(err)
			//return nil, err
			//fmt.Println(s)
		}
		studnet = append(studnet, s)

	}
	return studnet, nil
}

func studentdb() []student {
	db, err := sql.Open("mysql", "root:123@Passwo@tcp(localhost:3306)/student")
	//fmt.Println("errr")
	if err != nil {
		//fmt.Println(err)
	}
	defer db.Close()

	list, err := DatabaseReader(db)
	return list

}
func Handler(w http.ResponseWriter, req *http.Request) {

	if req.URL.Path == ("/") {
		w.WriteHeader(http.StatusOK)
		sql := studentdb()
		fmt.Fprintln(w, sql)
	} else if req.URL.Path == "/student" {
		p := studentdb()
		fmt.Fprintln(w, p)
		//w.WriteHeader(http.StatusOK)
	} // else {
	//	fmt.Fprintln(w, req.URL.Path[1:])
	//}

}
func main() {

	err := http.ListenAndServe(":8000", http.HandlerFunc(Handler))
	if err != nil {
		fmt.Println(err)
	}

}

// var db *sql.DB

// func GetPeople(w http.ResponseWriter, r *http.Request) {
// 	json.NewEncoder(w).Encode(db.QueryRow("SELECT * from student"))
// }

// func main() {
// 	var err error
// 	db, err = sql.Open("mysql", "root:123@Passwo@tcp(127.0.0.1:3306)/student")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer db.Close()

// 	router := mux.NewRouter()
// 	router.HandleFunc("/student", GetPeople).Methods("GET")
// 	log.Fatal(http.ListenAndServe(":8000", router))
// }

// func countCalls(h http.HandlerFunc) http.HandlerFunc {
//         var do
//         var count int
//         return func(w http.ResponseWriter, r *http.Request) {
//             do.b()
//             count++
//         }
//     }

// func main() {
// 	http.HandleFunc("/", HelloServer)
// 	http.ListenAndServe(":8000", nil)
// }

// func HelloServer(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
// }

// func (a apihandler) ServeHTTP(b http.ResponseWriter, r *http.Request) {
// 	b.Write([]byte(a))
// }

//func main() {
//var a apihandler = "tpe abc without func"
/*http.HandlerFunc("/abc", func(b http.ResponseWriter, r *http.Request) {
	b.Write([]byte(a))
})

//http.HandleFunc("/", HelloServer)*/
//http.HandleFunc("/a",HandlerFunc)
//http.ListenAndServe(":8000", http.HandlerFunc)   }

// func Write(name string,w http.ResponseWriter){
//     if name=="hello"{
//         w.Write([]byte("Hello"))
//     } else{
//         b:=fmt.Sprintf("Hello %s",name)
//         w.Write([]byte(b))
//     }
// }
// func countCalls(h http.HandlerFunc) http.HandlerFunc {
//     var do a
//     var count int
//     return func(w http.ResponseWriter, r *http.Request) {
//         do.b()
//         count++
//     }
// }

//  func HandlerFunc(w http.ResponseWriter, r *http.Request) {
//     if r.URL.Path =="/"{
//         w.Write([]byte("Hello"))
//     } else{

//         b:=fmt.Sprintf("Hello %s",r.URL.Path)
//         w.Write([]byte(b))
//     }
// }
