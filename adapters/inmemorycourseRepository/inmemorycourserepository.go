package inmemorycourseRepository

import "github.com/pekrj/labgolang/domain"

type InMemoryCourseRepository struct{}

func (ir *InMemoryCourseRepository) GetAll() ([]domain.Course, error) {
	return []domain.Course{
		domain.Course{"123", "Teste 1", "Descricao 1", 10},
		domain.Course{"456", "Teste 2", "Descricao 2", 18},
		domain.Course{"789", "Teste 3", "Descricao 3", 7},
		domain.Course{"012", "Teste 4", "Descricao 4", 2},
	}, nil
}
