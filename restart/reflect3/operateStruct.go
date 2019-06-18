package operateStruct

import (
	"fmt"
	"reflect"
)

type user struct {
	name string
	age  int8
}

func reflectUsed(data interface{}) int8 {
	reValue := reflect.ValueOf(data)
	fmt.Printf("reValue: %v", reValue)
	data = reValue.Interface()
	userStruct := data.(user)
	return userStruct.age
}
