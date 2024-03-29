# bbox_exporter

A Prometheus exporter for the Bbox, a Set-Top-Box (TV box) provided by French Internet Service Provider Bouygues Telecom.

[![License Apache 2][badge-license]](LICENSE)
[![GitHub version](https://badge.fury.io/gh/nlamirault%2Fbbox_exporter.svg)](https://badge.fury.io/gh/nlamirault%2Fbbox_exporter)

Metrics are :

| Name                                               | Exposed informations                                  | Labels               |
| -------------------------------------------------- | ------------------------------------------------------| ---------------------|
| `bbox_device_cpu`                                  | CPU Time                                              | `mode`               |
| `bbox_device_memory`                               | Memory in kB                                          | ̀`type`               |
| `bbox_device_process`                              | Processus                                             | `type`               |
| `bbox_device_status`                               | Current status                                        |
| `bbox_device_temperature`                          | Current internal temperature in °C                    |
| `bbox_dns_average`                                 | Average of average dns response time                  |
| `bbox_dns_max`                                     | Maximun of average dns response time                  |
| `bbox_dns_min`                                     | Minimun of average dns response time                  |
| `bbox_dns_number_of_queries`                       | Number of queries                                     |
| `bbox_lan_received_bytes`                          | RX bytes                                              |
| `bbox_lan_received_packets`                        | RX packets                                            |
| `bbox_lan_received_packets_discards`               | RX packets discards                                   |
| `bbox_lan_received_packets_errors`                 | RX packets in error                                   |
| `bbox_lan_transmitted_bytes`                       | TX bytes                                              |
| `bbox_lan_transmitted_packets`                     | TX packets                                            |
| `bbox_lan_transmitted_packets_discards`            | TX packets discards                                   |
| `bbox_lan_transmitted_packets_errors`              | TX packets in error                                   |
| `bbox_up`                                          | Was the last query of BBox successful.                |
| `bbox_wan_ftth_state`                              | LinkState of the GEth FTTH port                       |
| `bbox_wan_received_bandwidth`                      | RX bandwith available                                 |
| `bbox_wan_received_bandwidth_max`                  | RX bandwith available                                 |
| `bbox_wan_received_bytes`                          | RX bytes                                              |
| `bbox_wan_received_packets`                        | RX packets                                            |
| `bbox_wan_received_packets_discards`               | RX packets discards                                   |
| `bbox_wan_received_packets_errors`                 | RX packets in error                                   |
| `bbox_wan_transmitted_bandwidth`                   | TX bandwith available                                 |
| `bbox_wan_transmitted_bandwidth_max`               | TX maximum bandwith available                         |
| `bbox_wan_transmitted_bytes`                       | TX bytes                                              |
| `bbox_wan_transmitted_packets`                     | TX packets                                            |
| `bbox_wan_transmitted_packets_discards`            | TX packets discards                                   |
| `bbox_wan_transmitted_packets_errors`              | TX packets in error                                   |


![Dashboard](dashboard.png)
## Usage

Launch the Prometheus exporter :

    > bbox_exporter --help
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
## Contact

Nicolas Lamirault <nicolas.lamirault@gmail.com>

[badge-license]: https://img.shields.io/badge/license-Apache2-green.svg?style=flat
