#!/bin/sh
set -e

makeEnv() {
    sudo docker build --force-rm -t rrb .
}

build() {
    sudo docker run --rm -v $(pwd)/target:/build/target rrb package
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
