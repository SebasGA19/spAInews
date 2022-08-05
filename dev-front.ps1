# Remove any running instance
docker-compose -f .\compose-dev-front.yaml down --remove-orphans -v
# Start the new dev instance locally
docker-compose -f .\compose-dev-front.yaml up -d --build
# Start Frontend
# TODO: Implement me!