package schema

const Text = `

schema {
	query: Query
	mutation: Mutation
}

type Query {
	# The matches that the current player is part of.
	matches(first:Int, after:ID, last:Int, before:ID): MatchConnection!
	# The lobby contains matches currently available for joining.
	lobby: Lobby!
}

type Mutation {
	createMatch(boardSize:Int = 19): Match!
	joinMatch(matchId: ID!): Match!

	playStone(matchId: ID!, x: Int!, y:Int!): Match!
	skip(matchId: ID!): Match!
	resign(matchId: ID!): Match!
}

type Player {
	id: ID!
	username: String!
	colour: Colour!
}

type Lobby {
	playersOnlineCount: Int!
	matches(first:Int, after:ID, last:Int, before:ID): MatchConnection!
}

type MatchConnection {
	nodes: [Match!]!
	totalCount: Int!
}

type Match {
	id: ID!

	# The player who created this match.
	createdBy: Player!
	# Gets a player by stone colour.
	player(colour: Colour!): Player!
	# The player who will have the next move.
	next: Player!

	board: Board!
	events: EventConnection!
}

type Board {
	size: Int!
	stones: [Stone!]!
}

type Stone {
	colour: Colour!
	x: Int!
	y: Int!
}

enum Colour {
	BLACK
	WHITE
}

type EventConnection {
	nodes: [Event!]!
	totalCount: Int!
}

interface Event {
	player: Player!
	localisedDescription: String!
}

type PlaceStone implements Event {
	player: Player!
	localisedDescription: String!
	x: Int!
	y: Int!
}

type Skip implements Event {
	player: Player!
	localisedDescription: String!
}

type Resign implements Event {
	player: Player!
	localisedDescription: String!
}
`
