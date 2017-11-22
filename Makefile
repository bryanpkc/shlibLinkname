all:
	cd linkname1 && go install -buildmode=shared -linkshared
	go run -linkshared linkname2.go
