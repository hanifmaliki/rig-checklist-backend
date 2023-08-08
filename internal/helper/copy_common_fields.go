package helper

import (
	"reflect"
)

// CopyCommonFields copies the common fields from the struct
// pointed to src to the struct pointed to by dest.
func CopyCommonFields(destp, srcp interface{}) {
	destv := reflect.ValueOf(destp).Elem()
	srcv := reflect.ValueOf(srcp).Elem()

	destt := destv.Type()
	for i := 0; i < destt.NumField(); i++ {
		sf := destt.Field(i)
		v := srcv.FieldByName(sf.Name)
		if !v.IsValid() || !v.Type().AssignableTo(sf.Type) {
			continue
		}
		destv.Field(i).Set(v)
	}
}
