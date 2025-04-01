package constants

import "errors"

var Error = struct {
	SomethingWentWrong error
}{
	SomethingWentWrong: errors.New("Something went wrong"),
}
