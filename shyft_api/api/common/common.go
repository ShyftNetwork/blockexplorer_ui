package common

import (
	"strconv"
	"github.com/ShyftNetwork/blockexplorer_ui/shyft_api/logger"
)

// StringToInteger returns int64 from string params
func StringToInteger(str string) int64 {
	response, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		logger.Warn("Error converting params: " + err.Error())
	}
	return response
}
