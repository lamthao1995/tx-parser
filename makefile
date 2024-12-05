# Makefile for running the TxParser project

.PHONY: run

run:
	go run main.go
run:
	mockery --dir=domain --name=Parser --output=mocks --outpkg=mocks
run:
	mockery --dir=domain --name=Repository --output=mocks --outpkg=mocks