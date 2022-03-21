package course

import (
	"net/http"
	"test/db"

	"github.com/gorilla/mux"
)

func Delete(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	for index, course := range db.Courses {
		if course.CourseID == params["id"] {
			db.Courses = append(db.Courses[:index], db.Courses[index+1:]...)
			w.Write([]byte(`
				{
					success: true,
					msg: "course deleted"
				}
			`))
			return
		}
	}
	w.Write([]byte("no course with id: " + params["id"]))
}
