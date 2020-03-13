FROM scratch

COPY ncov-stats /

ENTRYPOINT ["/ncov-stats"]
