#!/bin/bash

echo "download Unicode database ..."
curl --request GET -sL --url 'https://www.unicode.org/Public/UCD/latest/ucdxml/ucd.all.flat.zip' --output './testdata/ucd.all.flat.zip'
unzip -o testdata/ucd.all.flat.zip -d testdata/
echo "done"
