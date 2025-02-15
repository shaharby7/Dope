#!/usr/bin/bash

# prepare and build the example project
## update the "staging" code of "dope" on "example"
rm -rf ./example/staging/Dope
rsync -av ./ ./example/staging/Dope --exclude example/ --exclude hack/ --exclude .git
## remove old build
rm -rf ./example/build
## build "example" project with the new dope build
go run ./cmd/dopecli build -p ./example/.dope -d ./example/build # TODO: check if it can be done with the staged code
