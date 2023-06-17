package monitor

import (
	"fmt"
	"time"

	"procmon/pkg/slack"

	"github.com/shirou/gopsutil/process"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func Start(pid int32, slackToken string, channelID string) {
	p, err := process.NewProcess(pid)
	if err != nil {
		log.WithFields(logrus.Fields{
			"pid": pid,
		}).Error("No process found with the provided PID: ", err)
		return
	}

	name, err := p.Name()
	if err != nil {
		log.WithFields(logrus.Fields{
			"pid": pid,
		}).Error("Unable to fetch process name: ", err)
		return
	}

	cpuPercent, err := p.CPUPercent()
	if err != nil {
		log.WithFields(logrus.Fields{
			"pid": pid,
		}).Error("Unable to fetch CPU percent: ", err)
		return
	}

	memInfo, err := p.MemoryInfo()
	if err != nil {
		log.WithFields(logrus.Fields{
			"pid": pid,
		}).Error("Unable to fetch memory info: ", err)
		return
	}

	memMB := float64(memInfo.RSS) / 1024.0 / 1024.0

	slackMessage := fmt.Sprintf(
		":satellite: *Monitoring started for process:* :satellite:\n\n"+
			":label: *Name:* `%s`\n"+
			":id: *PID:* `%d`\n"+
			":chart_with_upwards_trend: *Initial CPU Usage:* `%.2f%%`\n"+
			":bar_chart: *Initial Memory Usage:* `%.2f MB`",
		name,
		pid,
		cpuPercent,
		memMB,
	)
	slack.Send(slackToken, channelID, slackMessage)

	for {
		cpuPercent, err := p.CPUPercent()
		if err != nil {
			running, err := p.IsRunning()
			if err != nil {
				log.WithFields(logrus.Fields{
					"pid": pid,
				}).Error("Unable to determine if process is running: ", err)
				time.Sleep(5 * time.Second)
				continue
			}
			if !running {
				slack.Send(slackToken, channelID, fmt.Sprintf(":warning: The monitored process with PID `%d` has terminated.", pid))
				log.WithFields(logrus.Fields{
					"pid": pid,
				}).Info("Monitored process has terminated")
				return
			}

			log.WithFields(logrus.Fields{
				"pid": pid,
			}).Error("Cannot get CPU percent: ", err)
			time.Sleep(5 * time.Second)
			continue
		}

		if cpuPercent > 80.0 {
			slack.Send(slackToken, channelID, fmt.Sprintf(":fire: Process with PID `%d` is using high CPU: `%.2f`", pid, cpuPercent))
			log.WithFields(logrus.Fields{
				"pid": pid,
				"cpu": cpuPercent,
			}).Warn("Process using high CPU")
		}

		time.Sleep(5 * time.Second)
	}
}
