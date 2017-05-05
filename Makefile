all:runcheck

build:
	@go install

runcheck:build
	@archeck
