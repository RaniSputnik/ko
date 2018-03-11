Feature: Join match

Scenario: Bob joins a match
Given Bob is logged in
And Alice has created 1 match
When Bob joins Alice's match
Then the match opponent should be Bob
And the match status should be READY