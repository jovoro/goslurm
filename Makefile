CC := go

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
