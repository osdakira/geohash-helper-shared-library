.PHONY: build

build:
	os=$$(uname -s | tr '[A-Z]' '[a-z]')
	arch=$$(uname -m | tr '[A-Z]' '[a-z]')
	arch=$$([[ "${arch}" = "x86_64" ]] && echo "amd64" || echo "${arch}")
	go build -buildmode=c-shared -o "lib/build/${os}_${arch}_geohash_lib_go.so" ./geohash_lib.go

test:
	go test -v ./

fmt:
	go fmt ./...
