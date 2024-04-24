b:
	go build -o pingchecker main.go

r:
	go run main.go

run:
	./pingchecker -f ./input.csv -w 10

release:
	goreleaser release --snapshot --clean 

build:
	goreleaser build --single-target --snapshot --clean

