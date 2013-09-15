EXECUTABLES := autocompile staticserve
CONFIGDIR		:= $(PWD)/config
VERSION			:= 0.1.0

all: test

compile:
	for exe in $(EXECUTABLES); do \
		go build -o build/$$exe src/$$exe.go ; \
	done
		#upx -qq build/$$exe ;

test: compile
	env CONFIGDIR=$(CONFIGDIR) go test webdevutils
	cucumber

clean:
	rm -rf build/*
	rm webdevutils*.rpm
	rm webdevutils*.deb

package: linux_amd64

linux_386:
	export GOOS=linux GOARCH=386
	go build -o build/staticserve src/staticserve.go
	go build -o build/autocompile src/autocompile.go
	fpm -s dir -t rpm -n "webdevutils" -C build -v $(VERSION) \
		-a i386 --prefix /usr/bin .
	fpm -s dir -t deb -n "webdevutils" -C build -v $(VERSION) \
		-a i386 --prefix /usr/bin .
	unset GOOS GOARCH

linux_amd64:
	export GOOS=linux GOARCH=amd64
	go build -o build/staticserve src/staticserve.go
	go build -o build/autocompile src/autocompile.go
	fpm -s dir -t rpm -n "webdevutils" -C build -v $(VERSION) \
		-a x86_64 --prefix /usr/bin .
	fpm -s dir -t deb -n "webdevutils" -C build -v $(VERSION) \
		-a x86_64 --prefix /usr/bin .
	unset GOOS GOARCH

linux_arm:
	export GOOS=linux GOARCH=arm
	go build -o build/staticserve src/staticserve.go
	go build -o build/autocompile src/autocompile.go
	fpm -s dir -t rpm -n "webdevutils" -C build -v $(VERSION) \
		-a x86_64 --prefix /usr/bin .
	fpm -s dir -t deb -n "webdevutils" -C build -v $(VERSION) \
		-a x86_64 --prefix /usr/bin .
	unset GOOS GOARCH
