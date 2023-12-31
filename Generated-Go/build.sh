# Generated from trgen 0.21.2

version=`grep 'github.com/antlr4-go/antlr/v4' go.mod | awk '{print $2}' | sed 's/v//'`

export GO111MODULE=on
for i in {1..5}; do go get github.com/antlr4-go/antlr/v4@v$version; if [ "$?" = "0" ]; then break; fi; done; if [ "$?" != "0" ]; then exit 1; fi

set -e

if [ -f transformGrammar.py ]; then python3 transformGrammar.py ; fi

# Because there is no integrated build script for Dart targets, we need
# to manually look at the version in pubspec.yaml and extract the
# version number. We can then use this with antlr4 to generate the
# parser and lexer.
# Go has no version, just the latest version.

antlr4 -v $version -encoding utf-8 -Dlanguage=Go -o parser -lib parser -package parser -visitor parser/Expr.g4


go build Test.go
