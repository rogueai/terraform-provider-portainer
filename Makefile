default: install

generate:
	go generate ./...

install:
	go install .

test:
	go test -count=1 -parallel=4 ./portainer

testacc:
	TF_ACC=1 go test -count=1 -parallel=4 -timeout 10m -v ./portainer

.PHONY: testacc
