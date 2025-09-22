package logs

/*
 * Constants and categories for the API server logging system.
 * This file contains log category constants and log directory paths.
 */

import (
	"path/filepath"

	types "github.com/E-Reboul/go-server/types/logs"

	"go.uber.org/zap"
)

const (
	ServerCategory      types.LogCategory = "server"
	HealthcheckCategory types.LogCategory = "healthcheck"
)

func getLogsCategoriesPaths() map[types.LogCategory]string {
	return map[types.LogCategory]string{
		HealthcheckCategory: filepath.Join("logs", string(HealthcheckCategory), string(HealthcheckCategory)+".log"),
		ServerCategory:      filepath.Join("logs", string(ServerCategory), string(ServerCategory)+".log"),
	}
}

var Loggers = map[types.LogCategory]*zap.Logger{}
