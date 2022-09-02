app=nblog

build:
	@go build -o output/bin/${app}

run:
	output/bin/${app}

clean:
	rm -rf output