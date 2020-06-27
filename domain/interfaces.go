package domain

//CourseRepository to fetch courses from database
type CourseRepository interface {
	GetAll() ([]Course, error)
}
