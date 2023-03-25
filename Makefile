default: install

generate:
	go generate ./...

install:
	go install .

test:
	go test -count=1 -parallel=4 ./portainer

testacc:
	@echo "Starting mock server"
	./testacc/start-prism.sh
	sleep 3
	TF_ACC=1 go test -count=1 -parallel=4 -timeout 10m -v ./portainer
	if [ -e prism.PID ]; then \
    	kill -TERM $$(cat prism.PID) || true; \
		rm prism.PID; \
	fi;

.PHONY: testacc
