package status

import "github.com/samber/lo"

type Status string

func (s Status) String() string {
	return string(s)
}

func (s Status) IsFinal() bool {
	return lo.Contains([]Status{
		Status_SUCCESS,
		Status_FAILED,
		Status_UNKNOWN,
	}, s)
}

var (
	Status_CREATED Status = "CREATED"
	Status_PENDING Status = "PENDING"
	Status_SUCCESS Status = "SUCCESS"
	Status_FAILED  Status = "FAILED"
	Status_UNKNOWN Status = "UNKNOWN"
	Status_EXPIRED Status = "EXPIRED"
)
