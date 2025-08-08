package homework9

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Person struct {
	Name    string `properties:"name"`
	Address string `properties:"address,omitempty"`
	Age     int    `properties:"age"`
	Married bool   `properties:"married"`
}
type PersonEnh struct {
	Person    Person `properties:"person"`
	HasFamily bool   `properties:"family"`
}

func Serialize[T any](person T) string {
	sb := new(strings.Builder)
	dataType := reflect.TypeOf(person)
	dataValue := reflect.ValueOf(person)
	fieldsCount := dataType.NumField()

	for i := 0; i < fieldsCount; i++ {
		field := dataType.Field(i)
		props, exist := field.Tag.Lookup("properties")
		if !exist {
			continue
		}
		if props == "" {
			continue
		}
		parts := strings.Split(props, ",")
		omitempty := false
		name := parts[0]
		if len(parts) > 1 {
			if parts[1] == "omitempty" {
				omitempty = true
			}
		}
		fieldValue := dataValue.Field(i)

		if isEmptyValue(fieldValue) && omitempty {
			continue
		}
		sb.WriteString(fmt.Sprintf("%s=%s", name, getStringValue(fieldValue)))
		if i < fieldsCount-1 {
			sb.WriteString("\n")
		}
	}

	return sb.String()
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64,
		reflect.Interface, reflect.Pointer:
		return v.IsZero()
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if !isEmptyValue(v.Field(i)) {
				return false
			}
		}
		return true
	}
	return false
}
func getStringValue(v reflect.Value) string {
	if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		if v.IsNil() {
			return "null"
		}
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.String:
		return v.String()
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'g', -1, 64)
	case reflect.Struct:
		return "{" + Serialize(v.Interface()) + "}"
	case reflect.Slice, reflect.Array:
		var elements []string
		for i := 0; i < v.Len(); i++ {
			elements = append(elements, getStringValue(v.Index(i)))
		}
		return "[" + strings.Join(elements, ",") + "]"
	case reflect.Map:
		var pairs []string
		for _, key := range v.MapKeys() {
			pairs = append(pairs, fmt.Sprintf("%s:%s",
				getStringValue(key),
				getStringValue(v.MapIndex(key))))
		}
		return "{" + strings.Join(pairs, ",") + "}"
	default:
		return fmt.Sprintf("%v", v.Interface())
	}
}
