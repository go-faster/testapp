ARG IMG=gcr.io/distroless/static-debian11
FROM $IMG:nonroot

COPY testapp /usr/bin/local/testapp

ENTRYPOINT ["/usr/bin/local/testapp"]
