package monitor

import (
	"fmt"
	"time"

	"procmon/pkg/slack"

	"github.com/shirou/gopsutil/process"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func getProcessInfo(pid int32) (*process.Process, string, float64, float64, error) {
	p, err := process.NewProcess(pid)
	if err != nil {
		return nil, "", 0, 0, err
	}

	name, err := p.Name()
	if err != nil {
		return nil, "", 0, 0, err
	}

	cpuPercent, err := p.CPUPercent()
	if err != nil {
		return nil, "", 0, 0, err
	}

	memInfo, err := p.MemoryInfo()
	if err != nil {
		return nil, "", 0, 0, err
	}

	memMB := float64(memInfo.RSS) / 1024.0 / 1024.0

	return p, name, cpuPercent, memMB, nil
}

func checkCPULoad(pid int32, p *process.Process, cpuPercent float64, slackToken string, channelID string) {
	if cpuPercent > 80.0 {
		slack.Send(slackToken, channelID, fmt.Sprintf(":fire: Process with PID `%d` is using high CPU: `%.2f`", pid, cpuPercent))
		log.WithFields(logrus.Fields{
			"pid": pid,
			"cpu": cpuPercent,
		}).Warn("Process using high CPU")
	}
}

func Start(pid int32, slackToken string, channelID string) {
	p, name, cpuPercent, memMB, err := getProcessInfo(pid)
	if err != nil {
		log.WithFields(logrus.Fields{
			"pid": pid,
		}).Error("Unable to fetch process info: ", err)
		return
	}

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

		checkCPULoad(pid, p, cpuPercent, slackToken, channelID)

		time.Sleep(5 * time.Second)
	}
}
