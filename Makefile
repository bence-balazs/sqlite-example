build:
	go build -o bin/app
run: build
	(cd ./bin && exec ./app)
clean:
	(cd ./bin && rm -rf *)