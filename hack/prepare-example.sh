#!/usr/bin/bash
rm -rf ./example/staging/Dope
rsync -av ./ ./example/staging/Dope --exclude example/ --exclude hack/ --exclude .git
