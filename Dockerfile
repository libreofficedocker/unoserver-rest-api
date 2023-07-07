FROM libreoffice-docker/libreoffice-unoserver:test

COPY bin/unoserver-rest-api-linux /usr/bin/unoserver-rest-api
ADD rootfs /

EXPOSE 2003
