.PHONY: code-gen
code-gen:
	rm -rf redfish_client
	openapi-generator generate -i ./spec/openapi.yaml -g go --package-name "client" -o client/

.PHONY: deps
deps:
	go get github.com/stretchr/testify/assert
	go get golang.org/x/oauth2
	go get golang.org/x/net/context
	go get github.com/antihax/optional
