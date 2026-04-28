compile:
	go build -o ./bin/main ./game/Main.go
run: 
	go run ./game/Main.go
run_compile: 
	./bin/main
