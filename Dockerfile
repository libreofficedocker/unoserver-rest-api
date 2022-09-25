FROM socheatsok78/libreoffice-unoserver:nightly

COPY bin/unoserver-rest-api-linux /usr/bin/unoserver-rest-api
ADD rootfs /
