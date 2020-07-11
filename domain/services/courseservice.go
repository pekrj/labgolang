package services

import (
	"github.com/pekrj/labgolang/domain"
)



type CourseService struct {
	CourseRepository domain.CourseRepository
}

func (cs *CourseService) GetAllCourses() ([]domain.Course, error) {
	return cs.CourseRepository.GetAll()
}
