.PHONY: synced clean

synced:
	@echo "building $@..."
	@go build -o $(CURDIR)/bin/synced main.go

clean:
	rm -rf $(CURDIR)/bin/synced
