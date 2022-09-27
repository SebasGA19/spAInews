# Remove dependency folders
Remove-Item -Recurse -Force "mariadb-data"
Remove-Item -Recurse -Force "mongo-data"
# Remove any running instance
docker-compose -f .\compose-dev-databases.yaml down --remove-orphans -v
# Start the new dev instance locally
docker-compose -f .\compose-dev-databases.yaml up -d --build
