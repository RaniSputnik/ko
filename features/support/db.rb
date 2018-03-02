require 'mysql2'

def clear_db()
    client = Mysql2::Client.new(:host => "mysql", :port => "3306", :database => "ko", :username => "root", :password => "example")
    puts client.query("SELECT * FROM Matches")
end