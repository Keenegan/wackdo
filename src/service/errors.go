package service

import (
	"fmt"
	"reflect"
	"strings"
)

type InvalidParamError struct {
	Reason string
}

func (e *InvalidParamError) Error() string {
	return fmt.Sprintf("invalid parameter : %s'", e.Reason)
}

type DuplicateEmailError struct{}

func (e *DuplicateEmailError) Error() string {
	return "email already in use"
}

type EntityNotFoundError struct {
	Entity interface{}
}

func (e *EntityNotFoundError) Error() string {
	entityType := reflect.TypeOf(e.Entity)
	if entityType == nil {
		return "entity not found"
	}

	// Remove pointer and package prefix to get clean entity name
	typeName := entityType.String()
	typeName = strings.TrimPrefix(typeName, "*")
	if idx := strings.LastIndex(typeName, "."); idx != -1 {
		typeName = typeName[idx+1:]
	}

	return fmt.Sprintf("%s not found", strings.ToLower(typeName))
}
