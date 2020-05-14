package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(removeWipeCodeCmd)
}

var removeWipeCodeCmd = &cobra.Command{
	Use:   "removeWipeCode",
	Short: "Disable wipe code on the device",
	Long: `Disables the wipe code on the device. If there is currently
		a wipe code then it will prompt the user to enter the current pin
		before disabling`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := kk.RemoveWipeCode(); err != nil {
			log.Fatal(err)
		}
	},
}
