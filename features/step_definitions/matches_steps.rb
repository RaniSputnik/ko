Given(/^Alice is logged in$/) do
    #set_user "Alice"
end

When(/^(?:she|he) creates a new match$/) do
    @response = make_request("mutation { createMatch { id }}")
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