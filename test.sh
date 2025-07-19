#!/bin/sh

# Set your API credentials as environment variables
export MEXC_API_KEY=""
export MEXC_API_SECRET=""

# Run just the integration tests
go test -v -run TestQueryCurrencyInfo_RealAPI

# Or run all tests (unit + integration)
go test -v  