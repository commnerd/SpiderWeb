package id

import (
	"github.com/google/uuid"
	"../util"
)

func Derive(parentId Id, mask Mask) Id {
	mask++

	exists, _ := util.InArray(mask, BadMasks)

	if (exists) {
		mask++
	}

	return Id(uuid.MustParse(string(parentId.String()[0:mask+1] + uuid.UUID(New()).String()[mask+1:36])))
}
