#!/usr/bin/bash
git push --delete origin $1
git tag --delete $1
