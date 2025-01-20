package service

import (
	"eda/app/config"
	"eda/app/entities"
	"fmt"
)

func CategoryFindByName(n string) (*entities.ProductCategory, error) {
	var category entities.ProductCategory
	c := config.Db.Unscoped().Find(&category, "name = ?", n)
	if c.Error != nil {
		return &entities.ProductCategory{}, c.Error
	}

	return &category, nil
}

func CategoryFindById(id int) (*entities.ProductCategory, error) {
	var category entities.ProductCategory
	c := config.Db.Unscoped().Find(&category, "id = ?", id)
	if c.Error != nil {
		return &entities.ProductCategory{}, c.Error
	}

	return &category, nil
}

func CategoryCreate(c *entities.CreateProductCategory) (*entities.ProductCategory, error) {
	n := &entities.ProductCategory{
		Name: c.Name,
	}
	cr := config.Db.Create(&n)
	if cr.Error != nil {
		// Проверка на уже существующую категорию с таким именем
		f := config.Db.Unscoped().Find(&n, "name = ?", c.Name)
		if f.Error != nil {
			return &entities.ProductCategory{}, cr.Error
		}
		return n, nil // Возвращаем найденную категорию
	}

	if err := config.UpdateTSV("product_categories", "tsv", `'russian'`, `coalesce(name, '')`, n.ID); err != nil {
		return &entities.ProductCategory{}, fmt.Errorf("failed to update tsv for category: %v", err)
	}

	return n, nil
}
