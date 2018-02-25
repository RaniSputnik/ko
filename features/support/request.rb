require 'net/http'
require 'json'

def make_request(query)
    uri = URI('http://localhost:8080/graphql')
    req = Net::HTTP::Post.new(uri, 'Content-Type' => 'application/json')
    req.body = {query: query}.to_json
    res = Net::HTTP.start(uri.hostname, uri.port) do |http|
      http.request(req)
    end
    res
end