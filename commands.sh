#!/usr/bin/env bash
set -ex

test(){
	go test -v
}

$@