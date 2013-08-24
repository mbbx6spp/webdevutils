EXECUTABLES := autocompile staticserve
CONFIGDIR		:= $(PWD)/config

all: test

compile:
	for exe in $(EXECUTABLES); do \
		go build -o build/$$exe src/$$exe.go ; \
		upx -qq build/$$exe ; \
	done

test: compile
	env CONFIGDIR=$(CONFIGDIR) go test webdevutils
	cucumber
