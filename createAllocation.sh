#!/bin/bash

# Script to automate blobber token locking and allocation creation
# Usage: ./createAllocation.sh [--zbox-path PATH] [--tokens AMOUNT] [--lock AMOUNT]

set -e

# Default values
ZBOX_CMD="./zbox"
TOKENS_PER_BLOBBER=0.5
LOCK_AMOUNT=0.5
MAX_RETRIES=3
RETRY_DELAY=5

# Parse command line arguments
while [[ $# -gt 0 ]]; do
    case $1 in
        --zbox-path)
            ZBOX_CMD="$2"
            shift 2
            ;;
        --tokens)
            TOKENS_PER_BLOBBER="$2"
            shift 2
            ;;
        --lock)
            LOCK_AMOUNT="$2"
            shift 2
            ;;
        *)
            echo "Unknown option: $1"
            echo "Usage: $0 [--zbox-path PATH] [--tokens AMOUNT] [--lock AMOUNT]"
            exit 1
            ;;
    esac
done

# Function to calculate data and parity based on number of blobbers
calculate_data_parity() {
    local num_blobbers=$1
    local data=0
    local parity=0
    
    case $num_blobbers in
        2)
            data=1
            parity=1
            ;;
        3)
            data=2
            parity=1
            ;;
        4)
            data=3
            parity=1
            ;;
        5)
            data=3
            parity=2
            ;;
        6)
            data=4
            parity=2
            ;;
        *)
            echo "Error: Unsupported number of blobbers: $num_blobbers"
            echo "Supported: 2, 3, 4, 5, or 6 blobbers"
            exit 1
            ;;
    esac
    
    echo "$data $parity"
}

# Function to extract blobber IDs from zbox ls-blobbers output
extract_blobber_ids() {
    local output="$1"
    # Match lines like "- id: <blobber_id>" or "  id: <blobber_id>"
    # Extract the blobber ID (last field after "id:")
    echo "$output" | grep -E "^\s*-?\s*id:\s+" | sed 's/.*id:\s*//' | awk '{print $1}'
}

# Function to lock tokens for a blobber with retry
lock_tokens() {
    local blobber_id=$1
    local retry_count=0
    
    while [ $retry_count -lt $MAX_RETRIES ]; do
        echo "Locking $TOKENS_PER_BLOBBER tokens for blobber $blobber_id (attempt $((retry_count + 1))/$MAX_RETRIES)..."
        
        local result
        result=$($ZBOX_CMD sp-lock --blobber_id "$blobber_id" --tokens "$TOKENS_PER_BLOBBER" 2>&1)
        
        if echo "$result" | grep -q "tokens locked"; then
            echo "✓ Successfully locked tokens for blobber $blobber_id"
            return 0
        else
            retry_count=$((retry_count + 1))
            if [ $retry_count -lt $MAX_RETRIES ]; then
                echo "✗ Failed to lock tokens. Retrying in $RETRY_DELAY seconds..."
                # Show error if it's not a network/sharder confirmation issue
                if echo "$result" | grep -q "too less sharders"; then
                    echo "  (Network/sharder confirmation issue - will retry)"
                fi
                sleep $RETRY_DELAY
            else
                echo "✗ Failed to lock tokens for blobber $blobber_id after $MAX_RETRIES attempts"
                echo "Last error output:"
                echo "$result" | tail -5
                return 1
            fi
        fi
    done
}

# Function to create allocation with retry
create_allocation() {
    local blobber_ids=$1
    local data=$2
    local parity=$3
    local retry_count=0
    
    while [ $retry_count -lt $MAX_RETRIES ]; do
        echo "Creating allocation with $data data and $parity parity blobbers (attempt $((retry_count + 1))/$MAX_RETRIES)..."
        
        local result
        result=$($ZBOX_CMD newallocation --lock "$LOCK_AMOUNT" --data "$data" --parity "$parity" --preferred_blobbers "$blobber_ids" --force 2>&1)
        
        if echo "$result" | grep -q "Allocation created:"; then
            local allocation_id
            allocation_id=$(echo "$result" | grep "Allocation created:" | awk '{print $3}')
            echo "✓ Successfully created allocation: $allocation_id"
            echo "$allocation_id"
            return 0
        else
            retry_count=$((retry_count + 1))
            if [ $retry_count -lt $MAX_RETRIES ]; then
                echo "✗ Failed to create allocation. Retrying in $RETRY_DELAY seconds..."
                # Show error if it's not a network/sharder confirmation issue
                if echo "$result" | grep -q "too less sharders"; then
                    echo "  (Network/sharder confirmation issue - will retry)"
                fi
                sleep $RETRY_DELAY
            else
                echo "✗ Failed to create allocation after $MAX_RETRIES attempts"
                echo "Last error output:"
                echo "$result" | tail -10
                return 1
            fi
        fi
    done
}

