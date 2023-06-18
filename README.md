# ProcMon: Process Monitoring Tool
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) [![Go Version](https://img.shields.io/badge/Go-1.17-blue.svg)](https://golang.org/dl/) [![Go CI Build](https://github.com/nagstler/procmon/actions/workflows/main.yml/badge.svg)](https://github.com/nagstler/procmon/actions/workflows/main.yml) [![Maintainability](https://api.codeclimate.com/v1/badges/4020d2d5bb982e89047a/maintainability)](https://codeclimate.com/github/nagstler/procmon/maintainability)


ProcMon is a simple command-line tool written in Go. It monitors a specified system process by its PID (Process ID) and sends alerts to a designated Slack channel if the process is not found or if it's consuming too much CPU.

## Features

- Monitor any process by its PID
- Sends an alert to a Slack channel if the process is not found
- Sends an alert if the process's CPU usage exceeds 80%

## Prerequisites

- [Go](https://golang.org/dl/) (1.16 or higher)
- [Slack Bot Token](https://api.slack.com/authentication/basics)

## Installation

Clone the repository:

```bash
git clone https://github.com/nagstler/procmon.git
cd procmon
```

Build the tool:

```bash
go build -o procmon
```

## Usage

To use ProcMon, you need to provide a configuration file in YAML format (`config.yaml`). The file should contain your Slack bot token and the ID of the Slack channel where you want to send alerts:

```yaml
slack:
  token: 'your-slack-token'
  channel: 'your-channel-id'
```

Please remember to replace `'your-slack-token'` and `'your-channel-id'` with your actual Slack bot token and channel ID. **Never commit your Slack bot token to version control.**

To start monitoring a process, run:

```bash
./procmon --pid 12345 --config ./config.yaml
```

Replace `12345` with the PID of the process you want to monitor.

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/nagstler/procmon. This project is intended to be a safe, welcoming space for collaboration, and contributors are expected to adhere to the [code of conduct](https://github.com/nagstler/procmon/blob/main/CODE_OF_CONDUCT.md).

## License

This project is licensed under the terms of the MIT license. See the `LICENSE` file for the full text.