# saas services
SAAS_SERVICE_PROTO_FILES=$(shell cd $(PROJECT_PATH) && find api -name "*.proto")
.PHONY: protoc-service-protobuf
# protoc :-->: generate service protobuf
protoc-service-protobuf:
	@echo "# generate service protobuf"
	$(call protoc_protobuf,$(SAAS_SERVICE_PROTO_FILES))

# specified server
SAAS_SERVICE_SPECIFIED_FILES=$(shell cd $(PROJECT_PATH) && find ./api/${service} -name "*.proto")
.PHONY: protoc-specified-protobuf
# protoc :-->: example: make protoc-specified-protobuf service=ping-service
protoc-specified-protobuf:
	@echo "# generate ${service} protobuf"
	$(call protoc_protobuf,$(SAAS_SERVICE_SPECIFIED_FILES))
