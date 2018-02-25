Feature: Create match

Scenario: Alice creates a match
Given Alice is logged in
When she creates a new match
Then Alice should get a new match
And there should be no errors