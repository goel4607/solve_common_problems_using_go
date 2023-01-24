test:
	go test ./...

test_coverage:
	#go test ./... --coverprofile=coverage.out
	go test ./... --cover
