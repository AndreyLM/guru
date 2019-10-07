package errors

import (
	"errors"
	"fmt"
	"log"
	"runtime"
)

var (
	// UserExistError - user not exist error
	UserExistError = GenerateError("Cannot create user. Already exist")
	// UserNotExistError - user not exist error
	UserNotExistError = GenerateError("User does not exist")
	// InternalServerError - internal server error
	InternalServerError = GenerateError("internal server error")
)

// GenerateError - generates error
func GenerateError(err string) error {
	return errors.New(err)
}

// DebugPrintf - debug error
func DebugPrintf(err error, args ...interface{}) string {
	programCounter, file, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(programCounter)
	msg := fmt.Sprintf("[%s: %s %d] %s, %s", file, fn.Name(), line, err, args)
	log.Println(msg)
	return msg
}
