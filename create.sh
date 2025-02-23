#!/bin/bash

if [ -z "$1" ]; then
    echo "Usage: $0 <folder_name>"
    exit 1
fi

FOLDER_NAME=$1

mkdir -p "$FOLDER_NAME"
echo "package solution" > "$FOLDER_NAME/solution.go"