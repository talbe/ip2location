Ip 2 location assignment

prerequisites:
in order to use ip 2 location, you should set couple of environment variable-
API_VERSION=v1
DATA_STORE_PATH=<PATH_TO_DATASTORE_FILE>
FIND_COUNTRY_API=find-country
SERVER_PORT=8080
VISITOR_MINUTES_TO_LIVE=3
BUCKET_SIZE=3
TOKEN_INCREASE_RATE=1

build:
"go build" from the root directory.

test:
There are 2 tests files - in network/rate/limiter_test.go and handlers/iplocation_test.go.
You can run "go test" to run the tests/