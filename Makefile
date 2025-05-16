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

.PHONY: v0038 v0041 v0041-docker clean

# We could maybe use a template when adding different/new versions?
# Also, note you can specify several `--additional-properties` flags instead
# of concatenating options with commas.
v0038:
	@openapi-generator generate -g go -o $@ -i ./specs/$@.yaml                    \
		--git-host github.com --git-user-id pcolladosoto --git-repo-id goslurm/$@ \
		--additional-properties=packageName=$@,withGoMod=false

v0041:
	@openapi-generator generate -g go -o $@ -i ./specs/$@.json                    \
		--git-host github.com --git-user-id pcolladosoto --git-repo-id goslurm/$@ \
		--additional-properties=packageName=$@,withGoMod=false
v0041-docker:
	@docker run --rm -v "$${PWD}:/local" openapitools/openapi-generator-cli generate \
		-i /local/specs/v0041.json -g go -o /local/v0041 \
		--additional-properties=packageName=v0041,withGoMod=false

example:
	@echo "TODO..."

clean:
	@echo "TODO..."
	# @rm -f $(TRASH)
