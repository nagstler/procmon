# ProcMon: Process Monitoring Tool

ProcMon is a simple command-line tool written in Go. It monitors a specified system process by its name and sends alerts to a designated Slack channel if the process is not found or if it's consuming too much CPU.

## Features

- Monitor any process by its name
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
./procmon --name "your-process-name" --config ./config.yaml
```

Replace `"your-process-name"` with the name of the process you want to monitor.

## Contributing

Contributions are always welcome! Please see the `CONTRIBUTING.md` file for more information.

## License

This project is licensed under the terms of the MIT license. See the `LICENSE` file for the full text.

## Disclaimer

This is a simple project for monitoring system processes. It may not be suitable for production use. Always test tools in your environment before using them for important tasks.