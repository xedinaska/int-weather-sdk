#!make

lint:
	goimports -w $$(ls -d */ | grep -v vendor)
	golint $$(ls -d */ | grep -v vendor)
	gocyclo -over 10  $$(ls -d */ | grep -v vendor)

test:
	go test -v --cover --race -short `glide novendor | grep -v ./proto`

deps:
	glide install