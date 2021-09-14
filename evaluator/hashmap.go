package evaluator

import (
	"fmt"
	"git.kanersps.pw/loop/models/ast"
	"git.kanersps.pw/loop/models/object"
)

func evalHashmap(hashmap *ast.Hashmap, env *object.Environment) object.Object {
	pairs := make(map[object.HashKey]object.HashPair)

	for key, value := range hashmap.Values {
		keyVal := Eval(key, env)

		if _, ok := keyVal.(*object.Error); ok {
			return keyVal
		}

		hashedKey, ok := keyVal.(object.Hash)

		if !ok {
			return &object.Error{Message: fmt.Sprintf("key is not valid (not comparable). got=%q", keyVal.Type())}
		}

		valueVal := Eval(value, env)

		if _, ok := valueVal.(*object.Error); ok {
			return valueVal
		}

		hashedKeyValue := hashedKey.Hash()
		pairs[hashedKeyValue] = object.HashPair{
			Key:   keyVal,
			Value: valueVal,
		}
	}

	return &object.HashMap{Pairs: pairs}
}
