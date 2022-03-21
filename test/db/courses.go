package db

import "test/model"

var LastID = 1

var Courses = []model.Course{
	{CourseID: "1", CourseName: "ReactJS course", Author: &model.Author{Fullname: "Ajay Singh", Website: "example.com"}, CoursePrice: 299},
}
