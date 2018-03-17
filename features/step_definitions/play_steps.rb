Given("{user} is playing against {user}") do |creator, opponent|
    $mysql_client.query("INSERT INTO Matches (Owner, Opponent, BoardSize) VALUES (#{creator.id},#{opponent.id},#{19})")
end

When("{user} places a stone at {move}") do |user, move|
    match_id = match_id_from_db_id($mysql_client.last_id)
    @response = make_request("mutation { playStone(matchId:\"#{match_id}\", x:#{move.x}, y:#{move.y}) { message x y }}")
    @response_body = JSON.parse(@response.body)
    puts @response_body
end

Then("there should be a move at {int},{int}") do |movex, movey|
    event = @response_body.dig "data", "playStone"
    expect(event).to_not be_nil
    expect(event.x).to be(movex)
    expect(event.y).to be(movey)
end

Given("the folling moves have been played: {moves}") do |moves|
    moves.each do |move|
        puts "TODO put move {#{move.x},#{move.y}} in the database"
    end
end