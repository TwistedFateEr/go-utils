#!/usr/bin/env bash

case $1 in
"commit")
  git add .
  git commit -m "$2"
  git pull
  git push
;;
esac