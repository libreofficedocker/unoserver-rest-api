# unoserver-rest-api

The simple REST API for unoserver

## Usage

Unoserver needs to be installed, see [Installation](https://github.com/unoconv/unoserver#installation) guide.

```sh
NAME:
   unoserver-rest-api - The simple REST API for unoserver and unoconvert

GLOBAL OPTIONS:
   --addr value                The addr used by the unoserver api server (default: "0.0.0.0:2004")
   --unoserver-addr value      The unoserver addr used by the unoconvert (default: "127.0.0.1:2002") [$UNOSERVER_ADDR]
   --unoconvert-bin value      Set the unoconvert executable path. (default: "unoconvert") [$UNOCONVERT_EXECUTABLE_PATH]
   --unoconvert-timeout value  Set the unoconvert run timeout (default: 0s) [$UNOCONVERT_TIMEOUT]
   --help, -h                  show help
   --version, -v               print the version
```

### Using with Docker

You can use `unoserver-rest-api` with [libreoffice-docker/libreoffice-unoserver-alpine](https://github.com/libreoffice-docker/libreoffice-unoserver-alpine) or [libreoffice-docker/libreoffice-unoserver-ubuntu](https://github.com/libreoffice-docker/libreoffice-unoserver-ubuntu) by modifying the `Dockerfile` provided by the template.

```Dockerfile
# ...

# The unoserver-rest-api version number
ARG UNOSERVER_REST_API_VERSION=0.4.0
ADD https://github.com/libreoffice-docker/unoserver-rest-api/releases/download/v${UNOSERVER_REST_API_VERSION}/s6-overlay-module.tar.zx /tmp
ADD https://github.com/libreoffice-docker/unoserver-rest-api/releases/download/v${UNOSERVER_REST_API_VERSION}/s6-overlay-module.tar.zx.sha256 /tmp
RUN cd /tmp && sha256sum -c *.sha256 && \
    tar -C / -Jxpf /tmp/s6-overlay-module.tar.zx && \
    rm -rf /tmp/*.tar*
EXPOSE 2004
```

## API

There is only one POST `/request` API.

**Default payload**

```sh
curl -s -v \
   --request POST \
   --url http://127.0.0.1:2004/request \
   --header 'Content-Type: multipart/form-data' \
   --form "file=@/path/to/your/file.xlsx" \
   --form 'convert-to=pdf' \
   --output 'file.pdf'
```

- `file`: Type of `File`, required
- `convert-to`: Type of `String`, required

**Advance payload**

```sh
curl -s -v \
   --request POST \
   --url http://127.0.0.1:2004/request \
   --header 'Content-Type: multipart/form-data' \
   --form "file=@/path/to/your/file.xlsx" \
   --form 'convert-to=pdf' \
   --form 'opts[]=--landscape' \
   --output 'file.pdf'
```

- `file`: Type of `File`, required
- `convert-to`: Type of `String`, required
- `opts`: Type of `String[]`

## License

Licensed under [Apache-2.0 license](LICENSE).
