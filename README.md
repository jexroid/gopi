# gopi

this project was made with GORMand postgres
hashing strategy is argon2 because of speed efficiency

## env management
for this project `.env` was available for demo
but i recomend using env-vault for managing the project as the vault address and other details are available in the [Link](https://www.dotenv.org/)

## How can I test the Endpoints?
there is a JSON file that is used for testing the endpoints which are written with small test cases

## Testing
the testing unit lacks good Mocking techniques due to the lack of time.

## Docker image
the Docker image of this file is available at [dockerhub](https://hub.docker.com/repository/docker/jextoid/gauth-gopi/general)

## Security features:
 - HTTP headers
 - CORS
 - XSS protection
 - NoSniff (Don't let anyone sniff the credentials endpoints)
 - CSRF (Not implemented because of lack of frontend services)

   also remember to change the GO_ENV to production on production level. (you can makethis automatic in env-vault)
