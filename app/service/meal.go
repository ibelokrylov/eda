package service

import (
	"eda/app/config"
	"eda/app/entities"
	"errors"
	"strconv"
	"time"
)

func GetMealsByDay(d time.Time) ([]entities.Meal, error) {
	meal := []entities.Meal{}
	if err := config.Db.Where("day = ?", d).Find(&meal).Error; err != nil {
		return []entities.Meal{}, err
	}
	return meal, nil
}

func CreateMeal(m *entities.CreateMeal, id int64, d time.Time) (*entities.Meal, error) {
	if len(m.Food) == 0 {
		return nil, errors.New("create-meal@food empty")
	}
	fm := new(entities.Meal)
	config.Db.Where("user_id = ?", id).Where("day = ?", d).Where("meal_type = ?", m.MealType).Find(&fm)
	info, f, w, err := CalculateStatMeal(m.Food)
	if err != nil {
		return nil, err
	}
	meal := entities.Meal{
		Day:       d,
		UserID:    id,
		Info:      info,
		MealFoods: f,
		MealType:  m.MealType,
		Weight:    w,
	}

	if fm.ID != 0 {
		config.Db.Model(&fm).Updates(&meal)
		meal.ID = fm.ID
		meal.CreatedAt = fm.CreatedAt
		meal.UpdatedAt = fm.UpdatedAt
	} else {
		if err := config.Db.Create(&meal).Error; err != nil {
			return nil, err
		}
	}
	return &meal, nil
}

func CalculateStatMeal(sm []entities.MealFood) (entities.ProductStat, []entities.MealFood, float64, error) {
	grouped := make(map[string]*entities.MealFood)
	var m []entities.MealFood

	// Перебираем входной срез
	for _, mealFood := range sm {
		if mealFood.Weight == 0 {
			return entities.ProductStat{}, nil, 0, errors.New("calculate-stat-meal@weight-must-be-not-empty")
		}
		id := strconv.FormatInt(*mealFood.Id, 10) + mealFood.Type
		if existing, ok := grouped[id]; ok {
			existing.Weight += mealFood.Weight
		} else {
			grouped[id] = &entities.MealFood{
				Id:     mealFood.Id,
				Weight: mealFood.Weight,
				Type:   mealFood.Type,
			}
		}
	}

	for _, mf := range grouped {
		m = append(m, *mf)
	}

	stat := new(entities.ProductStat)
	var fw float64

	// Перебираем сгруппированные элементы
	for id, i := range m {
		var err error
		var name string
		k := i.Weight / 100

		if i.Type == "product" {
			p := new(entities.Product)
			err = config.Db.Where("id = ?", i.Id).Find(p).Error
			name = p.Name
			stat.Calories += p.Calories * k
			stat.Protein += p.Protein * k
			stat.Fat += p.Fat * k
			stat.Carbs += p.Carbs * k

			m[id].Info = &entities.ProductStat{
				Protein:  p.Protein,
				Fat:      p.Fat,
				Carbs:    p.Carbs,
				Calories: p.Calories,
			}

		} else {
			f := new(entities.Food)
			err = config.Db.Where("id = ?", i.Id).Find(f).Error
			name = f.Name

			stat.Calories += f.Calories * k
			stat.Protein += f.Protein * k
			stat.Fat += f.Fat * k
			stat.Carbs += f.Carbs * k

			m[id].Info = &entities.ProductStat{
				Protein:  f.Protein,
				Fat:      f.Fat,
				Carbs:    f.Carbs,
				Calories: f.Calories,
			}

		}

		if err != nil {
			return entities.ProductStat{}, nil, 0, err
		}

		m[id].Name = &name

		fw += i.Weight
	}

	return entities.ProductStat{
		Calories: stat.Calories,
		Protein:  stat.Protein,
		Fat:      stat.Fat,
		Carbs:    stat.Carbs,
	}, m, fw, nil
}

func GetMealByUserId(id int64, d time.Time) ([]entities.Meal, error) {
	meal := []entities.Meal{}
	if err := config.Db.Where("day = ?", d).Where("user_id = ?", id).Find(&meal).Error; err != nil {
		return nil, err
	}

	return meal, nil
}
