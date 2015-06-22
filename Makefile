
sysla.peg.go: sysla.peg
	peg -inline -switch sysla.peg

build: sysla.peg.go sysla_test.go
	go build

test: build
	go test
