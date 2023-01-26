package refactor

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

type testReadFail int

func (t testReadFail) Read(b []byte) (int, error) {
	return 0, errors.New("done")
}
func TestReadStudentFail(t *testing.T) {
	_, err := readStudent(testReadFail(6))
	if err.Error() != "done" {
		t.Errorf("Cant read student bad format")
	}
}

func TestreadStudent(t *testing.T) {
	testCases := []struct {
		input   string
		eoutput []student
	}{
		{input: `[{"Name":"a","Rollno":1,"Age":10,"Phone":[15000]},{"Name":"b","Rollno":2,"Age":11,"Phone":[00928392]}]`, eoutput: []student{{Name: "a", Rollno: 1, Age: 10, Phone: []string{"6"}}, {Name: "b", Rollno: 2, Age: 11, Phone: []string{"89348"}}}},
		{input: `[{"Name":"c","Rollno":3,"Age":3,"Phone":[15]},{"Name":"d","Rollno":4,"Age":4,"Phone":[898938]}] `, eoutput: []student{{Name: "a", Rollno: 1, Age: 1, Phone: []string{"15555"}}, {Name: "d", Rollno: 2, Age: 11, Phone: []string{"89348"}}}},
	}
	for _, tt := range testCases {
		got, _ := readStudent(strings.NewReader(tt.input))
		if !reflect.DeepEqual(got, tt.eoutput) {
			t.Errorf("readstudent() = %v,expected output %v", got, tt.eoutput)
		}
	}

}

/*func TestCreateFile()(t*testing.T) {
	a :=CreateFile("file.json")
	if _,err:=os.

}*/

func TestWriteFile(t *testing.T) {
	a := CreateFile("file.json")
	data := []student{{Name: "a", Rollno: 1, Age: 10, Phone: []string{"15000"}}, {Name: "b", Rollno: 2, Age: 11, Phone: []string{"00928392"}}}
	WriteFile(data, a)
	b := OpenFile("file.json")
	f, _ := readStudent(b)
	if !reflect.DeepEqual(f, data) {
		t.Errorf("Error")
	}
}

func TestFilterByAge(t *testing.T) {
	type args struct {
		student []student
	}
	tests := []struct {
		name          string
		args          args
		wantPrimary   []student
		wantSecondary []student
	}{
		//{input:=}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPrimary, gotSecondary := FilterByAge(tt.args.student)
			if !reflect.DeepEqual(gotPrimary, tt.wantPrimary) {
				t.Errorf("FilterByAge() gotPrimary = %v, want %v", gotPrimary, tt.wantPrimary)
			}
			if !reflect.DeepEqual(gotSecondary, tt.wantSecondary) {
				t.Errorf("FilterByAge() gotSecondary = %v, want %v", gotSecondary, tt.wantSecondary)
			}
		})
	}
}
