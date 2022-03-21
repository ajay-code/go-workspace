package course

import (
	"encoding/json"
	"net/http"
	"test/db"

	"github.com/gorilla/mux"
)

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	params := mux.Vars(r)

	for index, c := range db.Courses {
		if c.CourseID == params["id"] {
			db.Courses = append(db.Courses[:index], db.Courses[index+1:]...)

			// var course models.Course
			_ = json.NewDecoder(r.Body).Decode(&c)
			c.CourseID = params["id"]
			db.Courses = append(db.Courses, c)
			json.NewEncoder(w).Encode(c)
			return
		}
	}

	json.NewEncoder(w).Encode("no course with id: " + params["id"])

}
