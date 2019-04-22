REPLACE?=-replace flamingo.me/flamingo/v3=../flamingo -replace flamingo.me/flamingo-commerce/v3=../flamingo-commerce -replace flamingo.me/flamingo-commerce-adapter-standalone=../flamingo-commerce-adapter-standalone
DROPREPLACE?=-dropreplace flamingo.me/flamingo/v3 -dropreplace flamingo.me/flamingo-commerce/v3 -dropreplace flamingo.me/flamingo-commerce-adapter-standalone

.PHONY: local unlocal

local:
	git config filter.gomod-flamingo-commerce-magento2adapter.smudge 'go mod edit -fmt -print $(REPLACE) /dev/stdin'
	git config filter.gomod-flamingo-commerce-magento2adapter.clean 'go mod edit -fmt -print $(DROPREPLACE) /dev/stdin'
	git config filter.gomod-flamingo-commerce-magento2adapter.required true
	go mod edit -fmt $(REPLACE)

unlocal:
	git config filter.gomod-flamingo-commerce-magento2adapter.smudge ''
	git config filter.gomod-flamingo-commerce-magento2adapter.clean ''
	git config filter.gomod-flamingo-commerce-magento2adapter.required false
	go mod edit -fmt $(DROPREPLACE)

test:
	go test -race -v ./...
	