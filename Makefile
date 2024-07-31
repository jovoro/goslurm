CC := go

oapi-gen-ro:
	@openapi-generator generate -g go -o v0038-f -i api_ro.yaml \
		--git-host github.com --git-user-id pcolladosoto --git-repo-id goslurm/v0038 \
		--additional-properties=packageName=v0038,withGoMod=false

.PHONY: clean

clean:
	rm -f $(TRASH)
