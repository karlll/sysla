SRC=src/sysla
PEG=peg
BUILD=go build
TEST=go test

${SRC}/sysla.peg.go: ${SRC}/sysla.peg
	$(PEG) ${SRC}/sysla.peg

build: ${SRC}/sysla.peg.go ${SRC}/sysla_test.go
	$(BUILD) sysla

test: build
	${TEST} sysla

clean:
	rm -rf ${SRC}/sysla.peg.go sysla
