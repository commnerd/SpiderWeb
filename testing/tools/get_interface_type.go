package tools

import "reflect"

func GetInterfaceType(i interface{}) string {
    return reflect.TypeOf(i).Name()
}
