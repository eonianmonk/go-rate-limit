package backend

import "github.com/eonianmonk/go-rate-limit/backend/cli"

func startApp() {
	// can omit "--max-rate 50", because it is default value
	cli.Run([]string{"placeholder", "run", "--max-rate", "50"})
}
