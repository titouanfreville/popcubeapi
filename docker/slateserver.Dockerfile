FROM rails
LABEL MAINTAINER "clement@le-corre.eu"

## ADD ressources to serve slate ;)
ADD docker_resources/slate_resources /home/slate
ADD swagger.yaml /tmp/swagger
ADD docker_resources/swagger /home/swagger
ADD logo.png /home/slate/source/images/logo.png

RUN apt-get -q update && \
 apt-get -q install -yyy --auto-remove --no-install-recommends \
 git curl php5-cli ca-certificates default-jre && \
 rm -rf /var/cache/apt/*

WORKDIR /home/slate
RUN bundle install && bundle exec middleman build --clean

## Set workdir and expose the serveur port
EXPOSE 4567

# GO :)
CMD java -jar /home/swagger/swagger-codegen-cli.jar generate -i /tmp/swagger -l swagger -v -o /home/slate/json/popcubeapi > /dev/null && \
    rm -f /home/slate/popcubeapi.html.md && \
    php /home/swagger/index.php convert /home/slate/json/popcubeapi/swagger.json > /home/slate/source/index.html.md && \
    cd /home/slate && bundle exec middleman server
