package logs

/*
 * Logger initialization for the API server.
 * This file contains functions to initialize and close loggers.
 */

func InitializeLoggers() error {
	logsPaths := getLogsCategoriesPaths()
	if err := CreateLogDirectory("logs"); err != nil {
		return err
	}
	if err := CreateLoggersDirectories(logsPaths); err != nil {
		return err
	}
	CreateAllLoggersFiles(logsPaths)
	return nil
}

// Ferme tous les loggers (flush)
func CloseLoggers() {
	for _, logger := range Loggers {
		_ = logger.Sync()
	}
}
