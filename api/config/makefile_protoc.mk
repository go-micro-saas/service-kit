# config
SAAS_CONFIGS_API_PROTO=$(shell cd $(PROJECT_PATH) && find api/config -name "*.proto")
#SAAS_CONFIGS_INTERNAL_PROTO=$(shell cd $(PROJECT_PATH) && find app/config/internal/conf -name "*.proto")
SAAS_CONFIGS_INTERNAL_PROTO=
SAAS_CONFIGS_PROTO_FILES=""
ifneq ($(SAAS_CONFIGS_INTERNAL_PROTO), "")
	SAAS_CONFIGS_PROTO_FILES=$(SAAS_CONFIGS_API_PROTO) $(SAAS_CONFIGS_INTERNAL_PROTO)
else
	SAAS_CONFIGS_PROTO_FILES=$(SAAS_CONFIGS_API_PROTO)
endif
.PHONY: protoc-configs
# protoc :-->: generate configs protobuf
protoc-configs:
	@echo "# generate configs protobuf"
	if [ "$(SAAS_CONFIGS_PROTO_FILES)" != "" ]; then \
		cd $(PROJECT_PATH); \
		protoc \
			--proto_path=. \
			--proto_path=$(GOPATH)/src \
			--proto_path=./third_party \
			--go_out=paths=source_relative:. \
			--go-grpc_out=paths=source_relative:. \
			--go-http_out=paths=source_relative:. \
			--go-errors_out=paths=source_relative:. \
			--validate_out=paths=source_relative,lang=go:. \
			--openapiv2_out . \
			--openapiv2_opt logtostderr=true \
			--openapiv2_opt allow_delete_body=true \
			--openapiv2_opt json_names_for_fields=false \
			--openapiv2_opt enums_as_ints=true \
			--openapi_out=fq_schema_naming=true,enum_type=integer,default_response=true:. \
			$(SAAS_CONFIGS_PROTO_FILES) ; \
	fi
