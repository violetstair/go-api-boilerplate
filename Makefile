PROJECT_NAME=go-api-boilerplate
BUILD_DIR=$(shell pwd)/build
BINARY=${BUILD_DIR}/${PROJECT_NAME}
SOURCE=cmd/${PROJECT_NAME}/main.go
GOARCH=amd64
GOOS=linux

VERSION?=?
COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)

# Symlink into GOPATH
#BUILD_DIR=$(shell pwd)
CURRENT_DIR=$(shell pwd)
BUILD_DIR_LINK=$(shell readlink ${BUILD_DIR})

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS= -ldflags "-X main.VERSION=${VERSION} -X main.COMMIT=${COMMIT} -X main.BRANCH=${BRANCH}"

$(BUILD_DIR):
	if [ ! -d "$(BUILD_DIR)" ];then
		mkdir $(BUILD_DIR)
	fi

# Build the project
all: pre link clean test build

pre:
	dep ensure -update

link:
	BUILD_DIR=${BUILD_DIR}; \
	BUILD_DIR_LINK=${BUILD_DIR_LINK}; \
	CURRENT_DIR=${CURRENT_DIR}; \
	if [ "$${BUILD_DIR_LINK}" != "$${CURRENT_DIR}" ]; then \
		echo "Fixing symlinks for build"; \
		rm -f $${BUILD_DIR}; \
		ln -s $${CURRENT_DIR} $${BUILD_DIR}; \
	fi

build:
	GOOS=${GOOS} GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BINARY} ${SOURCE}

test:
	go test -v

clean:
	rm -f ${BINARY}

.PHONY: pre link build test clean