Feature: Matches

Background:
Given Alice is logged in

Scenario: Alice creates a match
When Alice creates a new match
Then Alice should get a new match
And there should be no errors

Scenario: Default board size
When Alice creates a new match without specifying board size
Then she should get a new match
And the board should be 19x19
And there should be no errors

Scenario: Alice gets matches
Given Alice has created 2 matches
When Alice requests her matches
Then she should get her 2 matches
And there should be no errors

Scenario: Alice does not see Bob's matches
Given Alice has created 1 match
And Bob has created 3 matches
When Alice requests her matches
Then she should get her 1 match