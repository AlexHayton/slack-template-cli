package main

import (
	"bytes"
	"log"
	"os"
	"text/template"

	"gopkg.in/urfave/cli.v1"
)

// GenerateFromTemplate : Generates a JSON string from the template provided
func GenerateFromTemplate(templateFile string, data EnvironmentData) (string, error) {
	funcMap := template.FuncMap{
		"env": os.Getenv,
	}

	parsedTemplate := template.Must(template.New("template").Parse(templateString))
	parsedTemplate.Funcs(funcMap)
	var buffer bytes.Buffer
	err := parsedTemplate.Execute(&buffer, data)

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return buffer.String(), nil
}

func main() {
	var slackToken string
	var channel string
	var templateFile string

	app := cli.NewApp()
	app.Name = "Slack Template CLI"
	app.Usage = ""

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "slackToken",
			Value:       "MY_SLACK_TOKEN",
			Usage:       "Your slack API token",
			Destination: &slackToken,
		},

		cli.StringFlag{
			Name:        "channel",
			Value:       "#debug",
			Usage:       "The channel you want to post to",
			Destination: &channel,
		},

		cli.StringFlag{
			Name:        "templateFile",
			Value:       "template.txt",
			Usage:       "The path to the template you want to fill in",
			Destination: &template,
		},
	}

	app.Action = func(c *cli.Context) error {
		var data EnvironmentData
		template, err := GenerateFromTemplate(templateFile, data)
		if err != nil {
			log.Fatal("Could not generate template!", err)
		}

		SendMessageToSlack(template)
	}

	app.Run(os.Args)
}
