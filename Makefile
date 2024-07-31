# The compiler to use, which is, well, Go!
CC := go

help:
	@echo "usage: make <target>"
	@echo "targets:"
	@echo "  help         show this message and exit"
	@echo ""
	@echo "  v0038        generate the v0.0.38 client based on ./specs/0038_ro.yaml"
	@echo ""
	@echo "  example      compile the example program under ./cmd/"
	@echo ""
	@echo "  clean        delete every built executable under ./cmd/bin"

v0038:
	@openapi-generator generate -g go -o v0038-f -i api_ro.yaml \
		--git-host github.com --git-user-id pcolladosoto --git-repo-id goslurm/v0038 \
		--additional-properties=packageName=v0038,withGoMod=false

example:
	@echo "TODO..."

.PHONY: clean

clean:
	@echo "TODO..."
	# @rm -f $(TRASH)
