package helpers

import "reflect"

func MergeStructs(dest, src interface{}) {
	destValue := reflect.ValueOf(dest).Elem()
	srcValue := reflect.ValueOf(src).Elem()

	for i := 0; i < srcValue.NumField(); i++ {
		srcField := srcValue.Field(i)
		destField := destValue.FieldByName(srcValue.Type().Field(i).Name)
		if destField.IsValid() && destField.CanSet() {
			destField.Set(srcField)
		}
	}
}
