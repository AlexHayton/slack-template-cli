Feature: Send a message to Slack based on a basic template
    As a user, I can use a simple text template to send messages to a Slack channel

Scenario: Simple text template
    Given I have a template with the following content:
        """
        {
            "channel": "hello",
            "username": "why hello", 
            "icon_emoji": ":ghost:"
        }
        """
    When I run the template tool
    Then I should have the following message sent to Slack:
        """
        {
            "channel": "hello",
            "username": "why hello", 
            "icon_emoji": ":ghost:"
        }
        """