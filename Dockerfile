FROM scratch
MAINTAINER Romain Baugue <romain.baugue@elwinar.com>

ADD build /app/

ENV APP_DATABASE /app/database.sqlite
ENV APP_DEBUG false
ENV APP_PORT 80
ENV APP_PUBLIC /app/public

EXPOSE 80

ENTRYPOINT ["/app/app"]

