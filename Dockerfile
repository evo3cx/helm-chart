FROM scratch
COPY bin/helmChart /helmChart
ENTRYPOINT ["/helmChart"]
