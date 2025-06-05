package root

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"gmg/config"
	"gmg/internal"
)

// Cmd returns the root command for the application
func Cmd(app *internal.App) *cobra.Command {
	cmd := &cobra.Command{
		Use:              "gmg",
		Short:            "Go Microservice Generator",
		TraverseChildren: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if isVersion, _ := cmd.Flags().GetBool("version"); isVersion {
				fmt.Println(app.Version())
				os.Exit(0)
			}

			return initializeConfig(cmd, app.Config())
		},
		Version: app.Version(),
	}

	cmd.SetVersionTemplate(app.Version() + "\n")
	cmd.PersistentFlags().BoolP("version", "v", false, "Show application version")

	return cmd
}

// initializeConfig reads in config file and sets configuration
// via environment variables
func initializeConfig(cmd *cobra.Command, cfg *config.Scheme) error {
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			// return fmt.Errorf("read config file: %w", err)
			return err
		}
	}

	// set config via env vars
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	viper.AllowEmptyEnv(true)

	bindFlags(cmd)
	return viper.Unmarshal(cfg)
}

// bindFlags binds flags to the command
func bindFlags(cmd *cobra.Command) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		if !f.Changed && viper.IsSet(f.Name) {
			val := viper.Get(f.Name)
			_ = cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})
}
