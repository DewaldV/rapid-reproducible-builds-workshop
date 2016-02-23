#!/bin/sh
set -e

makeEnv() {
    docker build --force-rm -t rrb .
}

build() {
    docker run --rm -ti -v $(pwd)/target:/build/target rrb package
}

clean() {
    rm -rf target/
}

case ${1} in
    'makeEnv')
    makeEnv)
    ;;
    'build')
    build
    ;;
    'clean')
    clean
    ;;
esac
