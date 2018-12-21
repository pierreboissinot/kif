// Copyright © 2018 Pierre Boissinot <perso@pierreboissinot.me>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"bytes"
	"github.com/BurntSushi/toml"
	"github.com/chzyer/readline"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"io/ioutil"
)

type KifConfig struct {
	WrikeApiToken  string
	GitlabApiToken string
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create configuration",
	Long: `Configure kif with the following parameters:
	- wrikeApiToken
	- gitlabApiToken`,
	Run: func(cmd *cobra.Command, args []string) {
		home, err := homedir.Dir()
		var conf KifConfig
		if _, tomlErr := toml.DecodeFile(home+"/.kif.toml", &conf); tomlErr != nil {
			// handle error
			panic(tomlErr)
		}
		rl, err := readline.NewEx(&readline.Config{
			UniqueEditLine: true,
		})
		if err != nil {
			panic(err)
		}
		defer rl.Close()

		rl.SetPrompt("Enter Wrike API token: ")
		wrikeApiToken, err := rl.Readline()
		conf.WrikeApiToken = wrikeApiToken
		if err != nil {
			return
		}
		rl.ResetHistory()

		rl.SetPrompt("Enter Gitlab API token:")
		gitlabApiToken, err := rl.Readline()
		conf.GitlabApiToken = gitlabApiToken
		buf := new(bytes.Buffer)
		encodeErr := toml.NewEncoder(buf).Encode(conf)
		if nil != encodeErr {
			panic(encodeErr)
		}
		writeErr := ioutil.WriteFile(home+"/.kif.toml", buf.Bytes(), 0664)
		if nil != writeErr {
			panic(writeErr)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
