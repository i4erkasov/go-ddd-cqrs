package cli

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	config string
	cfg    *viper.Viper
)

var cmd = &cobra.Command{
	Use:   "cmd",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		cfg = viper.New()
		cfg.AddConfigPath(".")
		cfg.AutomaticEnv()
		cfg.SetConfigFile(config)
		return cfg.ReadInConfig()
	},
}

func init() {
	cmd.PersistentFlags().StringVarP(
		&config, "config", "c", "config.yml",
		"path to config file",
	)
}
