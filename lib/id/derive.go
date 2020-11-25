package id

import (
	"github.com/google/uuid"
	"errors"
	"../util"
)

var hexMap []byte = []byte("0123456789abcdef")

func Derive(parentId Id, mask Mask) (Id, error) {
	if len(hexMap) == 0 {
		return parentId, errors.New("No more chars available.")
	}

	mask++

	exists, _ := util.InArray(mask, BadMasks)

	if (exists) {
		mask++
	}

	parentIdString := []byte(parentId.String())

	parentIdString[mask] = hexMap[0]

	hexMap = hexMap[1:]

	return Id(uuid.MustParse(string(parentIdString))), nil
}
