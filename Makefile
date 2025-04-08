# Makefile

# Variables
MOCKGEN = mockgen
MOCK_OUTPUT_DIR = ./config/mocks
CONTRACT_DIR = ./pkg/user/business/contract
OAUTHMANAGER_DIR = ./pkg/common/oauth
BUSINESS_DIR = ./pkg/user/business

.PHONY: all mocks clean test

all: mocks

mocks:
	@echo "Generating mocks..."
	@mkdir -p $(MOCK_OUTPUT_DIR)/contract
	@for file in $(CONTRACT_DIR)/*.go; do \
		mock_name=$$(basename $$file .go)_mock; \
		$(MOCKGEN) -source=$$file -destination=$(MOCK_OUTPUT_DIR)/contract/$$mock_name.go -package=contract; \
	done
	@mkdir -p $(MOCK_OUTPUT_DIR)/oauthmanager
	@for file in $(OAUTHMANAGER_DIR)/*.go; do \
		mock_name=$$(basename $$file .go)_mock; \
		$(MOCKGEN) -source=$$file -destination=$(MOCK_OUTPUT_DIR)/oauthmanager/$$mock_name.go -package=oauthmanager; \
	done
	@mkdir -p $(MOCK_OUTPUT_DIR)/business
	@for file in $(BUSINESS_DIR)/user.go; do \
		mock_name=$$(basename $$file .go)_mock; \
		$(MOCKGEN) -source=$$file -destination=$(MOCK_OUTPUT_DIR)/business/$$mock_name.go -package=business; \
	done
	@echo "Mocks generated successfully."

clean:
	@echo "Cleaning up generated mocks..."
	rm -rf $(MOCK_OUTPUT_DIR)
	@echo "Clean up completed."

test:
	@go test -cover -coverprofile=coverage.cov $$(go list ./... | grep -v /cache/ | grep -v /repository | grep -v /oauth/ | grep -v /mocks/)
	@go tool cover -func coverage.cov
	@go tool cover -html=coverage.cov -o coverage.html
	@rm -f coverage.cov
