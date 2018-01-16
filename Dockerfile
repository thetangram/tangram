FROM scratch
ARG version=unknown
LABEL name="Tangram" \
      version="$version" \
      description="Tangram, an edge-side html composition service"

ADD dist/tangramd /usr/local/tangramd/

WORKDIR /usr/local/tangramd/

ENTRYPOINT ["/usr/local/tangramd/tangramd"]
