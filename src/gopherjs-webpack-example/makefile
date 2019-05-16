GOPHERJS=gopherjs
GO=GO
NPM=npm
NPX=npx

.PHONY: clean
clean:
	@-rm -f ./libs/*.inc.js

.PHONY: build
build:
	@$(NPX) webpack --config webpack.config
	@$(GO) generate main.go
	@-rm -f ./output/js/*.map

.PHONY: run
run: clean build
