CC := go

TRASH := client_gen.go

all: client_gen.go
	$(CC) build

client_gen.go: api.yaml cfg.yaml generate.go
	$(CC) generate

oapi-gen:
	@openapi-generator generate -g go -o oapigen -i api.yaml \
		--git-host github.com --git-user-id pcolladosoto --git-repo-id goslurm/oapigen \
		-c cfg.json

oapi-gen-ro:
	@openapi-generator generate -g go -o oapigenro -i api_ro.yaml \
		--git-host github.com --git-user-id pcolladosoto --git-repo-id goslurm/oapigenro \
		-c cfg.json

.PHONY: clean

clean:
	rm -f $(TRASH)
