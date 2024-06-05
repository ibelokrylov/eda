package config

var statusMap = map[string]int{
	"OK":           0,
	"FAIL":         1,
	"SERVER_ERROR": 2,
	"NOT_FOUND":    3,
	"NOT_AUTH":     4,
}

func GetStatus(status string) int {
	if val, ok := statusMap[status]; ok {
		return val
	}
	return statusMap["NOT_FOUND"]
}
