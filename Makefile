prepare:
	cp config.yaml app/

build:
	cd app; go get ./...; go build -o photoaccess

run: prepare build
	cd app; ./photoaccess

doc:
	cd app; godoc -http=127.0.0.1:6060

clean:
	cd app; go clean
	cd app; rm config.yaml; rm photoaccess