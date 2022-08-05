# Setup Env
$env:SMTP_HOST = ""
$env:SMTP_PORT = ""
$env:SMTP_FROM = "dashboard@dev-spainews.co"
$env:SMTP_USERNAME = "dashboard@dev-spainews.co"
$env:SMTP_PASSWORD = "password"
$env:REDIS_ADDRESS = "127.0.0.1:6379"
$env:MARIA_URL = "spainews:spainews@tcp(127.0.0.1:3306)/spainews?parseTime=true"
$env:MONGO_URL = "mongodb://api:api@127.0.0.1:27017/spainews"
$env:LISTEN_ADDRESS = "127.0.0.1:5000"
# Start Backend
$path = Split-Path $MyInvocation.MyCommand.Path -Parent
go run "$path/api/cmd/api"