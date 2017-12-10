#!/bin/bash
set -euC


function join_list {
    local IFS=","
    echo "$*"
}

case "$1" in
    "standard")
        go test -race $(go list ./... | grep -v /vendor/)
    ;;
    "linter")
        go generate $(go list ./... | grep -v /vendor/) && git diff --exit-code
        gometalinter.v1 --config .linter.json
    ;;
    "coverage")
        echo "" > coverage.txt

        TEST_PKGS=$(go list ./...)

        for d in $TEST_PKGS; do
            go test -i $d
            DEPS=$((go list -f $'{{range $f := .TestImports}}{{$f}}\n{{end}}{{range $f := .Imports}}{{$f}}\n{{end}}' $d && echo $d) | sort -u | grep -v /vendor/ | grep ^github.com/go-kivik/kivikd | tr '\n' ' ')
            go test -coverprofile=profile.out -covermode=set -coverpkg=$(join_list $DEPS) $d
            if [ -f profile.out ]; then
                cat profile.out >> coverage.txt
                rm profile.out
            fi
        done

        bash <(curl -s https://codecov.io/bash)
esac
