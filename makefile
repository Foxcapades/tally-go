.PHONY: test
test: test-v1

.PHONY: test-v1
test-v1:
	go test -v ./v1/tally

.PHONY: gen
gen:

.PHONY: gen-v1
gen-v1:
	go run v1/tools/gen/main.go