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
	Short: "ShortDescription ",
	Long:  `Long Description`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		cfg = viper.New()
		cfg.AddConfigPath(".")
		cfg.AutomaticEnv()
		cfg.SetConfigFile(config)
		return cfg.ReadInConfig()
	},
}

func Execute() error {
	return cmd.Execute()
}

func init() {
	cmd.PersistentFlags().StringVarP(
		&config, "config", "c", "config.yml",
		"path to config file",
	)
}
