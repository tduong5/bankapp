package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/tduong5/simplebank/util"
)

var ValidCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		// cheeck currency is supported
		return util.IsSupporterCurrency(currency)
	}
	return false
}
