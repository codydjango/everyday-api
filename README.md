# Everday API

This is a golang authentication microservice for the everday app. 

Functionality is pretty basic: it takes a request, returns a message to be signed by a users' ethereum private key, verifies with the provided public key, and if it's correct returns a JWT with the appropriate session data. 

### Development notes

`source ./bin/scripts.sh`

`build && start`
