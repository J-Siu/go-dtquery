/*
Copyright © 2025 John, Sing Dao, Siu <john.sd.siu@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package cmd

import (
	"os"

	"github.com/J-Siu/go-dtquery/dq"
	"github.com/J-Siu/go-ezlog"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-dtquery",
	Short: "Query Devtools version and page information",
	Run: func(cmd *cobra.Command, args []string) {
		// Setup log functions
		ezlog.SetAllPrintln()
		// Setup log level
		debug, _ := cmd.Flags().GetBool("debug")
		if debug {
			ezlog.SetLogLevel(ezlog.DebugLevel)
		}

		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetInt("port")
		// dq.Get return a populated Devtools struct
		devtools := dq.Get(host, port)

		// Print out devtools
		ezlog.MsgP(dq.MustToJsonStrP(devtools))
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("host", "r", "localhost", "Hostname or IP")
	rootCmd.PersistentFlags().IntP("port", "p", 9222, "Port")
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "Debug mode")
}
