.PHONY: synced clean

synced:
	@echo "building $@..."
	@go build -v -o $(CURDIR)/bin/synced main.go

clean:
	-rm -rf $(CURDIR)/bin/synced
