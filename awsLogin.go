package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "Simple CLI commands for awssaml login/setup"
	app.Usage = "Lets you login into aws via cli if you have proper credentials"

	app.Commands = []cli.Command{
		{
			Name:        "login",
			HelpName:    "login",
			Action:      AwsLogin,
			ArgsUsage:   ` `,
			Usage:       `logins into aws accounts.`,
			Description: `Logging into AWS account via awssaml`,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "name",
					Usage: "login with name setup when awssaml was configured, run awssaml list to see your configuration",
				},
			},
		},
		{
			Name:     "list",
			HelpName: "list",
			Action:   listRoles,
			Usage:    `list all the roles you can log into for AwsLogin command.`,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func AwsLogin(c *cli.Context) error {
	if len(c.Args()) > 0 {
		return errors.New("no arguments are expected, use flags")
	}

	if c.String("name") == "" {
		return errors.New("name of the role you wish to sign is must be provided")
	}

	var accountID string
	var role string

	if c.IsSet("name") {
		fmt.Printf("logging in with %v", c.String("name"))

		if strings.HasPrefix(c.String("name"), "lab") || strings.HasPrefix(c.String("name"), "Lab") {
			accountID = "433838799813"
			if c.String("name") == "lab" {
				role = "swa/SWACSDeveloper"
			}

			if c.String("name") == "labOps" {
				role = "swa/SWACSOperations"
			}

			if c.String("name") == "labBg" {
				role = "swa/SWACSBreakGlassAdmin"
			}
		}

		if strings.HasPrefix(c.String("name"), "dev") || strings.HasPrefix(c.String("name"), "Dev") {
			accountID = "581854693576"

			if c.String("name") == "dev" {
				role = "swa/SWACSDeveloper"
			}

			if c.String("name") == "devOps" {
				role = "swa/SWACSOperations"
			}

			if c.String("name") == "devBg" {
				role = "swa/SWACSBreakGlassAdmin"
			}
		}

		if strings.HasPrefix(c.String("name"), "qa") || strings.HasPrefix(c.String("name"), "Qa") {
			accountID = "550460604145"

			if c.String("name") == "qa" {
				role = "swa/SWACSDeveloper"
			}

			if c.String("name") == "qaOp" {
				role = "swa/SWACSOperations"
			}

			if c.String("name") == "qaBg" {
				role = "swa/SWACSBreakGlassAdmin"
			}
		}

		if strings.HasPrefix(c.String("name"), "prod") || strings.HasPrefix(c.String("name"), "Prod") {
			accountID = "436431778785"

			if c.String("name") == "prod" {
				role = "swa/SWACSDeveloper"
			}

			if c.String("name") == "prodOps" {
				role = "swa/SWACSOperations"
			}

			if c.String("name") == "prodBg" {
				role = "swa/SWACSBreakGlassAdmin"
			}
		}

		// format command
		awsSamlLogin := fmt.Sprintf("awssaml get-credentials --account-id %v --name %v --role %v --user-name e143608 --duration 14400", accountID, c.String("name"), role)

		// execute command, inputs will be prompted, outputs will return as well as any errors
		cmd := exec.Command("bash", "-c", awsSamlLogin)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		_ = cmd.Run()
	}

	return nil
}

func listRoles(c *cli.Context) error {

	cmd := exec.Command("bash", "-c", "awssaml list")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()

	return nil
}
