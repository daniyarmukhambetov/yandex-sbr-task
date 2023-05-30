package models

import (
	"context"
	"gorm.io/gorm/schema"
	"reflect"
	"strings"
)

type Hours []string

func (h *Hours) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue interface{}) (err error) {
	*h = strings.SplitN(dbValue.(string), "-", 2)
	return nil
}

func (h Hours) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue interface{}) (interface{}, error) {
	return string(h[0]) + "-" + string(h[1]), nil
}
