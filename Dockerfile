FROM scratch
ARG TARGETOS
ARG TARGETARCH
COPY build/unoserver-rest-api-${TARGETOS}-${TARGETARCH} /unoserver-rest-api
