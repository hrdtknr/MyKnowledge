package repo

import(
	"time"
)

type SampleInterface interface{
	Method() time.Time
}

