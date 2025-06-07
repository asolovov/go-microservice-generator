package gen

import (
	"fmt"
	"github.com/misnaged/annales/logger"
	"github.com/spf13/cobra"
	"gmg/internal"
)

// Cmd returns the "serve" command of the application.
// This command is responsible for initializing and
func Cmd(app *internal.App) *cobra.Command {
	var cfg string

	cmd := &cobra.Command{
		Use:   "gen",
		Short: "Generate Microservice",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := app.Init(); err != nil {
				return fmt.Errorf("application initialisation: %w", err)
			}

			return app.Generate(cfg)
		},
		PreRun: func(cmd *cobra.Command, args []string) {
			logger.Log().Info(app.Version())
		},
	}

	cmd.Flags().StringVarP(&cfg, "config", "c", "gmg-cfg.yaml", "config file")

	return cmd
}
