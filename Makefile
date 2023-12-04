mock-expected-keepers:
	mockgen -source=x/checkers/types/expected_keepers.go \
		-package testutil \
		-destination=x/checkers/testutil/expected_keepers_mocks.go 

test:
	go test -v ./x/checkers/...

test-i:
	go test -v github.com/alice/checkers/tests/integration/checkers/keeper