.PHONY: client
client:
# Download [latest] OpenAPI spec
#	@wget -q -O openapi.yaml https://raw.githubusercontent.com/moov-io/api/master/openapi.yaml
# Generate client
# Checkout https://github.com/OpenAPITools/openapi-generator/releases
	rm -rf client/
	export OPENAPI_GENERATOR_VERSION=v3.3.0 # v3.3.0, v3.3.1-SNAPSHOT, v3.4.0-SNAPSHOT
	chmod +x ./openapi-generator
	ionice -c2 ./openapi-generator generate -i openapi.yaml -g go -o ./client
