package instance

import "time"

type DeletionQueueElement struct {
	Id string
	Ts time.Time
}
