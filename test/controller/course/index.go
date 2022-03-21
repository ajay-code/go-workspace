package course

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"test/db"
	"test/model"

	"github.com/gorilla/mux"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(db.Courses)
}

func AddOne(w http.ResponseWriter, r *http.Request) {
	db.LastID = db.LastID + 1
	db.Courses = append(db.Courses, model.Course{CourseID: strconv.Itoa(db.LastID), CourseName: "NodeJS course", Author: &model.Author{Fullname: "Ajay Singh", Website: "example.com"}, CoursePrice: 199})
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(db.Courses)
}

func Get(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	courseID := params["id"]

	for _, course := range db.Courses {
		if course.CourseID == courseID {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode(fmt.Sprintf("no course with given id: %v", courseID))

}

func Remove(w http.ResponseWriter, r *http.Request) {}
