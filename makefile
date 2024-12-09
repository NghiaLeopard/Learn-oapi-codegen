codegen-api:
	oapi-codegen \
	-generate fiber,strict-server,types,spec \
	-package api -o api/http.gen.go openapi/api.yaml