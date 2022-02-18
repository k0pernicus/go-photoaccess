prepare:
	cp config.yaml app/

build:
	cd app; go build -o photoaccess

run: prepare build
	cd app; ./photoaccess

clean:
	cd app; go clean
	cd app; rm config.yaml; rm photoaccess