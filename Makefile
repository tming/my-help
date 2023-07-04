.PHONY: tool

GITTAG=$(shell git describe --tags --always)
BUILDTIME = $(shell date +%Y-%m-%dT%T%z)
GITHASH=$(shell git rev-parse HEAD)
VERSION=${GITTAG}-$(shell date +%y.%m.%d)

LDFLAG=-s -w -X my-help/src/common.Version=${VERSION} \
 -X my-help/src/common.BuildTime=${BUILDTIME} \
 -X my-help/src/common.GitHash=${GITHASH} \
 -X my-help/src/common.Tag=${GITTAG}

BIN_PATH=./build

tool:
	mkdir -p ${BIN_PATH}
	go build -ldflags "${LDFLAG}" -o ${BIN_PATH}/my-help ./src/main.go

clean:
	rm -rf ${BIN_PATH}
