Given("{word} is logged in") do |user_name|
    login_user(user_name)
end

# TODO define custom parameter type 'user'
Given("{word} has created {int} match(es)") do |user_name, n_matches|
    user = get_user(user_name)
    for i in 1..n_matches do
        $mysql_client.query("INSERT INTO Matches (Owner, BoardSize) VALUES (#{user.id},#{19})")
    end
end

When(/^(?:she|he) creates a new match$/) do
    @response = make_request("mutation { createMatch { id }}")
    @response_body = JSON.parse(@response.body)
end

When(/^(?:she|he) creates a new match without specifying board size$/) do
    @response = make_request("mutation { createMatch { id, board { size }}}")
    @response_body = JSON.parse(@response.body)
end

When("Alice requests her matches") do
    @response = make_request("query { matches { nodes { id }}}")
    @response_body = JSON.parse(@response.body)
end

Then(/^Alice should get a new match$/) do
    created_match = @response_body.dig "data", "createMatch"
    expect(created_match["id"]).to match($id_regexp)
end

Then(/^there should be no errors$/) do
    expect(@response_body["errors"]).to eq(nil)
end

Then("the board should be {int}x{int}") do |sizex, sizey|
    got_board = @response_body.dig "data", "createMatch", "board"
    got_board_size = got_board["size"]

    expect(got_board).to_not be_nil
    # Boards must be square
    expect(got_board_size).to eq(sizex)
    expect(got_board_size).to eq(sizey)
end

Then("she should get her {int} match(es)") do |expected_number_of_matches|
    match_nodes = @response_body.dig "data", "matches", "nodes"

    expect(match_nodes.length).to eq(expected_number_of_matches)
    for match in match_nodes do
        expect(match["id"]).to match($id_regexp)
    end
end