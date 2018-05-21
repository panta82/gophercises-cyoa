#!/usr/bin/env bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

cp -f ${DIR}/../assets/gopher.json ${DIR}/../../../bin/
cp -f ${DIR}/../assets/*.gohtml ${DIR}/../../../bin/