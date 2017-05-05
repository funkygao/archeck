all:runcheck

build:
	@go generate
	@go install

runcheck:build
	@archeck
