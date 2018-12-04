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
	"fmt"
	"github.com/pierreboissinot/go-wrike"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xanzy/go-gitlab"
	"gopkg.in/src-d/go-git.v4"
	"log"
	"regexp"
	"strings"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import issue from Wrike task",
	Long: `Open gitlab issue from Wrike task, example:
	kif import $WRIKE_PERMALING

	The Gitlab project is defined by the remote origin
	of your current git repository.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		permalink := args[0]
		wrikeApiToken := viper.GetString("wrikeApiToken")
		wrikeClient := wrike.NewClient(nil, wrikeApiToken)
		fields := "[\"description\"]"
		tasks, _, err := wrikeClient.Tasks.QueryTasks(wrike.QueryTasksOptions{permalink, fields})
		if err != nil {
			fmt.Errorf("No results for permalink %v", permalink)
		}

		if len(tasks.Data) != 1 {
			fmt.Errorf("More than 1 task returned: %v", len(tasks.Data))
		}
		task := tasks.Data[0]
		path := "./"
		r, err := git.PlainOpen(path)
		if err != nil {
			fmt.Errorf("Can't open repository from path: %v", path)
		}
		remoteOrigin, err := r.Remote("origin")
		if err != nil {
			fmt.Errorf("No remote with name: %v", "origin")
		}
		if len(remoteOrigin.Config().URLs) < 1 {
			log.Fatal("No url associated with remote origin")
		}
		remoteUrl := remoteOrigin.Config().URLs[0]
		re := regexp.MustCompile(`(?:git|ssh|https?|git@[-\w.]+):(?P<Pid>(\/\/)?(.*?))(\.git)(\/?|\#[-\d\w._]+?)`)
		/*
			fmt.Printf("%#v\n", re.FindStringSubmatch(out.String()))
			fmt.Printf("%#v\n", re.SubexpNames())
			fmt.Println(re.SubexpNames()[1])
		*/
		format := fmt.Sprintf("${%s}", re.SubexpNames()[1])
		pid := strings.TrimSuffix(re.ReplaceAllString(remoteUrl, format), "\n")
		gitlabClient := gitlab.NewClient(nil, viper.GetString("gitlabApiToken"))
		i := &gitlab.CreateIssueOptions{Title: gitlab.String(task.Title), Description: gitlab.String(task.Description)}
		issue, _, err := gitlabClient.Issues.CreateIssue(pid, i)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Created issue: %v\nWith description: %s\nView issue: %s\n", issue.ID, issue.Description, issue.WebURL)
	},
}

func init() {
	rootCmd.AddCommand(importCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// importCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// importCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
