package cmd

import (
	"fmt"
	"github.com/codersgarage/emqx-influxdb-exporter/log"
	"github.com/codersgarage/emqx-influxdb-exporter/worker"
	"os"

	"github.com/spf13/cobra"
)

var (
	// RootCmd is the root command of emqx-influxdb-exporter service
	RootCmd = &cobra.Command{
		Use:   "emqx-influxdb-exporter",
		Short: "A http service",
		Long:  `An HTTP JSON API backend service`,
	}
)

func init() {
	RootCmd.AddCommand(serveCmd)
}

// Execute executes the root command
func Execute() {
	log.SetupLog()

	go worker.RunStatWorker()
	go worker.RunMetricsWorker()

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
