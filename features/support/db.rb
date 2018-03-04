require "mysql2"

mysql_host = ENV.fetch("KO_SQL_HOST", "localhost")
mysql_port = ENV.fetch("KO_SQL_PORT", "3306")
mysql_db   = ENV.fetch("KO_SQL_DB", "ko")
mysql_user = ENV.fetch("KO_SQL_USER", "root")
mysql_pwd  = ENV.fetch("KO_SQL_PWD", "example")

puts "Connecting to host: #{mysql_host}:#{mysql_port}"
$mysql_client = Mysql2::Client.new(
    :host => mysql_host, 
    :port => mysql_port, 
    :database => mysql_db, 
    :username => mysql_user, 
    :password => mysql_pwd
)

def clear_db()
    $mysql_client.query("DELETE FROM Matches")
end