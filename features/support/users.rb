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