package util

import (
	"os"
	"path/filepath"

	"gopkg.in/logger.v1"
)

func CreateFlowPath(catelog, userId string) string {
	flowPath := filepath.Join("/data/xtrader", catelog, "flowpath", userId)
	if _, err := os.Stat(flowPath); os.IsNotExist(err) {
		log.Println("create flowpath: ", flowPath)
		os.MkdirAll(flowPath, 0775)
	}
	return flowPath + string(filepath.Separator)
}
