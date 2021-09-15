
.PHONY: run
run:
	$(info #Building...)
	@$(MAKE) generate
	@$(MAKE) start-server


.PHONY: start-server
start-server:
	$(info #Start Application...)
	go run ./cmd/application/main.go

PHONY: generate
generate:
	$(info #Generate JSON file...)
	go run generate.go