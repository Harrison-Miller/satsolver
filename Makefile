test:
	go test -v .

bench:
	go test -run=xxx -bench=.

build:
	go build -o generator cmd/generator/generator.go
	go build -o solver cmd/solver/solver.go