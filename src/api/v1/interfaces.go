package api

type node interface{
	RegisterChild(child interface{}) (interface{}, error)
}