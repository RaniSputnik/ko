Feature: Join match

Background:
Given Bob is logged in

Scenario: Bob joins a match
Given Alice has created 1 match
When Bob joins Alice's match
Then the match opponent should be Bob
And the match status should be READY

Scenario: Bob joins his own match
Given Bob has created 1 match
When he joins his own match
Then he should get an error: 'A user can not join a match that they created.'

Scenario: Bob joins a match that does not exist
When Bob joins a match that does not exist
Then he should get an error: 'A match with the given id could not be found.'

Scenario: Bob joins a match that is already full
Given Alice has created a match against Clive
When Bob joins Alice's match
Then he should get an error: 'The match is already full.'