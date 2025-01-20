package service

import (
	"eda/app/config"
	"eda/app/entities"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

type groupParse struct {
	Href     string
	Name     string
	Calories float64
	Protein  float64
	Fat      float64
	Carbs    float64
	ImageUrl string
}

type productParsed struct {
	Href string
	Name string
}

func getPageDoc(url string) (*goquery.Document, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return goquery.NewDocumentFromReader(resp.Body)
}

func parseCategories() {
	baseUrl := "https://calorizator.ru"
	doc, err := getPageDoc(baseUrl + "/product")
	if err != nil {
		log.Fatalf("parseUrl@failed. err: %v", err)
	}

	products := []entities.ProductParesed{}
	exc := []string{"product/brand", "product/pix", "product/choice", "product/all"}

	doc.Find("ul.product a").Each(func(index int, element *goquery.Selection) {
		href, exists := element.Attr("href")
		if exists && strings.Contains(href, "product") {
			skip := false
			for _, excItem := range exc {
				if href == excItem {
					skip = true
					break
				}
			}
			if skip {
				return
			}

			text := strings.TrimSpace(element.Text())

			CategoryCreate(&entities.CreateProductCategory{Name: text})

			products = append(products, entities.ProductParesed{
				Url:          baseUrl + "/" + href,
				Name:         text,
				CountPage:    0,
				ProductCount: 0,
			})
		}
	})
	config.Db.Create(products)
}

func ParseProduct() {
	cat := []entities.ProductParesed{}

	ca := config.Db.Find(&cat)
	if ca.Error != nil {
		fmt.Println(ca.Error)
		return
	}

	if len(cat) > 0 {
		f := new(entities.ProductParesed)
		fa := config.Db.Where("product_count = ?", 0).Find(&f)

		if fa.Error != nil {
			fmt.Println(fa.Error)
			return
		}

		if f.ID == 0 {
			return
		}
		categories := new(entities.ProductCategory)
		cf := config.Db.Unscoped().Find(&categories, "name = ?", f.Name)

		if cf.Error != nil {
			fmt.Println(cf.Error)
			return
		}

		prs, cnt, err := parseGroupProductPagePagination(f.Url, categories.ID)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		go ProductCreateMany(prs)
		f.CountPage = cnt
		f.ProductCount = int32(len(prs))
		config.Db.Save(&f)
	} else {
		parseCategories()
		ParseProduct()
	}
}

func parseGroupProductPagePagination(url string, categoryId int64) ([]entities.CreateProduct, int32, error) {
	doc, err := getPageDoc(url)
	if err != nil {
		return nil, 0, err
	}

	countPage := 0
	doc.Find("ul.pager li").Each(func(int, *goquery.Selection) {
		countPage = countPage + 1
	})

	p := []entities.CreateProduct{}

	if countPage == 0 {
		pr, err := parsePage(0, url, doc, categoryId)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			p = append(p, pr...)
		}
	} else {
		for i := range countPage {
			pr, err := parsePage(i, url, doc, categoryId)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				p = append(p, pr...)
			}
		}
	}
	return p, int32(countPage + 1), nil
}

func parsePage(page int, url string, doc *goquery.Document, categoryId int64) ([]entities.CreateProduct, error) {
	if page != 0 {
		p := strconv.Itoa(page)
		url = url + "?page=" + p
		doc, _ = getPageDoc(url)
	}

	var wg sync.WaitGroup
	parsed := make([]*goquery.Selection, 0)
	var mutex sync.Mutex

	// Находим все строки таблицы
	doc.Find("tbody tr.odd, tbody tr.even").Each(func(i int, el *goquery.Selection) {
		wg.Add(1)
		go func(el *goquery.Selection) {
			defer wg.Done()
			mutex.Lock()
			parsed = append(parsed, el)
			mutex.Unlock()
		}(el)
	})

	wg.Wait() // Ждем завершения всех горутин, которые собирают элементы

	results := make([]entities.CreateProduct, 0, len(parsed))
	var parseWg sync.WaitGroup

	for _, el := range parsed {
		parseWg.Add(1)
		go func(el *goquery.Selection) {
			defer parseWg.Done()
			pr := new(entities.CreateProduct)
			pr.CategoryID = categoryId

			s, exists := el.Find("td.views-field-field-picture-fid a img").Attr("src")
			if exists {
				pr.ImageUrl = s
			}

			pr.Name = strings.TrimSpace(el.Find("td.views-field-title a").Text())

			prot := strings.TrimSpace(el.Find("td.views-field-field-protein-value").Text())
			if prot != "" {
				if f, err := strconv.ParseFloat(prot, 64); err == nil {
					pr.Protein = f
				}
			}

			fat := strings.TrimSpace(el.Find("td.views-field-field-fat-value").Text())
			if fat != "" {
				if f, err := strconv.ParseFloat(fat, 64); err == nil {
					pr.Fat = f
				}
			}

			carb := strings.TrimSpace(el.Find("td.views-field-field-carbohydrate-value").Text())
			if carb != "" {
				if f, err := strconv.ParseFloat(carb, 64); err == nil {
					pr.Carbs = f
				}
			}

			cal := strings.TrimSpace(el.Find("td.views-field-field-kcal-value").Text())
			if cal != "" {
				if f, err := strconv.ParseFloat(cal, 64); err == nil {
					pr.Calories = f
				}
			}

			if pr.Name != "" {
				mutex.Lock()
				results = append(results, *pr)
				mutex.Unlock()
			}
		}(el)
	}

	parseWg.Wait() // Ждем завершения всех горутин, которые парсят элементы

	return results, nil
}

func ProductFindById(id int) (*entities.Product, error) {
	var product entities.Product
	c := config.Db.Unscoped().Find(&product, "id = ?", id)
	if c.Error != nil {
		return &entities.Product{}, c.Error
	}

	return &product, nil
}

func ProductFindByName(n string) (*entities.Product, error) {
	var product entities.Product
	c := config.Db.Unscoped().Find(&product, "name = ?", n)
	if c.Error != nil {
		return &entities.Product{}, c.Error
	}

	return &product, nil
}

func ProductCreateMany(prs []entities.CreateProduct) ([]entities.Product, error) {
	var products []entities.Product
	var failedProducts []entities.CreateProduct
	for _, pr := range prs {
		product := entities.Product{
			Name:       pr.Name,
			Calories:   pr.Calories,
			CategoryID: pr.CategoryID,
			Protein:    pr.Protein,
			ImageUrl:   pr.ImageUrl,
			Fat:        pr.Fat,
			Carbs:      pr.Carbs,
		}
		tx := config.Db.Begin()
		if err := tx.Create(&product).Error; err != nil {
			tx.Rollback()
			failedProducts = append(failedProducts, pr)
			continue
		}

		if err := config.UpdateTSV("products", "tsv", `'russian'`, `coalesce(name, '')`, product.ID); err != nil {
			tx.Rollback()
			failedProducts = append(failedProducts, pr)
			continue
		}

		if err := tx.Commit().Error; err != nil {
			failedProducts = append(failedProducts, pr)
			continue
		}
		products = append(products, product)
	}

	if len(failedProducts) > 0 {
		return products, fmt.Errorf("some products could not be created: %d failures", len(failedProducts))
	}

	return products, nil
}
