package api

import (
	"github.com/TheDP66/simple_bank_go/db/util"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	// if ok is true
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		// check currency is supported
		return util.IsSupportedCurrency(currency)
	}

	return false
}
