#!/usr/bin/env bash
set -x -eou pipefail

# echo "Running tests before cutting the tag..."
# go test -cover
echo "Proceeding with tag creation..."
VERSION=`grep -E "MAJOR|MINOR|PATCH" version.go | cut -d '"' -f 2 | xargs echo -n | tr -s " " "."`
tag_name='v'$VERSION

git tag $tag_name
echo "$tag_name tag is created at current HEAD"
git push origin --tags $tag_name
echo "$tag_name pushed"
if [ $? -ne 0 ]; then
    echo 'An error has occurred! Aborting the script execution...'
    exit 1
fi
