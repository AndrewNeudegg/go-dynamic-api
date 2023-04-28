package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/andrewneudegg/go-dynamic-api/cmd/toady/subcmd/serve"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func configureLogger(verbose bool) {

	packageName := "/toady/"

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
	// Only log the warning severity or above.
	if verbose {
		log.SetReportCaller(true)
		log.StandardLogger().SetFormatter(&logrus.TextFormatter{
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				// s := strings.Split(f.Function, ".")
				// funcName := s[len(s)-1]
				relativeFPath := strings.SplitAfterN(f.File, packageName, 2)[1]
				return "", fmt.Sprintf(" %s:%d", relativeFPath, f.Line)
				// return funcName, fmt.Sprintf(" %s:%d", relativeFPath, f.Line)
			},

			DisableColors: false,
			FullTimestamp: true,
		})
		log.SetLevel(log.DebugLevel)
	} else {
		log.StandardLogger().SetFormatter(&logrus.TextFormatter{
			DisableColors: false,
			FullTimestamp: true,
		})
		log.SetLevel(log.InfoLevel)
	}
}

func main() {
	var verboseMode bool

	var rootCmd = &cobra.Command{
		Use: "toady",
		Long: `Toady: An example of a dynamic API application."
		`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			configureLogger(verboseMode)
			return nil
		},
	}

	rootCmd.PersistentFlags().BoolVarP(&verboseMode, "verbose", "v", false, "verbose output")
	rootCmd.AddCommand(serve.Cmd())

	err := rootCmd.Execute()
	if err != nil {
		log.Error(err)
	}
}
