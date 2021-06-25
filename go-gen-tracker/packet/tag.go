package packet

import (
	"reflect"
	"strconv"

	"github.com/pudjamansyurin/go-gen-tracker/util"
)

// type M map[string]interface{}

type Tag struct {
	Tipe      string
	Unit      string
	Scale     float32
	Chartable bool
}

func GetTag(packet interface{}) {
	tagStructWalk(reflect.ValueOf(packet).Elem())
}

func tagStructWalk(v reflect.Value) {
	switch v.Type().Kind() {
	case reflect.Struct:
		util.Debug(v.Type().Kind())
		for i := 0; i < v.NumField(); i++ {
			value := v.Field(i)
			t := v.Type().Field(i)

			if _, ok := t.Tag.Lookup("type"); !ok {
				// if t.Type.Kind() == reflect.Struct {
				tagStructWalk(value)
				continue
			}
		}

	case reflect.Array:
		// util.Debug(v)
		for i := 0; i < v.Len(); i++ {
			item := v.Index(i)
			util.Debug(item)

			if item.Kind() == reflect.Struct {
				v := reflect.Indirect(item)
				tagStructWalk(v)
				continue
				// for j := 0; j < v.NumField(); j++ {
				// 	fmt.Println(v.Type().Field(j).Name, v.Field(j).Interface())
				// }
			}
		}
	}

	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i)
		t := v.Type().Field(i)

		// if _, ok := tipe.Tag.Lookup("type"); !ok {
		if t.Type.Kind() == reflect.Struct {
			tagStructWalk(value)
			continue
		}

		// tag := getTags(tipe)
		// key := tag.Group + "." + strings.ToLower(tipe.Name)
		// val := value.Interface()

		// fmt.Println(key)
		// switch tag.Tipe {
		// case "string":
		// 	val = val.([]byte)
		// case "datetime":
		// 	val = formatter.ToUnixTime(val.([]byte))
		// }
		// buf[key] = val
	}
}

func getTags(field reflect.StructField) Tag {
	tag := Tag{}

	// if group, ok := field.Tag.Lookup("group"); ok {
	// 	tag.Group = group
	// }
	if tipe, ok := field.Tag.Lookup("type"); ok {
		tag.Tipe = tipe
	}
	if unit, ok := field.Tag.Lookup("unit"); ok {
		tag.Unit = unit
	}
	if factor, ok := field.Tag.Lookup("factor"); ok {
		if f, err := strconv.ParseFloat(factor, 32); err == nil {
			tag.Scale = float32(f)
		}
	}
	if _, ok := field.Tag.Lookup("chartable"); ok {
		tag.Chartable = true
	}

	return tag
}
