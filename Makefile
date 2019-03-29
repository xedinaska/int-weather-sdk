#!make

fmt:
	goimports -w $$(ls -d */ | grep -v vendor)
	goimports -w integration.go

test:
	go test -v --cover --race -short `glide novendor | grep -v ./proto`

deps:
	glide install