package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Student struct {
	Class   int       `json:"class"`
	RollNum int       `json:"roll_num"`
	Marks   []Subject `json:"marks"`
}

type Subject struct {
	Name     string `json:"name"`
	SubMarks int    `json:"sub_marks"`
}

var stuMap = make(map[string][]Subject)

func storeMarks(w http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodPost {
		var stuMarks Student
		err := json.NewDecoder(req.Body).Decode(&stuMarks)
		if err != nil {
			http.Error(w, "Cannot Decode request", http.StatusBadRequest)
			return
		}
		s := fmt.Sprintf("%d|%d", stuMarks.Class, stuMarks.RollNum)
		//storing marks
		stuMap[s] = stuMarks.Marks
		fmt.Fprintf(w, "Our record:%v", stuMap)

	} else {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}

}

func fetchMarks(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		var stuMarks Student
		err := json.NewDecoder(req.Body).Decode(&stuMarks)
		if err != nil {
			http.Error(w, "Cannot Decode request", http.StatusBadRequest)
			return
		}
		s := fmt.Sprintf("%d|%d", stuMarks.Class, stuMarks.RollNum)
		// fetching a student marks
		fmt.Fprintf(w, "Our Record: %v", stuMap[s])
	} else {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}
}

// func fetchMarks(w http.ResponseWriter, req *http.Request) {
// 	if req.Method == http.MethodGet {
// 		//var stuMarks Student
// 		rollParam := req.URL.Query().Get("roll_num")
// 		fmt.Println(rollParam)
// 		classParam := req.URL.Query().Get("class")
// 		fmt.Println(classParam)
// 		// classDecoder := json.NewDecoder(strings.NewReader(classParam))
// 		// rollDecoder := json.NewDecoder(strings.NewReader(rollParam))
// 		// classDecoder.Decode(&stuMarks)
// 		// rollDecoder.Decode(&stuMarks)
// 		s := fmt.Sprintf("%s|%s", classParam, rollParam)
// 		fmt.Println(s)
// 		fmt.Fprintf(w, "Our Record: %v", stuMap[s])
// 	} else {
// 		http.Error(w, "Bad request", http.StatusBadRequest)
// 	}
// }

func allStudentMarks(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		fmt.Fprintf(w, "Students marks: %v", stuMap)
	} else {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}
}

func main() {

	http.HandleFunc("/storemarks", storeMarks)
	http.HandleFunc("/fetchmarks", fetchMarks)
	http.HandleFunc("/allstudent", allStudentMarks)
	http.ListenAndServe(":8080", nil)

}
