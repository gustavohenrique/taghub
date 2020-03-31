#!/bin/bash

rm mocks/* 2>/dev/null
for i in `/bin/ls -1 src/domain/interfaces.go`; do
    out=`basename $i`
    mockgen -source ${i} -destination mocks/mock_${out} -package mocks
done
mockgen -source libs/mundipagg/interfaces.go -destination mocks/mock_mundipagg.go -package mocks
