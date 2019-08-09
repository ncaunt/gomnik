# gomnik

A tool written in `golang` which reads data from certain Omnik solar inverters and exposes it as `Prometheus` metrics.

## Usage

```shell
usage: gomnik --serial=SERIAL --address=ADDRESS [<flags>]

Flags:
      --help                    Show context-sensitive help (also try --help-long and --help-man).
  -s, --serial=SERIAL           Serial number of the inverter.
  -a, --address=ADDRESS         Address on which the inverter listens (example: 10.0.0.1:8899).
  -i, --interval=10             Number of seconds between queries of the inverter.
  -m, --metrics="0.0.0.0:9100"  Endpoint on which to serve Prometheus metrics.
      --version                 Show application version.
```

`gomnik` will open a TCP connection to the inverter whose address is specified by the parameter `--address` and request data every `--interval` seconds.

The parameter `--serial` must be set to the serial number of the inverter to be queried.

It serves `Prometheus` metrics on the HTTP endpoint specified in `--metrics`.

Tested with `go1.12.7`.

## TODO

* Tidy the code.
* Write proper tests.

## References

Based upon [Omnik-Data-Logger](https://github.com/Woutrrr/Omnik-Data-Logger) by Wouter van der Zwan.