# Main execution
echo "========================================="
echo "Blobber Token Locking and Allocation Script"
echo "========================================="
echo ""

# Step 1: List blobbers
echo "Step 1: Listing available blobbers..."
BLOBBER_LIST_OUTPUT=$($ZBOX_CMD ls-blobbers 2>&1)

if [ $? -ne 0 ]; then
    echo "Error: Failed to list blobbers. Make sure zbox is installed and configured."
    exit 1
fi

# Extract blobber IDs
BLOBBER_IDS=($(extract_blobber_ids "$BLOBBER_LIST_OUTPUT"))
NUM_BLOBBERS=${#BLOBBER_IDS[@]}

if [ $NUM_BLOBBERS -eq 0 ]; then
    echo "Error: No blobbers found. Make sure blobbers are running and registered."
    exit 1
fi

echo "Found $NUM_BLOBBERS blobber(s):"
for i in "${!BLOBBER_IDS[@]}"; do
    echo "  $((i+1)). ${BLOBBER_IDS[$i]}"
done
echo ""

# Calculate data and parity
read -r DATA PARITY <<< "$(calculate_data_parity $NUM_BLOBBERS)"

echo "Configuration:"
echo "  Data blobbers: $DATA"
echo "  Parity blobbers: $PARITY"
echo "  Tokens per blobber: $TOKENS_PER_BLOBBER"
echo "  Lock amount: $LOCK_AMOUNT"
echo ""

# Step 2: Lock tokens for each blobber
echo "Step 2: Locking tokens for each blobber..."
FAILED_LOCKS=0
for blobber_id in "${BLOBBER_IDS[@]}"; do
    if ! lock_tokens "$blobber_id"; then
        FAILED_LOCKS=$((FAILED_LOCKS + 1))
    fi
    echo ""
done

if [ $FAILED_LOCKS -gt 0 ]; then
    echo "Warning: Failed to lock tokens for $FAILED_LOCKS blobber(s). Continuing anyway..."
    echo ""
fi

# Step 3: Create allocation
echo "Step 3: Creating allocation..."
BLOBBER_IDS_CSV=$(IFS=,; echo "${BLOBBER_IDS[*]}")

ALLOCATION_ID=$(create_allocation "$BLOBBER_IDS_CSV" "$DATA" "$PARITY")

if [ $? -eq 0 ] && [ -n "$ALLOCATION_ID" ]; then
    echo ""
    echo "========================================="
    echo "✓ Success! Allocation created:"
    echo "  Allocation ID: $ALLOCATION_ID"
    echo "  Data blobbers: $DATA"
    echo "  Parity blobbers: $PARITY"
    echo "  Total blobbers: $NUM_BLOBBERS"
    echo "========================================="
    exit 0
else
    echo ""
    echo "========================================="
    echo "✗ Failed to create allocation"
    echo "========================================="
    exit 1
fi

# #!/bin/bash
# set -euo pipefail

# BLOBBER_URL="http://127.0.0.1:5051"
# WALLET_JSON="/var/0chain/blobber/blob_op_wallet.json"
# ROUND_EXPIRY=600   # seconds before the ticket expires
# JQ_BIN="${JQ_BIN:-jq}"

# usage() {
#   echo "Usage: $0 [--blobber-url URL] [--wallet PATH] [--round SECONDS]"
#   exit 1
# }

# while [[ $# -gt 0 ]]; do
#   case "$1" in
#     --blobber-url) BLOBBER_URL="$2"; shift 2 ;;
#     --wallet)      WALLET_JSON="$2"; shift 2 ;;
#     --round)       ROUND_EXPIRY="$2"; shift 2 ;;
#     *) usage ;;
#   esac
# done

# [[ -x "$(command -v "$JQ_BIN")" ]] || { echo "jq is required"; exit 1; }
# [[ -f "$WALLET_JSON" ]] || { echo "Wallet file not found: $WALLET_JSON"; exit 1; }

# PUBLIC_KEY=$("$JQ_BIN" -r '.keys[0].public_key' "$WALLET_JSON")
# PRIVATE_KEY=$("$JQ_BIN" -r '.keys[0].private_key' "$WALLET_JSON")
# CLIENT_ID=$("$JQ_BIN" -r '.client_id' "$WALLET_JSON")

# read -r -d '' PAYLOAD <<JSON
# {
#   "public_key": "$PUBLIC_KEY",
#   "private_key": "$PRIVATE_KEY",
#   "blobber_url": "$BLOBBER_URL",
#   "client_id": "$CLIENT_ID",
#   "round_expiry": $ROUND_EXPIRY
# }
# JSON

# echo "Requesting auth ticket from $BLOBBER_URL ..."
# curl -sS -X POST "$BLOBBER_URL/v1/auth/getAuthTicketWithClientId" \
#   -H 'Content-Type: application/json' \
#   -d "$PAYLOAD" | tee /tmp/auth-ticket.json

