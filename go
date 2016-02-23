#!/bin/sh
set -e

makeEnv() {
    sudo docker build --force-rm -t rrb .
}

build() {
    sudo docker run --rm --link nexus rrb deploy
}

case ${1} in
    'makeEnv')
    makeEnv
    ;;
    'build')
    build
    ;;
esac
