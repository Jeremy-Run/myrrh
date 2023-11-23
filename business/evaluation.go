package business

import "reflect"

func Eval(feature string) int {
	user := User{}
	value := reflect.ValueOf(&user)
	f := value.MethodByName(feature)
	results := f.Call([]reflect.Value{})
	return int(results[0].Int())
}
