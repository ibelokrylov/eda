package helpers

import (
	"reflect"
	"regexp"
)

type Replacer struct {
	patterns []string
}

func NewReplacer(patterns []string) *Replacer {
	return &Replacer{
		patterns: patterns,
	}
}

func (r *Replacer) Replace(text string) (string, error) {
	for _, pattern := range r.patterns {
		re, err := regexp.Compile(pattern)
		if err != nil {
			return "", err
		}
		text = re.ReplaceAllStringFunc(text, func(s string) string {
			matched := re.FindStringSubmatch(s)
			if len(matched) == 3 {
				return matched[1] + "XXX"
			}
			return s
		})
	}
	return text, nil
}

func RemoveFields(input interface{}, fieldsToRemove []string) interface{} {
	// Получаем отражение типа и значения входной структуры
	val := reflect.ValueOf(input).Elem()
	typ := val.Type()

	// Создаем новую структуру с удаленными полями
	newStruct := reflect.New(typ).Elem()

	// Копируем поля, которых нет в fieldsToRemove, в новую структуру
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if !contains(fieldsToRemove, field.Name) {
			newStruct.Field(i).Set(val.Field(i))
		}
	}

	return newStruct.Interface()
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
