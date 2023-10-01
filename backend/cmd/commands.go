package cmd

import "github.com/alecthomas/kingpin"

var (
	app = kingpin.New("ratel", "api with limited rate")

	runCmd = app.Command("run", "run svc")

	migrateCmd   = app.Command("migrate", "migrate cmd")
	migrateUpCmd = migrateCmd.Command("up", "migrate up cmd")
	// flags
	configName = app.Flag("cfg", "config file name").Default("config.yaml").String()
	rateLimit  = app.Flag("max-rate", "rate limit").Default("50").Int16()
)
