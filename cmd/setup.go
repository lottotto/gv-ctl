package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/xanzy/go-gitlab"
)

type gvclt gitlab.Client

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "variables set in the GitLab project",
	Run: func(cmd *cobra.Command, args []string) {
		filePath, err := cmd.PersistentFlags().GetString("filePath")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		manifest, err := readManifest(filePath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		client, err := connectGitLab(manifest.BaseURL, manifest.Token)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		ID := manifest.ID

		if manifest.Type == "project" {
			if err = deleteProjectVariables(client, ID); err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
	
			for _, v := range manifest.Variables {
				err := createProjectVariables(client, v, ID)
				if err != nil {
					log.Fatal(err)
					os.Exit(1)
				}
	
				logMessage := fmt.Sprintf("[Info] Set Succeed. ProjectID: %s, Key: %s, Value: %s", ID, v.Key, v.Value)
				log.Println(logMessage)
			}
		} else if manifest.Type == "group" {
			if err = deleteGroupVariables(client, ID); err != nil {
				log.Fatal(err)
				os.Exit(1)
			}

			for _, v := range manifest.Variables {
				err := createGroupVariables(client, v, ID)
				if err != nil {
					log.Fatal(err)
					os.Exit(1)
				}

				logMessage := fmt.Sprintf("[Info] Set Succeed. GroupID: %s, Key: %s, Value: %s", ID, v.Key, v.Value)
				log.Println(logMessage)
			}
		} else {
			logMessage := fmt.Sprintf("[Fatal] Enter the appropriate value in %s (type)", filePath)
			log.Println(logMessage)
		}
	},
}

func init() {
	setupCmd.PersistentFlags().StringP("filePath", "f", "gvctl.yaml", "variables config file path")
	rootCmd.AddCommand(setupCmd)
}




func createProjectVariables(client *gitlab.Client, v Variable, ID string) (err error) {
	var vtype gitlab.VariableTypeValue = "env_var"
	variable := &gitlab.CreateProjectVariableOptions{
		Key:          &v.Key,
		Value:        &v.Value,
		VariableType: &vtype,
		Protected:    &v.Protected,
		Masked:       &v.Masked,
	}
	_, _, err = client.ProjectVariables.CreateVariable(ID, variable, nil)
	if err != nil {
		return err
	}
	return nil
}

func deleteProjectVariables(client *gitlab.Client, ID string) (err error) {
	projectVariablesOptions := &gitlab.ListProjectVariablesOptions{PerPage: 100}
	currentVariables, _, err := client.ProjectVariables.ListVariables(ID, projectVariablesOptions, nil)
	if err != nil {
		return err
	}
	for _, v := range currentVariables {
		_, err = client.ProjectVariables.RemoveVariable(ID, v.Key, nil)
		if err != nil {
			return err
		}
	}
	return nil
}


func createGroupVariables(client *gitlab.Client, v Variable, ID string) (err error) {
	var vtype gitlab.VariableTypeValue = "env_var"
	variable := &gitlab.CreateGroupVariableOptions{
		Key:          &v.Key,
		Value:        &v.Value,
		VariableType: &vtype,
		Protected:    &v.Protected,
		Masked:       &v.Masked,
	}
	_, _, err = client.GroupVariables.CreateVariable(ID, variable, nil)
	if err != nil{
		return err
	}
	return err
}

func deleteGroupVariables(client *gitlab.Client, ID string) (err error) {
	groupVariablesOptions := &gitlab.ListGroupVariablesOptions{}
	groupVariablesOptions.PerPage = 100
	currentVariables, _, err := client.GroupVariables.ListVariables(ID, groupVariablesOptions, nil)
	if err != nil {
		return err
	}
	for _, v:= range currentVariables {
		_, err = client.GroupVariables.RemoveVariable(ID, v.Key, nil)
		if err != nil{
			return err
		}
		logMessage := fmt.Sprintf("[Info] Remove Succeed. GroupID: %s, Key: %s", ID, v.Key)
		log.Println(logMessage)

	}
	return nil
}