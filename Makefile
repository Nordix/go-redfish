.PHONY: code-gen
code-gen:
	rm -rf client
	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate \
         -i  /local/spec/openapi.yaml \
         -g go \
	 --package-name "client" \
        --git-repo-id go-redfish/client \
        --git-user-id Nordix \
	 -p enumClassPrefix=true \
         -o /local/client
	#openapi-generator generate -i ./spec/openapi.yaml -g go --package-name "client" --git-repo-id go-redfish/client --git-user-id Nordix -o client/ -p enumClassPrefix=true
	mkdir -p api
	go run api_generator.go | gofmt > api/service_interface.go
	go generate api/service_interface.go


.PHONY: deps
deps:
	go get github.com/stretchr/testify/assert
	go get golang.org/x/oauth2
	go get golang.org/x/net/context
	go get github.com/antihax/optional
