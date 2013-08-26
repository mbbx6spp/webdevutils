EXECUTABLES := autocompile staticserve
CONFIGDIR		:= $(PWD)/config
VERSION			:= 0.1.0

all: test

compile:
	for exe in $(EXECUTABLES); do \
		go build -o build/$$exe src/$$exe.go ; \
		upx -qq build/$$exe ; \
	done

test: compile
	env CONFIGDIR=$(CONFIGDIR) go test webdevutils
	cucumber

package: linux_arm linux_amd64 linux_386

linux_386:
	mkdir -p build/linux_386
	export GOOS=linux GOARCH=386
	go build -o build/$$GOOS_$$GOARCH/staticserve src/staticserve.go
	go build -o build/$$GOOS_$$GOARCH/autocompile src/autocompile.go
	fpm -s dir -t rpm -n "webdevutils" -v $(VERSION) -a i386 --prefix /usr/bin build/$$GOOS_$$GOARCH
	fpm -s dir -t deb -n "webdevutils" -v $(VERSION) -a i386 --prefix /usr/bin build/$$GOOS_$$GOARCH
	unset GOOS GOARCH

linux_amd64:
	mkdir -p build/linux_amd64
	export GOOS=linux GOARCH=amd64
	go build -o build/$$GOOS_$$GOARCH/staticserve src/staticserve.go
	go build -o build/$$GOOS_$$GOARCH/autocompile src/autocompile.go
	fpm -s dir -t rpm -n "webdevutils" -v $(VERSION) -a x86_64 --prefix /usr/bin build/$$GOOS_$$GOARCH
	fpm -s dir -t deb -n "webdevutils" -v $(VERSION) -a x86_64 --prefix /usr/bin build/$$GOOS_$$GOARCH
	unset GOOS GOARCH

linux_arm:
	mkdir -p build/linux_arm
	export GOOS=linux GOARCH=arm
	go build -o build/$$GOOS_$$GOARCH/staticserve src/staticserve.go
	go build -o build/$$GOOS_$$GOARCH/autocompile src/autocompile.go
	fpm -s dir -t rpm -n "webdevutils" -v $(VERSION) -a x86_64 --prefix /usr/bin build/$$GOOS_$$GOARCH
	fpm -s dir -t deb -n "webdevutils" -v $(VERSION) -a x86_64 --prefix /usr/bin build/$$GOOS_$$GOARCH
	unset GOOS GOARCH
