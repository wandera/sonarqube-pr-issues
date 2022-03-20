package server

import (
	"github.com/spf13/cobra"
)

var serverPort int
var workers int
var event string

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Executes the server commands",
}

func init() {
	ServerCmd.PersistentFlags().IntVarP(&serverPort, "port", "p", 8080, "Server port")
	ServerCmd.PersistentFlags().IntVarP(&workers, "workers", "w", 30, "Workers count")
	ServerCmd.PersistentFlags().StringVarP(&event, "event", "e", "REQUEST_CHANGES", "GitHub review event type")
	ServerCmd.AddCommand(RunCmd)
}
