
.PHONY: usage
usage:
	@echo This shows some ways you can use the Makefile
	@echo make test               - run golang tests
	@echo make example1           - run a example using cmd/sl2tab

.PHONY: test
test:
	go test -coverprofile=coverage.out

cover:
	go tool cover -func=coverage.out


.PHONY: example1
example1:
	go run ./cmd/sl2tab/ -count 20 "./testdata/sample-data-lowrance/Elite_4_Chirp/Chart 05_11_2018 [0].sl2"

.PHONY: example2
example2:
	go run ./cmd/sl3tab/ -count 20 "./testdata/sample-data-lowrance/other/format3_version2.sl3"
