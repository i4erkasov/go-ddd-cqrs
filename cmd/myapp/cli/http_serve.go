package cli

import (
	"os"
	"time"

	"github.com/i4erkasov/go-ddd-cqrs/internal/infrastructure/api/rest"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const HttpServerCommand = "http-server"
const VersionHttpServer = "1.0.0"

var httpServer = &cobra.Command{
	Use:     HttpServerCommand,
	Short:   "Start http server",
	Version: VersionHttpServer,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		bws := &zapcore.BufferedWriteSyncer{
			WS:            os.Stderr,
			Size:          512 * 1024,
			FlushInterval: time.Minute,
		}
		defer bws.Stop()
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core := zapcore.NewCore(consoleEncoder, bws, zapcore.DebugLevel)
		log := zap.New(core)

		cnf := cfg.Sub("app.api.rest")

		var server *rest.Server
		if server, err = rest.New(cnf, log); err != nil {
			return err
		}

		return server.Start(cmd.Context())
	},
}

func init() {
	cmd.AddCommand(httpServer)
}
