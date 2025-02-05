# unoserver-rest-api

The simple REST API for unoserver

> **Warning**
>
> It is important to know that the REST API layer DOES NOT provide any type of security whatsoever.  
> It is NOT RECOMMENDED to expose this container image to the internet.

## Usage

Unoserver needs to be installed, see [Installation](https://github.com/unoconv/unoserver#installation) guide.

```sh
NAME:
   unoserver-rest-api - The simple REST API for unoserver and unoconvert

GLOBAL OPTIONS:
   --addr value                The addr used by the unoserver api server (default: "0.0.0.0:80")
   --unoserver-addr value      The unoserver addr used by the unoconvert (default: "127.0.0.1:2002") [$UNOSERVER_ADDR]
   --unoconvert-bin value      Set the unoconvert executable path. (default: "unoconvert") [$UNOCONVERT_BIN]
   --unoconvert-timeout value  Set the unoconvert run timeout (default: 0s) [$UNOCONVERT_TIMEOUT]
   --help, -h                  show help
   --version, -v               print the version
```

### Using with Docker

The [libreofficedocker/libreoffice-unoserver](https://github.com/libreofficedocker/libreoffice-unoserver) already have `unoserver-rest-api` included within the Docker image.

## API

There is only one POST `/request` API.

**Default payload**

```sh
curl -s -v \
   --request POST \
   --url http://127.0.0.1:80/request \
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
   --url http://127.0.0.1:80/request \
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
