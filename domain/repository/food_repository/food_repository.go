package food_repository

import (
	entity "github.com/TheEgid/articles-site-go/domain/entity/food_entity"
)

type FoodRepository interface {
	SaveFood(*entity.Food) (*entity.Food, map[string]string)
	GetFood(uint64) (*entity.Food, error)
	GetAllFood() ([]entity.Food, error)
	UpdateFood(*entity.Food) (*entity.Food, map[string]string)
	DeleteFood(uint64) error
}
