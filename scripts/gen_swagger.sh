#!/bin/bash
if [[ $? -gt 0 ]]
then
  echo "moving to $1"
  cd $1
fi
echo "Generating json spec"
swagger generate spec -m -o swagger.json || exit 1
echo "Generating yaml spec"
java -jar swagger-codegen-cli.jar generate -i swagger.json  -l swagger-yaml || exit 2
echo "Removing swagger.json"
# rm swagger.json || return 3
echo "Correctly generated doc"