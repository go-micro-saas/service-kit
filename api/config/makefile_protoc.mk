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
.PHONY: protoc-config-protobuf
# protoc :-->: generate config api protobuf
protoc-config-protobuf:
	@echo "# generate config api protobuf"
	$(call protoc_protobuf,$(SAAS_CONFIGS_PROTO_FILES))
