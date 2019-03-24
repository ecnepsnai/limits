// Package limits provides a way to define minimum and maximum length requirements on fields of a struct,
// then check a instance against those rules.
// Add the `min:"<int>"` or `max:"<int>"`  tag to any struct field you wish to enforce.
// Limits only works for strings and slices (of any types, enforcing the count).
package limits

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Check if the given struct has any values that violate the field limits. Returns a single error for the first
// matching field that is too long or short, or nil if all fields are valid.
func Check(obj interface{}) error {
	valueOf := reflect.Indirect(reflect.ValueOf(obj))
	typeOf := valueOf.Type()
	count := typeOf.NumField()
	if count == 0 {
		return nil
	}
	i := 0
	for i < count {
		field := typeOf.FieldByIndex([]int{i})
		i++

		var minStr, maxStr string
		minStr = field.Tag.Get("min")
		maxStr = field.Tag.Get("max")

		if minStr != "" {
			minimum, err := strconv.Atoi(minStr)
			if err != nil {
				panic(err)
			}

			value := valueOf.FieldByName(field.Name)
			isArray := strings.HasPrefix(value.Type().String(), "[]")
			thing := "characters"
			if isArray {
				thing = "elements"
			}
			length := value.Len()

			if length < minimum {
				return fmt.Errorf("Value for '%s' requires at least %d %s", field.Name, minimum, thing)
			}
		}

		if maxStr != "" {
			maximum, err := strconv.Atoi(maxStr)
			if err != nil {
				panic(err)
			}

			value := valueOf.FieldByName(field.Name)
			isArray := strings.HasPrefix(value.Type().String(), "[]")
			thing := "characters"
			if isArray {
				thing = "elements"
			}
			length := value.Len()

			if length > maximum {
				return fmt.Errorf("Value for '%s' exceeds maximum %d %s", field.Name, maximum, thing)
			}
		}
	}

	return nil
}
