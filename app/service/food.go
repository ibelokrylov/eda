package service

import (
	"eda/app/config"
	"eda/app/entities"
	"errors"
	"fmt"
)

func FoodFindByName(n string) (*entities.Food, error) {
	var food entities.Food
	c := config.Db.Unscoped().Find(&food, "name = ?", n)
	if c.Error != nil {
		return &entities.Food{}, c.Error
	}

	return &food, nil
}

func FoodFindById(id int) (*entities.Food, error) {
	var food entities.Food
	c := config.Db.Unscoped().Find(&food, "id = ?", id)
	if c.Error != nil {
		return &entities.Food{}, c.Error
	}
	return &food, nil
}

func FoodSearch(s string, id int64) ([]entities.Food, error) {
	var foods []entities.Food

	err := config.Db.
		Where("tsv @@ to_tsquery(?) OR name LIKE ?", s+"*", "%"+s+"%").
		Where("user_id = ? OR user_id IS NULL OR is_public = true", id).
		Where("deleted_at IS NULL").
		Limit(30).
		Find(&foods).Error
	if err != nil {
		return nil, err
	}

	if len(foods) == 0 {
		return nil, nil
	}

	return foods, nil
}

func ProductSearch(s string) ([]entities.Product, error) {
	var products []entities.Product

	err := config.Db.
		Where("tsv @@ to_tsquery(?) OR LOWER(name) LIKE LOWER(?)", s+"*", "%"+s+"%").
		Where("deleted_at IS NULL").
		Preload("Category").
		Limit(30).
		Find(&products).Error
	if err != nil {
		return nil, err
	}

	if len(products) < 30 {

		var cp []entities.ProductCategory
		err = config.Db.
			Where("tsv @@ to_tsquery(?) OR LOWER(name) LIKE LOWER(?)", s+"*", "%"+s+"%").
			Where("deleted_at IS NULL").
			Preload("Products").
			Find(&cp).Error
		if err != nil {
			return nil, err
		}

		if len(cp) > 0 {
			for _, category := range cp {
				for i := range category.Products {
					newProduct := category.Products[i]
					categoryCopy := entities.ProductCategory{
						ID:        category.ID,
						CreatedAt: category.CreatedAt,
						UpdatedAt: category.UpdatedAt,
						DeletedAt: category.DeletedAt,
						Name:      category.Name,
					}
					newProduct.Category = categoryCopy
					products = append(products, newProduct)
				}
			}
		}

		uniqueProducts := make(map[int64]entities.Product)
		for _, product := range products {
			uniqueProducts[product.ID] = product
		}

		var result []entities.Product
		for _, product := range uniqueProducts {
			result = append(result, product)
		}
		return result, nil
	}

	return products, nil
}

func CreateFood(f entities.CreateFood, userID int64) (entities.Food, error) {
	nf := entities.Food{
		UserID:   &userID,
		IsPublic: f.IsPublic || false,
		Recepit:  f.Receipt,
		Name:     f.Name,
	}

	if f.Info != nil {
		nf.Protein = f.Info.Protein
		nf.Fat = f.Info.Fat
		nf.Carbs = f.Info.Carbs
		nf.Calories = f.Info.Calories
	} else {
		if f.Products != nil {
			if len(*f.Products) == 0 {
				return entities.Food{}, errors.New("createFood@product-must-be-not-empty")
			}
			nf.Products = *f.Products
			nps := new(entities.ProductStat)
			var prl float64 = 0
			for i, pr := range *f.Products {
				if pr.Weight == 0 {
					return entities.Food{}, errors.New("createFood@product-weight-must-be-not-empty")
				}

				cf := pr.Weight / 100

				switch {
				case pr.ProductID != nil:
					fbd := new(entities.Product)
					if err := config.Db.Unscoped().First(&fbd, "id = ?", *pr.ProductID).Error; err != nil {
						return entities.Food{}, errors.New("createFood@product-not-found")
					}
					prl += 1
					nps.Protein += fbd.Protein * cf
					nps.Fat += fbd.Fat * cf
					nps.Carbs += fbd.Carbs * cf
					nps.Calories += fbd.Calories * cf

					nf.Products[i].Name = &fbd.Name
					nf.Products[i].Info = &entities.ProductStat{
						Calories: fbd.Calories,
						Fat:      fbd.Fat,
						Carbs:    fbd.Carbs,
						Protein:  fbd.Protein,
					}

				case pr.Name != nil:
					if pr.Info != nil {
						prl += 1
						nps.Protein += pr.Info.Protein * cf
						nps.Fat += pr.Info.Fat * cf
						nps.Carbs += pr.Info.Carbs * cf
						nps.Calories += pr.Info.Calories * cf
					} else {
						return entities.Food{}, errors.New("createFood@products-in-product-info-must-be-not-empty")
					}
				default:
					return entities.Food{}, errors.New("createFood@productid-required")
				}
				nf.Protein = nps.Protein
				nf.Fat = nps.Fat
				nf.Carbs = nps.Carbs
				nf.Calories = nps.Calories
			}

			nps.Protein = nps.Protein / prl
			nps.Calories = nps.Calories / prl
			nps.Carbs = nps.Carbs / prl
			nps.Fat = nps.Fat / prl
		} else {
			return entities.Food{}, errors.New("createFood@food info must be not empty")
		}
	}

	if err := config.Db.Save(&nf).Error; err != nil {
		return entities.Food{}, err
	}

	config.UpdateTSV(`foods`, `tsv`, `russian`, nf.Name, nf.ID)

	if nf.Recepit != nil {
		err := config.UpdateTSV(`"foods"`, `"receipt_tsv"`, `'russian'`, *nf.Recepit, nf.ID)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	return nf, nil
}
