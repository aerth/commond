all: bin/server bin/complex
bin/server: *.go tester/server/*.go
	GOBIN=${PWD}/bin go install -v ./tester/server
bin/complex: *.go tester/complex/*.go
	GOBIN=${PWD}/bin go install -v ./tester/complex
