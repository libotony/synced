.PHONY: synced clean

synced:
	@echo "building $@..."
	@go build -o $(CURDIR)/bin/synced main.go
	@echo "done. executable created at 'bin/$@'"

clean:
	rm -rf $(CURDIR)/bin/synced
