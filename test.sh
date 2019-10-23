#!/bin/bash

# Run all unit tests.
echo "Running unit tests .."
go test github.com/yashmurty/wealth-park/wpark/core -cover -count=1
go test github.com/yashmurty/wealth-park/wpark/service -cover -count=1
go test github.com/yashmurty/wealth-park/wpark/apiserver -cover -count=1
go test github.com/yashmurty/wealth-park/wpark/controller -cover -count=1

if [[ ! -z "$1" ]]; then
  # Run all e2e tests.
  echo "Running e2e tests .."
  go test github.com/yashmurty/wealth-park/wpark/mysql -cover -count=1

fi

echo "Done."