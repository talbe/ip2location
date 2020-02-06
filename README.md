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
There are 3 tests files - 

Main tests: server_tests.go - system testing to verify that the server api work as expected

network/rate/limiter_test.go  - unit testing to verify that the rate limit engine works as epected

handlers/iplocation_test.go = unit testing to verify that the iplocation handler working as epected

You can run "go test" to run the tests/