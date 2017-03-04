package main

import (
	"os"
	"testing"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
)

func iHaveATemplateWithTheFollowingContent(arg1 *gherkin.DocString) error {
	return godog.ErrPending
}

func iRunTheTemplateTool() error {
	return godog.ErrPending
}

func iShouldHaveTheFollowingMessageSentToSlack(arg1 *gherkin.DocString) error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I have a template with the following content:$`, iHaveATemplateWithTheFollowingContent)
	s.Step(`^I run the template tool$`, iRunTheTemplateTool)
	s.Step(`^I should have the following message sent to Slack:$`, iShouldHaveTheFollowingMessageSentToSlack)
}

func TestMain(m *testing.M) {
	status := godog.RunWithOptions("godogs", func(s *godog.Suite) {
		FeatureContext(s)
	}, godog.Options{
		Format: "progress",
		Paths:  []string{"features"},
	})

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}
