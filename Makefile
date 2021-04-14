PKGS=$(shell go list ./... | grep -v /vendor/)

test:
	@ mkdir -p build
	@ go test $(PKGS) | tee build/test.log
	@ go2xunit -input build/test.log -output build/tests.xml

clean:
	-rm -r build/*

.PHONY: clean
