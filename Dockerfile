FROM scratch
ARG version=unknown
LABEL name="Tangram" \
      version="$version" \
      description="Tangram, an edge-side html composition service"

ADD dist/tangram /usr/local/tangram/

WORKDIR /usr/local/tangram/

ENTRYPOINT ["/usr/local/tangram/tangram"]
