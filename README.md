# Revealer
Revealer reveals secrets.

## Usage

### Print Version
```sh
# From this directory
$ ./revealer -v # or long form --version
v0.1.0
```

---
### Print Help
```sh
$ ./revealer -h # or long form --help
Usage of ./revealer:
  -h, --help                  Prints this help content.
  -f, --secrets-file string   Secrets file to parse.
  -v, --version               Prints the version of Revealer.
```

---
### Validate Infra Settings
>  Infra-settings are expected to be under a YAML structrure like the one below:
```yaml
infraSettings:
  KEY: VALUE
  ...
```

```
$ ./revealer -f sample-infra-secrets.yaml
WARNING! Newlines are added in the supplied settingsValue: bG9jYWxob3N0OjMzMDYK for settingsKey: DATABASE_HOST Correct value should be: bG9jYWxob3N0OjMzMDY=
DATABASE_HOST: localhost:3306
```

---
## Building binary from source
```sh
go build -o revealer $PWD/cmd/revealer

# Or simply just run this ready-made shell script
./generate-binaries.sh # This will be prepare linux and mac binaries based on the latest version (defined in version.go)
```

---
## Make a fresh release
```sh
./cut-tag.sh # This will be cut a fresh tag based on version defined in version.go and pushes the same to Github.
```
