#!/bin/bash

go get github.com/go-delve/delve/cmd/dlv
dlv debug 16.9.go
