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
        diff -u <(echo -n) <(gofmt -e -d $(find . -type f -name '*.go' -not -path "./vendor/*"))
        go install # to make gotype (run by gometalinter) happy
        gometalinter.v1 --config .linter_test.json
        gometalinter.v1 --config .linter.json
    ;;
    "coverage")
        echo "" > coverage.txt

        TEST_PKGS=$(find -name "*_test.go" | grep -v /vendor/ | xargs dirname | sort -u | sed -e "s#^\.#github.com/go-kivik/couchdb#" )

        for d in $TEST_PKGS; do
            go test -i $d
            DEPS=$((go list -f $'{{range $f := .TestImports}}{{$f}}\n{{end}}{{range $f := .Imports}}{{$f}}\n{{end}}' $d && echo $d) | sort -u | grep -v /vendor/ | grep -v /pouchdb | grep -v /kivik/test | grep ^github.com/flimzy/kivik | tr '\n' ' ')
            go test -coverprofile=profile.out -covermode=set -coverpkg=$(join_list $DEPS) $d
            if [ -f profile.out ]; then
                cat profile.out >> coverage.txt
                rm profile.out
            fi
        done

        bash <(curl -s https://codecov.io/bash)
esac
