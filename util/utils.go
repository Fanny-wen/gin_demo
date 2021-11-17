package util

import (
	"reflect"
)

// Struct2Map struct to map
func Struct2Map(u interface{}, size int) map[string]interface{} {
	t := reflect.TypeOf(u)
	v := reflect.ValueOf(u)
	// 指定size, 在 make 中进行初始化大小, 可以更快
	m := make(map[string]interface{}, size)
	for i := 0; i < t.NumField(); i++ {
		m[t.Field(i).Name] = v.Field(i).Interface()
	}
	return m
}

// Map2Struct map to struct
func Map2Struct(beforeMap map[string]interface{}, afterStruct interface{}) interface{} {
	v := reflect.ValueOf(afterStruct)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		if v.Kind() != reflect.Struct {
			panic("must struct")
		}
		findFromMap := func(key string,nameTag string ) interface {} {
			for k,v := range beforeMap {
				if k == key || k == nameTag {
					return v
				}
			}
			return nil
		}
		for i := 0; i < v.NumField(); i++ {
			val := findFromMap(v.Type().Field(i).Name,v.Type().Field(i).Tag.Get("name"))
			if val != nil && reflect.ValueOf(val).Kind() == v.Field(i).Kind() {
				v.Field(i).Set(reflect.ValueOf(val))
			}
		}
	} else {
		panic("must ptr")
	}
	return nil
}
