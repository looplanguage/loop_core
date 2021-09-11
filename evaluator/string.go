package evaluator

import (
	"git.kanersps.pw/loop/models/object"
)

func evalString(string string) object.Object {
	return &object.String{Value: string}
}
