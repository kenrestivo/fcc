default: help

DUMPFILE := artifacts/fcc-license-view-data-csv-format.zip
BOLTDB := artifacts/fcc.db
FCC2BOLT := bin/fcc2bolt

.PHONY: help
help:
	@echo
	@echo "ingest: ingest FCC database into local database format"
	@echo

.PHONY: ingest
ingest: $(BOLTDB)

.PHONY: clean
clean:
	rm -rf artifacts/*
	rm -rf bin/*

$(DUMPFILE):
	cd artifacts && wget http://data.fcc.gov/download/license-view/fcc-license-view-data-csv-format.zip

$(BOLTDB): $(DUMPFILE) | $(FCC2BOLT)
	$(FCC2BOLT) -dump $(DUMPFILE) -db $(BOLTDB)

$(FCC2BOLT):
	go build -o $@ cmd/fcc2bolt/main.go

