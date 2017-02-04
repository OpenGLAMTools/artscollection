// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/OpenGLAMTools/artscollection/cmd/artscollection/server"
	"github.com/OpenGLAMTools/artscollection/collection"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	Long: `The server enables the gui. Just open the url with the correct
port.`,
	Run: func(cmd *cobra.Command, args []string) {
		server.ServerPort = viper.GetString("serverPort")
		c := viper.GetStringMapString("collections")
		server.Artscollection = loadCollections(c)
		server.Serve()
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func loadCollections(paths map[string]string) map[string]*collection.Collection {
	ac := make(map[string]*collection.Collection)
	for k, cp := range paths {
		var err error
		ac[k], err = collection.LoadTxt(cp)
		if err != nil {
			errorLog(err, "loadCollection error:")
		}
	}
	return ac
}

func errorLog(err error, s string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s:%v", s, err)
	}
}
