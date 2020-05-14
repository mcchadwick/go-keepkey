package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(changeWipeCodeCmd)
}

var changeWipeCodeCmd = &cobra.Command{
	Use:   "changeWipeCode",
	Short: "Change or add a wipe code to the device",
	Long:  "Change or add a wipe code to the device",
	Run: func(cmd *cobra.Command, args []string) {
		if err := kk.ChangeWipeCode(); err != nil {
			log.Fatal(err)
		}
	},
}
