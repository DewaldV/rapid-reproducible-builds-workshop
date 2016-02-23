#!/bin/sh
set -e

makeEnv() {
    sudo docker build --force-rm -t rrb .
}

build() {
    target_dir="$(pwd)/target"
    mkdir -p ${target_dir}
    sudo docker run --rm --link nexus -v ${target_dir}:/build/target rrb package
    sudo chown -R go:go target/
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
