Given("{user} is logged in") do |user|
    login_user(user.name)
end

Given("{user} has created {int} match(es)") do |user, n_matches|
    for i in 1..n_matches do
        $mysql_client.query("INSERT INTO Matches (Owner, BoardSize) VALUES (#{user.id},#{19})")
    end
end

Given("{user} has created a match against {user}") do |creator, opponent|
    $mysql_client.query("INSERT INTO Matches (Owner, Opponent, BoardSize) VALUES (#{creator.id},#{opponent.id},#{19})")
end

When("{user} creates a new match") do |user|
    @response = make_request("mutation { createMatch { id }}")
    @response_body = JSON.parse(@response.body)
end

When("{user} creates a new match without specifying board size") do |user|
    @response = make_request("mutation { createMatch { id, board { size }}}")
    @response_body = JSON.parse(@response.body)
end

When("{user} requests (her)(his) matches") do |user|
    @response = make_request("query { matches { nodes { id }}}")
    @response_body = JSON.parse(@response.body)
end

Then("{user} should get a new match") do |user|
    created_match = @response_body.dig "data", "createMatch"
    expect(created_match["id"]).to match($id_regexp)
end

Then("the board should be {int}x{int}") do |sizex, sizey|
    got_board = @response_body.dig "data", "createMatch", "board"
    got_board_size = got_board["size"]

    expect(got_board).to_not be_nil
    # Boards must be square
    expect(got_board_size).to eq(sizex)
    expect(got_board_size).to eq(sizey)
end

Then("{user} should get (her)(his) {int} match(es)") do |user, expected_number_of_matches|
    match_nodes = @response_body.dig "data", "matches", "nodes"

    expect(match_nodes.length).to eq(expected_number_of_matches)
    for match in match_nodes do
        expect(match["id"]).to match($id_regexp)
    end
end