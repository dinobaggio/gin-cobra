package cmd

import (
	"fmt"
	"gin-cobra/bin"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "command untuk menjalankan aplikasi server gin-cobra",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		app := bin.NewApp()
		port := "3000"
		app.Run(fmt.Sprint(":", port))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
