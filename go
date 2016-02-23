#!/bin/sh
set -e

makeEnv() {
    sudo docker build --force-rm -t rrb .
}

build() {
    sudo docker run --rm --link nexus rrb deploy
}

clean() {
    sudo rm -rf target/
}

case ${1} in
    'makeEnv')
    makeEnv
    ;;
    'build')
    build
    ;;
    'clean')
    clean
    ;;
esac
