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
$Clive = User.new(3, "Clive")

def login_user(name)
    user = get_user(name)
    @current_user = user
    payload = {:sub => encode_user_id(user.id) }
    @auth_token = JWT.encode payload, nil, 'none'
end

def encode_user_id(user_id)
    Base64.strict_encode64("User:#{user_id}")
end

def get_user(name)
    user = case name.downcase
        when "alice" then $Alice
        when "bob" then $Bob
        when "clive" then $Clive
        when "she", "he" then @current_user
        else raise "Unknown user: '#{name}'"
    end

    if !user then raise "Unknown user: '#{name}'" end
    user
end