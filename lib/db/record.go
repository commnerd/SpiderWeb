package db

type Record struct{
	Collection string
	Id interface{}
	Nodes []node
	Body interface{}
}