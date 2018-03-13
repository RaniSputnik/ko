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
    # This property is required for the user
    # parameter type. TODO is there a way to
    # change this back to an instance property?
    $current_user = user
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
        when "she", "he", "her", "his" then $current_user
        else raise "Unknown user: '#{name}'"
    end

    if !user then raise "Unknown user: '#{name}'" end
    user
end

# TODO how can I move these to their own file
# without getting 'uninitialized constant User (NameError)'

ParameterType({
    :name => 'user',
    :regexp => /Alice|Bob|Clive|she|he|her|his/,
    :type => User,
    :transformer => lambda {|s| get_user(s) }
})

class Move
    def initialize(name)
        first_letter = name[0,1]
        second_letter = name[1,2]
        @name = name
        @x = first_letter.ord - 'A'.ord
        @y = second_letter.to_i - 1
    end

    attr_reader :name
    attr_reader :x
    attr_reader :y
end

ParameterType({
    :name => 'move',
    :regexp => /[A-Z][1-9]/,
    :type => Move,
    :transformer => lambda {|s| Move.new(s) }
})