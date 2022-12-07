package constants

import "fmt"

type LogType int

const (
	InternalServerError = "internal server error"
	Success             = "success"
)

// Enum
const (
	User LogType = 1
)

func (e LogType) String() string {
	switch e {
	case User:
		return "user"
	default:
		return fmt.Sprintf("%d", int(e))
	}
}
