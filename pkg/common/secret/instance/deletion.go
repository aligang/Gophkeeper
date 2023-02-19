package instance

import "time"

type DeletionQueueElement struct {
	ObjectId  string
	DeletedAt time.Time
}
