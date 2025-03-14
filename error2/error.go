package error2

import "fmt"

func ErrorWrap(msg string, err error) error {
	return fmt.Errorf(msg, err)
}
