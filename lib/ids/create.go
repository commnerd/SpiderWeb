package ids

import "github.com/google/uuid"

func CreateID() uuid.UUID {
	return uuid.New();
}
