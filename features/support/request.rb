require 'net/http'
require 'json'

$api_endpoint = ENV.fetch('KO_API_ENDPOINT', 'http://localhost:8080/graphql')

def make_request(query)
    uri = URI($api_endpoint)
    req = Net::HTTP::Post.new(uri, 'Content-Type' => 'application/json')
    req.body = {query: query}.to_json
    res = Net::HTTP.start(uri.hostname, uri.port) do |http|
      http.request(req)
    end
    res
end