# Start Frontend
$path = Split-Path $MyInvocation.MyCommand.Path -Parent
npm start --prefix "$path/frontend"