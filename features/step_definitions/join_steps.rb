When("{user} joins {user}('s)( own) match") do |opponent, creator|
    # TODO find a way to make this less janky
    match_id = match_id_from_db_id($mysql_client.last_id)
    @response = make_request("mutation { joinMatch(matchId:\"#{match_id}\") { id opponent { id } status }}")
    @response_body = JSON.parse(@response.body)
end

When("{user} joins a match that does not exist") do |user|
    match_id = Base64.strict_encode64("Match:0")
    @response = make_request("mutation { joinMatch(matchId:\"#{match_id}\") { id opponent { id } status }}")
    @response_body = JSON.parse(@response.body)
end

Then("the match opponent should be {user}") do |user|
    got_opponent = @response_body.dig "data", "joinMatch", "opponent"
    user_id = encode_user_id(user.id)

    expect(got_opponent).to_not be_nil
    expect(got_opponent["id"]).to eq(user_id)
    # TODO add username
    #expect(got_opponent["username"]).to eq(user.username)
end

Then("the match status should be {word}") do |status|
    got_match = @response_body.dig "data", "joinMatch"
    expect(got_match["status"]).to eq(status)
end