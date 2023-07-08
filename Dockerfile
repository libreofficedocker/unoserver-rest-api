ARG ALPINE_VERSION=edge
FROM libreofficedocker/libreoffice-unoserver:${ALPINE_VERSION}

COPY bin/unoserver-rest-api-linux /usr/bin/unoserver-rest-api
ADD rootfs /

EXPOSE 2003
