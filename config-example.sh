#!/bin/bash

set -a

# API server id, host and port.
WPARK_SERVER_ID="mercari-server-1"

# Mercari API. This is where all external clients connect.
# Expose this port to the internet / outside the cluster's
# firewall / network security group.
WPARK_CORE_API_HOST="localhost"
WPARK_CORE_API_PORT="11111"

# MYSQL credentials.
WPARK_MYSQL_URL="root:password@tcp(localhost:13306)"
WPARK_MYSQL_DB_NAME="wpark"

# Passing the string "true" puts the billing into production mode.
WPARK_PRODUCTION="false"

set +a
