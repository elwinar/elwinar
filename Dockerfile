FROM scratch
MAINTAINER Romain Baugue <romain.baugue@elwinar.com>

ADD elwinar /app
ADD public /public
ADD database.sqlite /database.sqlite

ENV APP_PORT 80
ENV APP_DATABASE /database.sqlite
ENV APP_PUBLIC /public/

ENTRYPOINT ["/app"]

