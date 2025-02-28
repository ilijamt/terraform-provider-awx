#!/bin/bash
DIRNAME=$(dirname -- "$0")
TARGETDIR=$(realpath "$DIRNAME/../examples")
set -x
find $TARGETDIR  -type f -name "terraform.tfstate*" -delete
find $TARGETDIR -type f -name "plan" -delete