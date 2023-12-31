package cli

import (
	"context"
	logger "log"
	"os"

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

	cmd, err := app.Parse(args[1:])
	if err != nil {
		log.Fatal(errors.Wrap(err, "Failed to parse cli command"))
	}
	cfg := config.New(fig.File(*configName), rateLimit)
	ctx := context.Background()
	switch cmd {
	case migrateUpCmd.FullCommand():
		MigrateUp(cfg)
	case runCmd.FullCommand():
		http.Run(cfg, ctx)
	default:
		log.Fatalf("Unknown cmd :(")
	}
}

func main() {
	Run(os.Args)
}
