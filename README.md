# gopi

this project was made with GORM and postgres
hashing strategy is argon2 because of speed efficiency

## env management
for this project `.env` is available for demo but you can configure env-vault to download the `.env` for you
but i recomend using env-vault for managing the environment variables of the project as the details are available in the [Link](https://www.dotenv.org/)

## How can I test the Endpoints?
there is a JSON file named [go Auth.postman_collection](https://github.com/jexroid/gopi/blob/main/go%20Auth.postman_collection.json) that is used for testing the endpoints which are written with small test cases

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
