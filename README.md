# bbox_exporter

A Prometheus exporter for the Bbox Miami, a Set-Top-Box (TV box) provided by French Internet Service Provider Bouygues Telecom.

[![License Apache 2][badge-license]](LICENSE)
[![GitHub version](https://badge.fury.io/gh/nlamirault%2Fbbox_exporter.svg)](https://badge.fury.io/gh/nlamirault%2Fbbox_exporter)

* Master : [![Circle CI](https://circleci.com/gh/nlamirault/bbox_exporter/tree/master.svg?style=svg)](https://circleci.com/gh/nlamirault/bbox_exporter/tree/master)
* Develop : [![Circle CI](https://circleci.com/gh/nlamirault/bbox_exporter/tree/develop.svg?style=svg)](https://circleci.com/gh/nlamirault/bbox_exporter/tree/develop)


Metrics are :



## Installation

You can download the binaries :

* Architecture amd64 [ [linux](https://bintray.com/artifact/download/nlamirault/oss/bbox_exporter-0.1.0_linux_amd64) / [darwin](https://bintray.com/artifact/download/nlamirault/oss/bbox_exporter-0.1.0_darwin_amd64) / [windows](https://bintray.com/artifact/download/nlamirault/oss/bbox_exporter-0.1.0_windows_amd64.exe) ]
* Architecture arm [ [linux](https://bintray.com/artifact/download/nlamirault/oss/bbox_exporter-0.1.0_linux_arm) ]
* Architecture arm64 [ [linux](https://bintray.com/artifact/download/nlamirault/oss/bbox_exporter-0.1.0_linux_arm) ]


## Usage

Launch the Prometheus exporter :

    $ bbox_exporter -log.level=debug


## Development

* Initialize environment

        $ make init

* Build tool :

        $ make build

* Launch unit tests :

        $ make test


## Local Deployment

* Launch Prometheus using the configuration file in this repository:

        $ prometheus -config.file=prometheus.yml

* Launch exporter:

        $ bbox_exporter -log.level=debug

* Check that Prometheus find the exporter on `http://localhost:9090/targets`


## Contributing

See [CONTRIBUTING](CONTRIBUTING.md).


## License

See [LICENSE](LICENSE) for the complete license.


## Changelog

A [changelog](ChangeLog.md) is available


## Contact

Nicolas Lamirault <nicolas.lamirault@gmail.com>

[badge-license]: https://img.shields.io/badge/license-Apache2-green.svg?style=flat