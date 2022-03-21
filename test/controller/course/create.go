package course

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"test/db"
	"test/model"
	"time"
)

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	var course model.Course

	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("no date inside json")
	}

	// generate random id, convert it into string

	rand.Seed(time.Now().UnixNano())
	course.CourseID = strconv.Itoa(rand.Intn(100))

	db.Courses = append(db.Courses, course)

	json.NewEncoder(w).Encode(course)

}
