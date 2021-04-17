
.PHONY: usage
usage:
	@echo This shows some ways you can use the Makefile
	@echo make test               - run golang tests
	@echo make example1           - run a example using cmd/sl2tab

.PHONY: test
test:
	go test -cover


.PHONY: example1
example1:
	go run ./cmd/sl2tab/ -count 20 "./test-fixtures/sample-data-lowrance/Elite_4_Chirp\Chart 05_11_2018 [0].sl2"
