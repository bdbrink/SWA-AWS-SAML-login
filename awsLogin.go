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

// need name, role, e/xid

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

		if strings.HasPrefix(c.String("name"), "lab") {
			accountID = "217906394988"
			if c.String("name") == "lab" {
				role = "swa/SWACSDeveloper"
			}

			if c.String("name") == "labOp" {
				role = "swa/SWACSOperations"
			}
		}

		if strings.HasPrefix(c.String("name"), "dev") {
			accountID = "988101568216"

			if c.String("name") == "devDev" {
				role = "swa/SWACSDeveloper"
			}

			if c.String("name") == "dev" {
				role = "swa/SWACSOperations"
			}
		}

		if strings.HasPrefix(c.String("name"), "qa") {
			accountID = "042808334126"

			if c.String("name") == "qa" {
				role = "swa/SWACSDeveloper"
			}

			if c.String("name") == "qaOp" {
				role = "swa/SWACSOperations"
			}
		}

		if strings.HasPrefix(c.String("name"), "prod") {
			accountID = "707239158216"

			if c.String("name") == "prod" {
				role = "swa/SWACSDeveloper"
			}

			if c.String("name") == "ProdOps" {
				role = "swa/SWACSOperations"
			}
		}

		awsSamlLogin := fmt.Sprintf("awssaml get-credentials --account-id %v --name %v --role %v --user-name e143608 --duration 14400", accountID, c.String("name"), role)

		cmd := exec.Command("bash", "-c", awsSamlLogin)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		_ = cmd.Run()
	}

	fmt.Println("Logged into AWS")
	return nil
}
