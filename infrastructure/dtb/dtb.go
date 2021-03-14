package dtb

import (
	"errors"
	"fmt"
	"github.com/TheEgid/articles-site-go/domain/entity/food_entity"
	"github.com/TheEgid/articles-site-go/domain/repository/food_repository"
	"github.com/TheEgid/articles-site-go/infrastructure/repo_methods"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repositories struct {
	Food food_repository.FoodRepository
	db   *gorm.DB
}

func NewRepositories(DbUser, DbPassword, DbPort, DbHost, DbName string) (*Repositories, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s TimeZone=Asia/Shanghai sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		fmt.Println(errors.Unwrap(err))
	}
	return &Repositories{
		Food: repo_methods.NewFoodRepository(db),
		db:   db,
	}, nil
}

//closes the database connection
func (s *Repositories) Close() error {
	sqlDB, err := s.db.DB()
	if err != nil {
		fmt.Println(errors.Unwrap(err))
	}
	return sqlDB.Close()
}

////This migrate all tables
func (s *Repositories) Automigrate() error {
	return s.db.AutoMigrate(&food_entity.Food{})
}
