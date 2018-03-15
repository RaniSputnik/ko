
def match_id_from_db_id(db_id)
    Base64.strict_encode64("Match:#{db_id}")
end

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
    :regexp => /[A-Z][1-9][1-9]?/,
    :type => Move,
    :transformer => lambda {|s| Move.new(s) }
})

ParameterType({
    :name => 'moves',
    :regexp => /((?:[A-Z][1-9][1-9]?,?)+)/,
    :type => Array,
    :transformer => lambda {|s| s.split(',').map {|m| Move.new(m)} }
})