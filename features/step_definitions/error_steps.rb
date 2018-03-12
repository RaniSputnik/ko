Then(/^there should be no errors$/) do
    expect(@response_body["errors"]).to eq(nil)
end

Then("{word} should get an error: {string}") do |user_name, error_message|
    errors = @response_body["errors"]

    expect(errors.length).to eq(1)
    expect(errors[0]["message"]).to eq(error_message)
end