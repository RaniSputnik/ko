Feature: Join match

Scenario: Bob joins a match
Given Bob is logged in
And Alice has created 1 match
When Bob joins Alice's match
Then the match opponent should be Bob
And the match status should be READY

Scenario: Bob joins his own match
Given Bob is logged in
And he has created 1 match
When Bob joins his own match
Then he should get an error: 'A user can not join a match that they created.'

Scenario: Bob joins a match that does not exist
Given Bob is logged in
When Bob joins a match that does not exist
Then he should get an error: 'A match with the given id could not be found.'

Scenario: Bob joins a match that is already full
Given Bob is logged in
And Alice has created a match against Clive
When Bob joins Alice's match
Then he should get an error: 'The match is already full.'