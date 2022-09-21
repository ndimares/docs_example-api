#!/bin/bash

if [ -f /run/secrets/git_auth_token ]; then
   echo "Found Git secret"
   export TOKEN=$(cat /run/secrets/git_auth_token)
fi

git config --global url."https://${GIT_USERNAME}speakeasybot:${TOKEN}@github.com".insteadOf "https://github.com"
