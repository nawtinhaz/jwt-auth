#!/bin/sh

#  Check if dependencies have changed and modules are being used
echo ""
echo "Checking for unresolved dependency changes..."

if [ ! -e go.mod ]
then 
    echo "Not using go modules... Please create go.mod to continue"
    exit 1
fi

goModDigest=$(shasum go.mod)
if [ -e go.sum ]; then goSumDigest=$(shasum go.sum); fi

go mod tidy
if [ -e go.sum ]
then
    go mod vendor
    goSumDigest=$(shasum go.sum)
fi

goModNewDigest=$(shasum go.mod)
if [ -e go.sum ]; then goSumNewDigest=$(shasum go.sum); fi

if [ "$goModDigest" != "$goModNewDigest" ]
then
    echo "Failed to commit. go.mod has changed..."
    exit 1
fi

if [ "$goSumDigest" != "$goSumNewDigest" ]
then
    echo "Failed to commit. go.sum has changed..."
    exit 1
fi

# Static check
echo ""
echo "Static checks..."

staticcheck -checks all ./...

if [ "$?" != "0" ]
then
    echo "" 
    printf "Failed to commit. Static check failed!\n"
    exit 1
fi

# Go vet
echo ""
echo "Go vet checks..."

go vet ./... && go vet -vettool=$(which shadow) ./...

if [ "$?" != "0" ]
then
    echo "" 
    printf "Failed to commit. Go vet failed!\n"
    exit 1
fi

# Go sec
echo ""
echo "Go sec checks..."

gosec ./...

if [ "$?" != "0" ]
then
    echo "" 
    printf "Failed to commit. Go sec failed!\n"
    exit 1
fi