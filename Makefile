.PHONY: check install-hooks git-skipchecks push-skipchecks

install-hooks:
	go install mvdan.cc/gofumpt@v0.1.1
	go install honnef.co/go/tools/cmd/staticcheck@v0.2.0
	go install github.com/securego/gosec/v2/cmd/gosec@v2.8.1
	cp scripts/hooks/pre-commit.sh .git/hooks/pre-commit
	cp scripts/hooks/pre-push.sh .git/hooks/pre-push
check: fmt test static vet sec

run: 
	go run cmd/http/main.go

test:
	go test ./...

fmt:
	gofumpt -w -l .

static:
	staticcheck -checks all ./...

vet:
	go vet ./...
	go vet -vettool=$(which shadow) ./...

sec:
	gosec ./...

skipcheck-precommit:
	git commit --no-verify

skipcheck-prepush:
	git push --no-verify