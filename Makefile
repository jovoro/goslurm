CC := go

TRASH := client_gen.go

all: client_gen.go
	$(CC) build

client_gen.go: api.yaml cfg.yaml generate.go
	$(CC) generate

.PHONY: clean

clean:
	rm -f $(TRASH)
