package cmd

import (
	logger "log"

	"github.com/alecthomas/kingpin"
	"github.com/eonianmonk/go-rate-limit/backend/config"
	"github.com/eonianmonk/go-rate-limit/backend/http"
	"github.com/kkyr/fig"
	"github.com/pkg/errors"
)

func Run(args []string) {
	log := logger.Default()
	defer func() {
		if rvr := recover(); rvr != nil {
			logger.Fatal(rvr, "-> app panicked")
		}
	}()

	app := kingpin.New("ratel", "api with limited rate")

	runCmd := app.Command("run", "run svc")

	migrateCmd := app.Command("migrate", "migrate cmd")
	migrateUpCmd := migrateCmd.Command("up", "migrate up cmd")
	// flags
	configName := app.Flag("cfg", "config file name").Default("config.yaml").String()
	rateLimit := app.Flag("max-rate", "rate limit").Default("50").Int16()

	cmd, err := app.Parse(args[:1])
	if err != nil {
		log.Fatal(errors.Wrap(err, "Failed to parse cli command"))
	}
	cfg := config.New(fig.File(*configName), rateLimit)

	switch cmd {
	case migrateUpCmd.FullCommand():
		MigrateUp(cfg)
	case runCmd.FullCommand():
		http.Run(cfg)
	default:
		log.Fatalf("Unknown cmd :(")
	}
}
