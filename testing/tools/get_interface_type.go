package tools

import "reflect"

// GetInterfaceType : Get the interface type using "reflect"
func GetInterfaceType(i interface{}) string {
    return reflect.TypeOf(i).Name()
}
