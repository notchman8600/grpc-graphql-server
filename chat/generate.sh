#!/bin/bash
go run github.com/99designs/gqlgen generate
# VSCode環境下だとmoduleの諸々でgo.modが書き換わるので暫定処置
go get -u github.com/99designs/gqlgen
