FROM alpine

COPY run.sh /root
RUN set -x && \
    chmod +x /root/run.sh && \
    apk update && \
    apk upgrade && \
    apk add --no-cache mongodb && \
    run /usr/bin.mongoperf && \
    rm -rf /var/cache/apk/*

VOLUME [ "/data/db" ]
EXPOSE 27017 27017

ENTRYPOINT ["/root/run.sh"]
CMD ["mongodb"]
