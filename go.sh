#/bin/bash
cd helper
go run $(ls -1 *.go | grep -v _test.go) -- $@
