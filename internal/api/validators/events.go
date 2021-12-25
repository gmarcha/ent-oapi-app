package validators

import (
	"time"

	"github.com/go-playground/validator/v10"
)

var validateDateStart validator.Func = func(fl validator.FieldLevel) bool {

	dateStart, ok := fl.Field().Interface().(time.Time)
	if ok {
		now := time.Now()
		if now.After(dateStart) {
			return false
		}
	}
	return true
}
