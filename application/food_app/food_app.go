package food_app

import (
	"github.com/TheEgid/articles-site-go/domain/entity/food_entity"
	"github.com/TheEgid/articles-site-go/domain/repository/food_repository"
)

type foodApp struct {
	fr food_repository.FoodRepository
}

var _ FoodAppInterface = &foodApp{}

type FoodAppInterface interface {
	SaveFood(*food_entity.Food) (*food_entity.Food, map[string]string)
	GetAllFood() ([]food_entity.Food, error)
	GetFood(uint64) (*food_entity.Food, error)
	UpdateFood(*food_entity.Food) (*food_entity.Food, map[string]string)
	DeleteFood(uint64) error
}

func (f *foodApp) SaveFood(food *food_entity.Food) (*food_entity.Food, map[string]string) {
	return f.fr.SaveFood(food)
}

func (f *foodApp) GetAllFood() ([]food_entity.Food, error) {
	return f.fr.GetAllFood()
}

func (f *foodApp) GetFood(foodId uint64) (*food_entity.Food, error) {
	return f.fr.GetFood(foodId)
}

func (f *foodApp) UpdateFood(food *food_entity.Food) (*food_entity.Food, map[string]string) {
	return f.fr.UpdateFood(food)
}

func (f *foodApp) DeleteFood(foodId uint64) error {
	return f.fr.DeleteFood(foodId)
}
