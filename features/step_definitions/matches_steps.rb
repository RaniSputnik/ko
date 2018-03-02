Given(/^Alice is logged in$/) do
    #set_user "Alice"
end

Given("she has created {int} matches") do |int|
    pending # Write code here that turns the phrase above into concrete actions
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
    base64encodedRegexp = /^(?:[A-Za-z0-9+\/]{4}\n?)*(?:[A-Za-z0-9+\/]{2}==|[A-Za-z0-9+\/]{3}=)?$/

    expect(@response_body["data"]).to_not be_nil
    expect(@response_body["data"]["createMatch"]).to_not be_nil
    expect(@response_body["data"]["createMatch"]["id"]).to match(base64encodedRegexp)
end

Then(/^there should be no errors$/) do
    expect(@response_body["errors"]).to eq(nil)
end

Then("the board should be {int}x{int}") do |sizex, sizey|
    expect(@response_body["data"]).to_not be_nil
    expect(@response_body["data"]["createMatch"]).to_not be_nil
    expect(@response_body["data"]["createMatch"]["board"]).to_not be_nil

    gotBoardSize = @response_body["data"]["createMatch"]["board"]["size"]
    # Boards must be square
    expect(gotBoardSize).to eq(sizex)
    expect(gotBoardSize).to eq(sizey)
end

Then("she should get her {int} matches") do |int|
    pending # Write code here that turns the phrase above into concrete actions
end