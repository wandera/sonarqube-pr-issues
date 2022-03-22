package server

import (
	"github.com/herlon214/sonarqube-pr-issues/pkg/scm"
	"github.com/spf13/cobra"
)

var serverPort int
var workers int
var reviewEvent string

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Executes the server commands",
}

func init() {
	ServerCmd.PersistentFlags().IntVarP(&serverPort, "port", "p", 8080, "Server port")
	ServerCmd.PersistentFlags().IntVarP(&workers, "workers", "w", 30, "Workers count")
	ServerCmd.PersistentFlags().StringVarP(&reviewEvent, "review-event", "e", scm.REVIEW_EVENT_REQUEST_CHANGES, "GitHub review event type")
	ServerCmd.AddCommand(RunCmd)
}
