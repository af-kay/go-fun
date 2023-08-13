package reflection

import "reflect"

func NaiveIsNil(o any) bool {
	return o == nil
}

func IsNil(o any) bool {
	if NaiveIsNil(o) {
		return true
	}

	value := reflect.ValueOf(o)
	if value.Kind() != reflect.Ptr {
		return false
	}

	return value.IsNil()
}
