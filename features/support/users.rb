require "jwt"
require "base64"

class User
    def initialize(id, name)
        @id = id
        @name = name
    end

    attr_reader :id
    attr_reader :name
end

$Alice = User.new(1, "Alice")
$Bob = User.new(2, "Bob")

def login_user(name) 
    user = get_user(name)
    @current_user = user
    payload = {:sub => Base64.strict_encode64("User:#{user.id}") }
    @auth_token = JWT.encode payload, nil, 'none'
end

def get_user(name)
    user = case name.downcase
        when "alice" then $Alice
        when "bob" then $Bob
        when "she", "he" then @current_user
        else raise "Unknown user: '#{name}'"
    end

    if !user then raise "Unknown user: '#{name}'" end
    user
end