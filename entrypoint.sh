#!/bin/sh

if [ -z "$VALUE_TO_SET" ] || [ -z "$RPC_API_KEY" ] || [ -z "$WALLET_PRIVATE_KEY" ]; then
  echo "One or more required environment variables are missing: VALUE_TO_SET, RPC_API_KEY, WALLET_PRIVATE_KEY"
  exit 1
fi

/app/deploy -value "$VALUE_TO_SET" -rpc "$RPC_API_KEY" -private "$WALLET_PRIVATE_KEY"
