FROM redis:7.2-alpine

USER root

WORKDIR /redis

# COPY redis.conf /usr/local/etc/redis/redis.conf

COPY entry.sh ./

RUN chmod +x ./entry.sh

ENTRYPOINT ["sh", "-c", "./entry.sh"]
