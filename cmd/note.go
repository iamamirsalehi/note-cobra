package cmd

import (
	"github.com/spf13/cobra"
)

// noteCmd represents the note command
var noteCmd = &cobra.Command{
	Use:   "note",
	Short: "You can create your notes",
	Long: `You can create your notes`,
}

func init() {
	rootCmd.AddCommand(noteCmd)
}
