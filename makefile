MODULES_DIR := modules

# Default target
.PHONY: all
all: help

# Help command
.PHONY: help
help:
	@echo "Usage:"
	@echo "  make module <module_name> - Create a new module directory with a .go file"

# Create a new module
.PHONY: module
module:
	@mkdir -p $(MODULES_DIR)/$(module_name)/sql
	@echo "--sqlc queries file for $(module_name)" > $(MODULES_DIR)/$(module_name)/sql/queries.sql
	@echo "--sqlc schema file for $(module_name)" > $(MODULES_DIR)/$(module_name)/sql/schema.sql
	@echo "package $(module_name)" > $(MODULES_DIR)/$(module_name)/$(module_name).go
	@echo "Module $(module_name) created at $(MODULES_DIR)/$(module_name)/$(module_name).go"

# Allow passing the module name as a variable
module_name := $(word 2, $(MAKECMDGOALS))
