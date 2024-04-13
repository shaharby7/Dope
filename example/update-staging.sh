#!/usr/bin/bash
rm -rf ./staging/Dope
mkdir ./staging/Dope
rsync -av ../ ./staging/Dope --exclude example