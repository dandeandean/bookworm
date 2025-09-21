#!/bin/bash

echo "Testing install"

make install

if ! command -v bookworm 2>/dev/null ; then
	echo "failed to install"
	exit 1
fi

echo "===================="

echo "Testing pre-init"

bookworm

if [[ $? != 2 ]]; then
	exit 1
fi

echo "===================="

echo "Testing init"

if ! bookworm init; then
	echo "failed to init"
	exit 1
fi

echo "===================="
echo "Testing create new bookmark"

if ! bookworm new google www.google.com ; then
	echo "failed to create bookmark"
	exit 1
fi

if ! bookworm ls | grep -q "google"; then
	echo "Couldn't find google in the bookmarks"
	exit 1
fi


