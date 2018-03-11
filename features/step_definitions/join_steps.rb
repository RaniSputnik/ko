When("Bob joins Alice's match") do
    # TODO find a way to make this less janky
    match_db_id = $mysql_client.last_id
    match_id = Base64.strict_encode64("Match:#{match_db_id}")
    @response = make_request("mutation { joinMatch(matchId:\"#{match_id}\") { id opponent { id } status }}")
    @response_body = JSON.parse(@response.body)
end

Then("the match opponent should be Bob") do
    got_opponent = @response_body.dig "data", "joinMatch", "opponent"
    bob = get_user("Bob")
    bobs_id = encode_user_id(bob.id)

    expect(got_opponent).to_not be_nil
    expect(got_opponent["id"]).to eq(bobs_id)
    # TODO add username
    #expect(got_opponent["username"]).to eq(bobs_username)
end

Then("the match status should be {word}") do |status|
    got_match = @response_body.dig "data", "joinMatch"
    expect(got_match["status"]).to eq(status)
end