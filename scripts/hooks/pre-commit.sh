#!/bin/sh

# Get list of changed files
GO_FILES=$(git diff --cached --name-only | grep ".go$")

# If no files found, all good
if [[ "$GO_FILES" = "" ]]; then
  exit 0
fi

echo "Formatting code..."

# Run through each file and run format
PASS=true
for FILE in $GO_FILES
do
    beforeSum=$(shasum $FILE)
    gofumpt -w $FILE
    afterSum=$(shasum $FILE)

    if [ "$beforeSum" != "$afterSum" ] 
    then
        PASS=false
    fi
done

# If some file has changed, warn user.
if ! $PASS; then
  echo "Failed to commit. Some files have changed!"
  echo "Please review changes and commit them."
  exit 1
fi