package packet

import (
	"reflect"
	"strconv"

	"github.com/pudjamansyurin/go-gen-tracker/util"
)

type M map[string]interface{}

type Meta struct {
	Tipe      string
	Unit      string
	Scale     float32
	Chartable bool
}

func GetMeta(packet interface{}) {
	// TypeOf returns the reflection Type that represents the dynamic type of variable.
	// If variable is a nil interface value, TypeOf returns nil.
	buffer := M{}
	tagWalk(reflect.TypeOf(packet).Elem(), buffer)
}

func tagWalk(t reflect.Type, buf M) {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if field.Type.Kind() == reflect.Struct {
			tagWalk(field.Type, buf)
			continue
		}

		meta := decodeMeta(field)
		util.Debug(meta)
		// buf = append()
	}
}

func decodeMeta(field reflect.StructField) Meta {
	meta := Meta{
		Scale: 1.0,
	}

	if tipe, ok := field.Tag.Lookup("type"); ok {
		meta.Tipe = tipe
	}
	if unit, ok := field.Tag.Lookup("unit"); ok {
		meta.Unit = unit
	}
	if scale, ok := field.Tag.Lookup("scale"); ok {
		if d, err := strconv.ParseFloat(scale, 32); err == nil {
			meta.Scale = float32(d)
		}
	}
	if _, ok := field.Tag.Lookup("chartable"); ok {
		meta.Chartable = true
	}

	return meta
}
