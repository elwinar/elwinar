FROM scratch
MAINTAINER Romain Baugue <romain.baugue@elwinar.com>

ADD elwinar /app
ADD public /public

ENV APP_DATABASE /database.sqlite
ENV APP_DEBUG false
ENV APP_PORT 80
ENV APP_PUBLIC /public

EXPOSE 80

ENTRYPOINT ["/app"]

