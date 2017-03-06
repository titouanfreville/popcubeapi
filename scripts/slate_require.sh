#!/bin/bash
green="\\033[1;32m"
red="\\033[1;31m"
basic="\\033[0;39m"

COMPOSER="docker run --rm --name composer -v $(pwd)/docker_resources/swagger:/app composer/composer:1.1-alpine install"

sudo rm -rf ./docker_resources/slate_resources
git clone https://github.com/tripit/slate.git ./docker_resources/slate_resources

# cd tmp
# cp -r deploy.sh config.rb font-selection.json Gemfile Gemfile.lock 'source'  ../docker_resources/slate_resources

# cd ..
# rm -rf tmp
git clone https://github.com/E96/swagger2slate.git docker_resources/swagger
$COMPOSER
wget http://repo1.maven.org/maven2/io/swagger/swagger-codegen-cli/2.1.6/swagger-codegen-cli-2.1.6.jar -O docker_resources/swagger/swagger-codegen-cli.jar
