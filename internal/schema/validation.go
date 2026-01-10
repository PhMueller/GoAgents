package schema

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var IsStringValidUUID validator.Func = func(fl validator.FieldLevel) bool {
	/* This function is applied to fields in structs to verify that the parameter is a valid uuid.

	Reason: gin does not have built-in uuid validation, so we add it here. After validation is registered, we can remove
	the custom validation in the schema definitions.

	in the `ShouldBind` call an exception is raised if the validation fails.
	the `tag` is then the name of the validation

	This function is tested in the schema that use the function.
	*/
	value, ok := fl.Field().Interface().(string)
	if ok {
		_, err := uuid.Parse(value)
		return err == nil
	}
	return false
}
