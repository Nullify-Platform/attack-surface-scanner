<a href="https://nullify.ai">
  <img src="https://uploads-ssl.webflow.com/6492db86d53f84f396b6623d/64dad6c12b98dee05eb08088_nullify%20logo.png" alt="Nullify" width="300"/>
</a>

# Attack Surface Scanning CLI (PoC)

## Install

Download the relevant binary from the
[latest release](https://github.com/nullify-platform/attack-surface-scanner/releases/latest)
and add it to your path.

### Linux

```sh
curl -L -o as https://github.com/nullify-platform/attack-surface-scanner/releases/latest/download/as_linux_amd64
chmod +x as
sudo mv as /usr/local/bin
```

## Usage

```
Usage: as [--verbose] [--debug] <command> [<args>]

Options:
  --verbose, -v          enable verbose logging
  --debug, -d            enable debug logging
  --help, -h             display this help and exit
  --version              display version and exit

Commands:
  scan                   test the given app for vulnerabilities
```

## Running a REST API scan

```
Usage: as scan [--spec-path SPEC-PATH] [--target-host TARGET-HOST]

Options:
  --spec-path SPEC-PATH
                         The file path to the OpenAPI file (both yaml and json are supported) e.g. ./openapi.yaml
  --target-host TARGET-HOST
                         The base URL of the API to be scanned e.g. https://api.nullify.ai

Global options:
  --verbose, -v          enable verbose logging
  --debug, -d            enable debug logging
  --help, -h             display this help and exit
  --version              display version and exit
```

Example scan

```sh
as scan \
  --spec-path   openapi.yml \
  --target-host http://localhost:8888
```

Example output

```json
{
  "withAuth": [
    {
      "method": "post",
      "path": "/api/users",
      "status": 401
    }
  ],
  "withoutAuth": [
    {
      "method": "get",
      "path": "/api/users/{username}",
      "status": 200
    }
  ],
  "errors": [
    {
      "method": "get",
      "path": "/api/users/{username}/profile",
      "status": 500
    },
    {
      "method": "get",
      "path": "/api/users/{username}/email",
      "error": "connection closed"
    }
  ]
}
```

- `withAuth` - endpoints that have authentication
- `withoutAuth` - endpoints that do not have authentication
- `errors` - any errors that occurred during the scan
