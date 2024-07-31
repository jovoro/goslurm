CC := go

oapi-gen-ro:
	@openapi-generator generate -g go -o v0038 -i api_ro.yaml \
		--git-host github.com --git-user-id pcolladosoto --git-repo-id goslurm/v0038 \
		-c cfg.json

.PHONY: clean

clean:
	rm -f $(TRASH)
