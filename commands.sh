#!/usr/bin/env bash
set -ex

test(){
	go test -c
	./injection.test -alsologtostderr
}

$@