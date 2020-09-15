package ids

import "github.com/google/uuid"

func Create(parentId uuid.UUID, mask Mask) uuid.UUID {
	if (mask == Mask(-1)) {
		return uuid.New()
	}
	return uuid.MustParse(string(parentId.String()[0:mask+1] + uuid.New().String()[mask+1:36]))
}
