# unoserver-rest-api

The simple REST API for unoserver

## Usage

Unoserver needs to be installed, see [Installation](https://github.com/unoconv/unoserver#installation) guide.

```sh
NAME:
   unoserver-rest-api - The simple REST API for unoserver and unoconvert

GLOBAL OPTIONS:
   --addr value             The addr used by the unoserver api server (default: "0.0.0.0:2003")
   --unoconvert-addr value  The addr used by the unoconvert (default: "127.0.0.1:2002")
   --unoconvert-bin value   Set the unoconvert executable path. (default: "unoconvert") [$UNOCONVERT_EXECUTABLE_PATH]
   --help, -h               show help
   --version, -v            print the version
```

## API

There is only one POST `/request` API.

**Default payload**

```sh
curl -s -v \
   --request POST \
   --url http://127.0.0.1:2004/request \
   --header 'Content-Type: multipart/form-data' \
   --form "file=@/paht/to/your/file.xlsx" \
   --form 'convert-to=pdf'
```

- `file`: Type of `File`, required
- `convert-to`: Typeof `String`, required

**Advance payload**

```sh
curl -s -v \
   --request POST \
   --url http://127.0.0.1:2004/request \
   --header 'Content-Type: multipart/form-data' \
   --form "file=@/paht/to/your/file.xlsx" \
   --form 'convert-to=pdf' \
   --form 'opts[]=--landscape'
```

- `file`: Type of `File`, required
- `convert-to`: Typeof `String`, required
- `opts`: Type of `String[]`

## License

Licensed under [Apache-2.0 license](LICENSE).
