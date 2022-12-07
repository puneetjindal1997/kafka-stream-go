GOBUILD=go build
GOTEST=go test

dependencies: dep_server dep_client

dep_server:
	cd server && go get
dep_client:
	cd client && go get

all:clean stop build_server build_client
	server/server &
	client/client &
build_server:
	cd server && $(GOBUILD) -v .
build_client:
	cd client && $(GOBUILD) -v .
build_migrations:
	cd migrations && $(GOBUILD) -v .
clean:
	rm -f server/server
	rm -f client/client
stop:
	pkill server || true
	pkill client || true
test:
	cd client && $(GOTEST) -v .