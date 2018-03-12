When("{word} joins {word}('s)( own) match") do |user_name, creator_name|
    # TODO find a way to make this less janky
    match_db_id = $mysql_client.last_id
    match_id = Base64.strict_encode64("Match:#{match_db_id}")
    @response = make_request("mutation { joinMatch(matchId:\"#{match_id}\") { id opponent { id } status }}")
    @response_body = JSON.parse(@response.body)
end

When("{word} joins a match that does not exist") do |user_name|
    match_id = Base64.strict_encode64("Match:0")
    @response = make_request("mutation { joinMatch(matchId:\"#{match_id}\") { id opponent { id } status }}")
    @response_body = JSON.parse(@response.body)
end

Then("the match opponent should be {word}") do |user_name|
    got_opponent = @response_body.dig "data", "joinMatch", "opponent"
    user = get_user(user_name)
    user_id = encode_user_id(user.id)

    expect(got_opponent).to_not be_nil
    expect(got_opponent["id"]).to eq(user_id)
    # TODO add username
    #expect(got_opponent["username"]).to eq(user_username)
end

Then("the match status should be {word}") do |status|
    got_match = @response_body.dig "data", "joinMatch"
    expect(got_match["status"]).to eq(status)
end