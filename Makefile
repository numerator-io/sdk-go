# Commands to run different Go tests
go-unit-tests:
	go test -v ./... -count=1 -skip "Router|Repository" && echo "Unit tests passed"
go-integration-tests:
	INTEGRATION_TEST=1 go test -v ./... -run "Router" -count=1 && echo "Integration tests passed"
# Commands to stand up tests environments and run tests
unit-test: go-unit-tests
integration-test: go-integration-tests
