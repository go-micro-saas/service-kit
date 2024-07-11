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
.PHONY: protoc-api-config
# protoc :-->: generate configs protobuf
protoc-api-config:
	@echo "# generate ${service} protobuf"
	$(call protoc_protobuf,$(SAAS_CONFIGS_PROTO_FILES))
