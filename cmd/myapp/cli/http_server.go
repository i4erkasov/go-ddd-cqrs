package cli

import (
	"github.com/i4erkasov/go-ddd-cqrs/internal/infrastructure/api/http"
	"github.com/spf13/cobra"
)

const HttpServerCommand = "http-server"
const VersionHttpServer = "1.0.0"

var httpServer = &cobra.Command{
	Use:     HttpServerCommand,
	Short:   "Start http server",
	Version: VersionHttpServer,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		var server *http.Server
		if server, err = http.NewHttpServer(cfg); err != nil {
			return err
		}

		return server.Start(cmd.Context())
	},
}

func init() {
	cmd.AddCommand(httpServer)
}
