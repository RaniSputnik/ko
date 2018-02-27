Feature: Create match

Scenario: Alice creates a match
Given Alice is logged in
When she creates a new match
Then Alice should get a new match
And there should be no errors

Scenario: Default board size
Given Alice is logged in
When she creates a new match without specifying board size
Then Alice should get a new match
And the board should be 19x19
And there should be no errors