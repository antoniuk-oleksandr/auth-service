#!/usr/bin/env bash
set -e

TEMPL_DIR="./internal/view"
HASH_FILE="./tmp/.templ_hash"

# Calculate current hash of all .templ files
CURRENT_HASH=$(find "$TEMPL_DIR" -type f -name "*.templ" -exec sha1sum {} \; | sha1sum | awk '{print $1}')

# Compare with the previous hash
if [ -f "$HASH_FILE" ]; then
    PREV_HASH=$(cat "$HASH_FILE")
else
    PREV_HASH=""
fi

# Only generate if hash differs
if [ "$CURRENT_HASH" != "$PREV_HASH" ]; then
    echo "🔄 Changes detected in templ files, regenerating..."
    templ generate "$TEMPL_DIR"/...
    echo "$CURRENT_HASH" > "$HASH_FILE"
else
    echo "✅ No changes in templ files."
fi
