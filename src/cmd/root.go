package cmd

import (
	"github.com/cyberkarma/chatserver/cmd/runserver"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "app",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("execute root cmd")
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(runserver.RunServer)

}
