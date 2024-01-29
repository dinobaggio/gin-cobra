package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gin-cobra",
	Short: "gin-cobra menjalankan sebuah aplikasi server dengan command line",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		// some code
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
