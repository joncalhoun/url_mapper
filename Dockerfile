FROM alpine:latest

MAINTAINER Jon Calhoun <jon@calhoun.io>

WORKDIR "/opt"

ADD .docker_build/url_mapper /opt/bin/url_mapper
ADD ./templates /opt/templates
ADD ./static /opt/static

CMD ["/opt/bin/url_mapper"]
