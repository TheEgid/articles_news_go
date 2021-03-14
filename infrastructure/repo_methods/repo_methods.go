package repo_methods

import (
	"errors"
	"github.com/TheEgid/articles-site-go/domain/entity/food_entity"
	"github.com/TheEgid/articles-site-go/domain/repository/food_repository"
	"gorm.io/gorm"
	"strings"
)

type FoodRepo struct {
	db *gorm.DB
}

func NewFoodRepository(db *gorm.DB) *FoodRepo {
	return &FoodRepo{db}
}

//FoodRepo implements the repository.FoodRepository interface
var _ food_repository.FoodRepository = &FoodRepo{}

func (r *FoodRepo) GetFood(id uint64) (*food_entity.Food, error) {
	var food food_entity.Food
	err := r.db.Debug().Where("id = ?", id).Take(&food).Error
	if err != nil {
		return nil, errors.New("database error, please try again")
	}
	errors.Is(err, gorm.ErrRecordNotFound)
	return &food, nil
}

func (r *FoodRepo) SaveFood(food *food_entity.Food) (*food_entity.Food, map[string]string) {
	dbErr := map[string]string{}
	//The images are uploaded to digital ocean spaces. So we need to prepend the url. This might not be your use case, if you are not uploading image to Digital Ocean.
	//food.FoodImage = os.Getenv("DO_SPACES_URL") + food.FoodImage

	err := r.db.Debug().Create(&food).Error
	if err != nil {
		//since our title is unique
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			dbErr["unique_title"] = "food title already taken"
			return nil, dbErr
		}
		//any other db error
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	return food, nil
}

func (r *FoodRepo) GetAllFood() ([]food_entity.Food, error) {
	var foods []food_entity.Food
	err := r.db.Debug().Limit(100).Order("created_at desc").Find(&foods).Error
	if err != nil {
		return nil, err
	}
	errors.Is(err, gorm.ErrRecordNotFound)
	return foods, nil
}

func (r *FoodRepo) UpdateFood(food *food_entity.Food) (*food_entity.Food, map[string]string) {
	dbErr := map[string]string{}
	err := r.db.Debug().Save(&food).Error
	if err != nil {
		//since our title is unique
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			dbErr["unique_title"] = "title already taken"
			return nil, dbErr
		}
		//any other db error
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	return food, nil
}

func (r *FoodRepo) DeleteFood(id uint64) error {
	var food food_entity.Food
	err := r.db.Debug().Where("id = ?", id).Delete(&food).Error
	if err != nil {
		return errors.New("database error, please try again")
	}
	return nil
}
