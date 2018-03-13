Feature: Play a stone

Background:
Given Alice is logged in

Scenario: Alice plays a stone
Given Alice is playing against Bob
When she places a stone at A3
Then there should be 1 move played
And the match status should be IN_PROGRESS
And there should be no errors

Scenario: A stone is played when there is no opponent
Given she has created 1 match
When she places a stone at A3
Then she should get an error: 'The match does not have an opponent yet.'

Scenario: A stone is played in a match Alice is not a part of
Given Bob is playing against Clive
When she places a stone at A3
Then she should get an error: 'You are not participating in the match.'

Scenario: A stone is played in an occupied position
Given Alice is playing against Bob
And the folling moves have been played: A3,G5,D4
When she places a stone at A3
Then she should get an error: 'Can not play in an occupied position.'

#Feature: Taking turns

# Scenario: Black plays first
# Given Bob is logged in
# And he is playing against Alice
# And Bob is the black player
# When Bob requests the next player
# Then he should get himself

#Feature: Liberties

# A stone is played surrounding an opponents stone

# A group of stones loses final liberty

# A stone is played in a surrounded position
# Then she should get an error: 'Can not play in a position with no liberties.'

#Feature: Ko