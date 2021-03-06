package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "synthetic",
	Short: "Synthetic simulates operational scenarios for microservices",
	Long: `
........................................................................
:                                      __   __          __   __        :
:     **            .-----.--.--.-----|  |_|  |--.-----|  |_|__.----.  :
:   *    *          |__ --|  |  |     |   _|     |  -__|   _|  |  __|  :
:  * ---- * ---- *  |_____|___  |__|__|____|__|__|_____|____|__|____|  :
:          *    *         |_____|                                v1.0  :
:           **                                                         :
........................................................................
`,
}
