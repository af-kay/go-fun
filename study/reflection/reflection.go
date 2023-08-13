package reflection

import "reflect"

func NaiveIsNil(o interface{}) bool {
	return o == nil
}

func IsNil(o interface{}) bool {
	if o == nil {
		return true
	}

	value := reflect.ValueOf(o)
	if value.Kind() != reflect.Ptr {
		return false
	}

	return value.IsNil()
}
