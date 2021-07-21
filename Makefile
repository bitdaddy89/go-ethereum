# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: getd android ios getd-cross evm all test clean
.PHONY: getd-linux getd-linux-386 getd-linux-amd64 getd-linux-mips64 getd-linux-mips64le
.PHONY: getd-linux-arm getd-linux-arm-5 getd-linux-arm-6 getd-linux-arm-7 getd-linux-arm64
.PHONY: getd-darwin getd-darwin-386 getd-darwin-amd64
.PHONY: getd-windows getd-windows-386 getd-windows-amd64

GOBIN = ./build/bin
GO ?= latest
GORUN = env GO111MODULE=on go run

getd:
	$(GORUN) build/ci.go install ./cmd/getd
	@echo "Done building."
	@echo "Run \"$(GOBIN)/getd\" to launch getd."

all:
	$(GORUN) build/ci.go install

android:
	$(GORUN) build/ci.go aar --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/getd.aar\" to use the library."
	@echo "Import \"$(GOBIN)/getd-sources.jar\" to add javadocs"
	@echo "For more info see https://stackoverflow.com/questions/20994336/android-studio-how-to-attach-javadoc"

ios:
	$(GORUN) build/ci.go xcode --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/Getd.framework\" to use the library."

test: all
	$(GORUN) build/ci.go test

lint: ## Run linters.
	$(GORUN) build/ci.go lint

clean:
	env GO111MODULE=on go clean -cache
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# The devtools target installs tools required for 'go generate'.
# You need to put $GOBIN (or $GOPATH/bin) in your PATH to use 'go generate'.

devtools:
	env GOBIN= go install golang.org/x/tools/cmd/stringer@latest
	env GOBIN= go install github.com/kevinburke/go-bindata/go-bindata@latest
	env GOBIN= go install github.com/fjl/gencodec@latest
	env GOBIN= go install github.com/golang/protobuf/protoc-gen-go@latest
	env GOBIN= go install ./cmd/abigen
	@type "solc" 2> /dev/null || echo 'Please install solc'
	@type "protoc" 2> /dev/null || echo 'Please install protoc'

# Cross Compilation Targets (xgo)

getd-cross: getd-linux getd-darwin getd-windows getd-android getd-ios
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/getd-*

getd-linux: getd-linux-386 getd-linux-amd64 getd-linux-arm getd-linux-mips64 getd-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/getd-linux-*

getd-linux-386:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/getd
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/getd-linux-* | grep 386

getd-linux-amd64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/getd
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/getd-linux-* | grep amd64

getd-linux-arm: getd-linux-arm-5 getd-linux-arm-6 getd-linux-arm-7 getd-linux-arm64
	@echo "Linux ARM cross compilation done:"
	@ls -ld $(GOBIN)/getd-linux-* | grep arm

getd-linux-arm-5:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/arm-5 -v ./cmd/getd
	@echo "Linux ARMv5 cross compilation done:"
	@ls -ld $(GOBIN)/getd-linux-* | grep arm-5

getd-linux-arm-6:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/arm-6 -v ./cmd/getd
	@echo "Linux ARMv6 cross compilation done:"
	@ls -ld $(GOBIN)/getd-linux-* | grep arm-6

getd-linux-arm-7:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/arm-7 -v ./cmd/getd
	@echo "Linux ARMv7 cross compilation done:"
	@ls -ld $(GOBIN)/getd-linux-* | grep arm-7

getd-linux-arm64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/arm64 -v ./cmd/getd
	@echo "Linux ARM64 cross compilation done:"
	@ls -ld $(GOBIN)/getd-linux-* | grep arm64

getd-linux-mips:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/getd
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/getd-linux-* | grep mips

getd-linux-mipsle:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/getd
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/getd-linux-* | grep mipsle

getd-linux-mips64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/getd
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/getd-linux-* | grep mips64

getd-linux-mips64le:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/getd
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/getd-linux-* | grep mips64le

getd-darwin: getd-darwin-386 getd-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/getd-darwin-*

getd-darwin-386:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/getd
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/getd-darwin-* | grep 386

getd-darwin-amd64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/getd
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/getd-darwin-* | grep amd64

getd-windows: getd-windows-386 getd-windows-amd64
	@echo "Windows cross compilation done:"
	@ls -ld $(GOBIN)/getd-windows-*

getd-windows-386:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=windows/386 -v ./cmd/getd
	@echo "Windows 386 cross compilation done:"
	@ls -ld $(GOBIN)/getd-windows-* | grep 386

getd-windows-amd64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=windows/amd64 -v ./cmd/getd
	@echo "Windows amd64 cross compilation done:"
	@ls -ld $(GOBIN)/getd-windows-* | grep amd64
