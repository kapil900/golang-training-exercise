package main

import (
	"database/sql"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
)

func TestHandler(t *testing.T) {
	a := httptest.NewServer(http.HandlerFunc(Handler))
	resp, err := http.Get(a.URL)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 8000 but got %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	expected := string(`[{1 10 a avvbbc} {2 11 b avdas } {3 4 d svdsa}]`)
	b, err := ioutil.ReadAll(resp.Body)
	output := strings.Replace(string(b), "\n", "", -1)
	if err != nil {
		t.Error(err)
	}
	if string(b) != expected {
		t.Errorf("expected %s and got %s", expected, output)
	}
}

func TestDatabaseReader(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error("Failed to connect with database")
	}

	rs := mock.NewRows([]string{"Id", "Age", "Name", "Adress"}).AddRow(1, 10, "a", "avvbbc")
	mock.ExpectQuery("SELECT *").WillReturnRows(rs)

	tests := []struct {
		name     string
		database *sql.DB
		want     []student
		wantErr  error
	}{
		{name: "Pass", database: db, want: []student{{Id: 1, Age: 10, Name: "a", Adress: "avvbbc"}}, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DatabaseReader(tt.database)
			if err != tt.wantErr {
				t.Errorf("databaseReader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("databaseReader() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 	type args struct {
// 		w   http.ResponseWriter
// 		req *http.Request
// 	}
// 	tests := []struct {
// 		name string
// 		want string
// 	}{
// 		{name : "true" ,
// 		want : string(`[{1 10 "a" "avvbbc"} {2 11 "b" "avdas" } {3 4 "d" "svdsa"}]`),
// 	},
// }

// 	for _, tt := range tests {

// 		t.Run(tt.name, func(t *testing.T) {
// 			Handler(tt.args.w, tt.args.req)
// 		//})
//}
