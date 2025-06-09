.PHONY: all
all: build

.PHONY: build
build: bin
	go build -o bin/fedtpd ./cmd/fedtpd
	go build -o bin/login ./cmd/login
	go build -o bin/perivale ./cmd/perivale
	go build -o bin/telnetd ./cmd/telnetd
	go build -o bin/workbench ./cmd/workbench

bin:
	@mkdir $@

.PHONY: run
run: build ssl
	docker buildx build --tag=go-federation:latest deployments/docker
	./scripts/run.sh

.PHONY: ssl
ssl: ssl/telnetd.crt

ssl/telnetd.crt:
	-mkdir -p ssl
	openssl req -x509 -newkey rsa:4096 -keyout $(patsubst %.crt,%.key,$@) -out $@ -sha256 -days 365 -nodes -subj /CN=localhost -quiet

.PHONY: clean
clean:
	-$(RM) -r bin

.PHONY: test
test: # TODO
