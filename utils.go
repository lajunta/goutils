package goutils

import "reflect"

// Contains check a item in a slice
func Contains(array interface{}, val interface{}) (index int) {
	index = -1
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		{
			s := reflect.ValueOf(array)
			for i := 0; i < s.Len(); i++ {
				if reflect.DeepEqual(val, s.Index(i).Interface()) {
					index = i
					return
				}
			}
		}
	}
	return
}

// Keys get keys  array
func Keys(myMap map[string][]string) []string {
	keys := make([]string, 0, len(myMap))
	for k := range myMap {
		keys = append(keys, k)
	}
	return keys
}

// Values get map values array
func Values(myMap map[string][]string) [][]string {
	keys := make([][]string, 0, len(myMap))
	for _, v := range myMap {
		keys = append(keys, v)
	}
	return keys
}
