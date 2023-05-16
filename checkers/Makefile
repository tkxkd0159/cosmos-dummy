
.PHONY: mock-expected-keepers install-mockgen

mock-expected-keepers: install-mockgen
	mockgen -source=x/checkers/types/expected_keepers.go \
			-package testutil \
			-destination=x/checkers/testutil/expected_keepers_mocks.go

install-mockgen:
	go install github.com/golang/mock/mockgen@v1.6.0