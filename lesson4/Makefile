.PHONY: groupimports
groupimports:
	find . -type f -name '*.go' ! -path './.git/*' | xargs -L 1 python3 third_party/groupimports.py

.PHONY: lint
lint:
	 golangci-lint run -c ./configs/golangci.yml -v ./...