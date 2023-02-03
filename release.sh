#!/bin/bash

NEW_VERSION=$(git tag -l --sort=-v:refname | head -n 1 | awk -F . 'BEGIN{OFS=""}{print $1,".",$2,".",$3+1}')
git tag ${NEW_VERSION}
git push origin ${NEW_VERSION}
