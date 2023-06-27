# ProcMon: Multiprocess Monitoring with Slack Integration 🕵️‍♀️
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) [![Go Version](https://img.shields.io/badge/Go-1.17-blue.svg)](https://golang.org/dl/) [![Go CI Build](https://github.com/nagstler/procmon/actions/workflows/main.yml/badge.svg)](https://github.com/nagstler/procmon/actions/workflows/main.yml) [![Maintainability](https://api.codeclimate.com/v1/badges/4020d2d5bb982e89047a/maintainability)](https://codeclimate.com/github/nagstler/procmon/maintainability) [![GitHub release](https://img.shields.io/github/release/nagstler/procmon.svg)](https://github.com/nagstler/procmon/releases/)


ProcMon is a simple command-line tool, built in Go, that monitors one or more system processes simultaneously by their names. If a process isn't found, consumes too much CPU, or gets terminated, it sends alerts directly to a specified Slack channel.

![SLACK-IMG](https://github.com/nagstler/procmon/assets/1298480/a61602ab-5f58-43d9-b563-216e386af486)


## Features

- Monitor any process by its name
- Monitor multiple processes simultaneously
- Sends an alert to a Slack channel if the process is not found or terminated
- Sends an alert if a process's CPU usage exceeds 80%
- Resumes monitoring when the process is restarted

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
processes:
  - 'process-name-1'
  - 'process-name-2'
  - 'process-name-3'
```

Please remember to replace 'your-slack-token' and 'your-channel-id' with your actual Slack bot token and channel ID, and 'process-name-1', 'process-name-2', 'process-name-3' with the names of the processes you want to monitor.

**Never commit your Slack bot token to version control.**

To start monitoring a process, run the following:

```bash
./procmon
```

## Contributing

Bug reports and pull requests are welcome on GitHub at [https://github.com/nagstler/procmon](https://github.com/nagstler/procmon). 
This project is intended to be a safe, welcoming space for collaboration, and contributors are expected to adhere to the [code of conduct](https://github.com/nagstler/procmon/blob/main/CODE_OF_CONDUCT.md) .

## License

This project is licensed under the terms of the MIT license. See the `LICENSE` file for the full text.
