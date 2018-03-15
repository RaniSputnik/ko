Given("{user} is playing against {user}") do |creator, opponent|
    $mysql_client.query("INSERT INTO Matches (Owner, Opponent, BoardSize) VALUES (#{creator.id},#{opponent.id},#{19})")
end

When("{user} places a stone at {move}") do |user, move|
    match_id = match_id_from_db_id($mysql_client.last_id)
    @response = make_request("mutation { playStone(matchId:\"#{match_id}\", x:#{move.x}, y:#{move.y}) { id events { nodes { message }}}}")
    @response_body = JSON.parse(@response.body)
end

Then("there should be {int} move(s) played") do |number_of_moves|
    events = @response_body.dig "data", "playStone", "events", "nodes"
    expect(events).to_not be_nil
    expect(events.length).to be(number_of_moves)
end

Given("the folling moves have been played: {moves}") do |moves|
    moves.each do |move|
        puts "TODO put move {#{move.x},#{move.y}} in the database"
    end
end