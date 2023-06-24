package monitor

import (
	"fmt"
	"time"

	"procmon/pkg/slack"

	"github.com/shirou/gopsutil/process"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func getProcessByName(name string) (*process.Process, float64, float64, error) {
	procs, err := process.Processes()
	if err != nil {
		return nil, 0, 0, err
	}

	for _, p := range procs {
		pName, _ := p.Name()
		if pName == name {
			cpuPercent, _ := p.CPUPercent()
			memInfo, _ := p.MemoryInfo()
			memMB := float64(memInfo.RSS) / 1024.0 / 1024.0

			return p, cpuPercent, memMB, nil
		}
	}

	return nil, 0, 0, fmt.Errorf("process not found")
}

func checkCPULoad(name string, p *process.Process, cpuPercent float64, slackToken string, channelID string) {
	if cpuPercent > 80.0 {
		slack.Send(slackToken, channelID, fmt.Sprintf(":fire: *Alert:* Process `%s` is using high CPU: `%.2f%%`", name, cpuPercent))
		log.WithFields(logrus.Fields{
			"Process": name,
			"CPU":     fmt.Sprintf("%.2f%%", cpuPercent),
		}).Warn("High CPU usage detected")
	}
}

func Start(name string, slackToken string, channelID string) {
	var p *process.Process
	processIsRunning := false
	for {
		var cpuPercent float64
		var err error
		if !processIsRunning {
			p, cpuPercent, _, err = getProcessByName(name)
			if err == nil {
				processIsRunning = true
				slack.Send(slackToken, channelID, fmt.Sprintf(":satellite: *Monitoring started for process:* `%s`\n:chart_with_upwards_trend: *Initial CPU Usage:* `%.2f%%`", name, cpuPercent))
				log.WithFields(logrus.Fields{
					"Process": name,
					"CPU":     fmt.Sprintf("%.2f%%", cpuPercent),
				}).Info("Monitoring started")
			}
		} else {
			cpuPercent, err = p.CPUPercent()
			if err != nil {
				running, _ := p.IsRunning()
				if processIsRunning && !running {
					slack.Send(slackToken, channelID, fmt.Sprintf(":warning: *Alert:* The monitored process `%s` has terminated.", name))
					log.WithFields(logrus.Fields{
						"Process": name,
					}).Error("Process terminated")
					processIsRunning = false
				}
			} else {
				checkCPULoad(name, p, cpuPercent, slackToken, channelID)
			}
		}
		time.Sleep(5 * time.Second)
	}
}
