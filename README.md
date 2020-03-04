# alexandria

A simple DNS server

## Getting started

Build the docker image and configure with the following environment variables:

```ini
LOG_LEVEL="TRACE|DEBUG|INFO|WARN|ERROR|FATAL|PANIC - the log level"
LOG_JSON="true|false - whether or not to log in JSON format to STDIO"
PRINT_TITLE="true|false - whether or not to print the title screen/version number to STDIO"
HOSTNAME="the hostname on which to serve the DNS server"
PORT="the port on which to server the DNS server"
```

## Contributors

- [Ariel Simulevski](https://github.com/Azer0s) - creator and maintainer

## Credits

Mad props go to

- Simon Eskildsen for [Logrus](https://github.com/sirupsen/logrus)
