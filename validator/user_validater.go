package validator

import (
	"go-rest-api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)



type ITaskValidetor interface {
	TaskValidate(task model.Task) error
}

type taskValidate struct{}

type NewTaskValidate() ITaskValidetor {
	return &taskValidate{}
}

func (tv *taskValidator) TaskValidate(task model.Task) error {
	return validation.ValidateStruct(&task,
		validation.Field(
			&task.Title,
			validation.Required.Error("title is required"),
			validation.RuneLength(1, 10).Error("limited max 10 char"),
		),
	)
}