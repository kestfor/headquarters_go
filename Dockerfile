FROM alpine:3.9 as builder



RUN apk update && apk add ca-certificates tzdata



FROM scratch



COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo/

ADD data ./data
COPY bin/main /
CMD ["chmod", "777", "main"]

CMD ["/main"]