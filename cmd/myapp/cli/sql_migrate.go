package cli

import (
	"errors"

	"github.com/i4erkasov/go-pgsql/migrate"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

const (
	MigrationCommand      = "sql-migrate"
	VersionMigrateService = "0.0.1"
	MigrationUpArg        = "up"
	MigrationDownArg      = "down"
)

var steps int // Глобальная переменная для хранения количества шагов

var sqlMigrate = &cobra.Command{
	Use:        MigrationCommand,
	Short:      "Run database migration",
	Version:    VersionMigrateService,
	Args:       cobra.MaximumNArgs(1),
	ArgAliases: []string{MigrationUpArg, MigrationDownArg},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		var migration migrate.SQLMigrator
		if migration, err = migrate.NewWithViper(cfg.Sub("app")); err != nil {
			return err
		}

		switch true {
		case slices.Contains(args, MigrationUpArg):
			return migration.Up(cmd.Context())
		case slices.Contains(args, MigrationDownArg):
			return migration.Down(cmd.Context(), steps)
		default:
			return errors.New("invalid argument. please specify 'up' or 'down' for migration")
		}
	},
}

func init() {
	sqlMigrate.Flags().IntVarP(&steps, "steps", "s", 1, "Number of steps to migrate down")

	cmd.AddCommand(sqlMigrate)
}
