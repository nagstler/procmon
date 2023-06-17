// Importing required packages
package monitor

import (
	"fmt"
	"procmon/pkg/slack"
	"time"

	"github.com/shirou/gopsutil/process"
	"github.com/sirupsen/logrus"
)

// Initializing the logrus logger
var log = logrus.New()

// Start function starts the monitoring of the process with the provided processName
// It also requires a slackToken and channelID for sending slack alerts
func Start(processName string, slackToken string, channelID string) {
	// An infinite loop which checks for the process and CPU usage every 5 seconds
	for {
		// Finding the process ID by the process name
		pid, err := findProcessByName(processName)
		// If the process is not found or any error occurred while fetching the process
		// it logs the error and waits for 5 seconds before the next iteration
		if err != nil {
			log.WithFields(logrus.Fields{
				"process": processName,
			}).Error("Error while finding the process: ", err)
			time.Sleep(5 * time.Second)
			continue
		}

		// Creating a new process instance with the found process ID
		p, err := process.NewProcess(pid)
		// If any error occurred while creating the new process instance,
		// it logs the error and waits for 5 seconds before the next iteration
		if err != nil {
			log.WithFields(logrus.Fields{
				"process": processName,
				"pid":     pid,
			}).Error("Cannot create process instance: ", err)
			continue
		}

		// Getting the current CPU percent of the process
		cpuPercent, err := p.CPUPercent()
		// If any error occurred while fetching the CPU percent,
		// it logs the error and waits for 5 seconds before the next iteration
		if err != nil {
			log.WithFields(logrus.Fields{
				"process": processName,
				"pid":     pid,
			}).Error("Cannot get CPU percent: ", err)
			continue
		}

		// If CPU percent is over 80, it sends an alert to the provided slack channel
		if cpuPercent > 80.0 {
			slack.Send(slackToken, channelID, fmt.Sprintf("Process %s is using too much CPU: %.2f%%", processName, cpuPercent))
			log.WithFields(logrus.Fields{
				"process": processName,
				"cpu":     cpuPercent,
			}).Warn("Process using high CPU")
		}

		// Wait for 5 seconds before the next iteration
		time.Sleep(5 * time.Second)
	}
}

// findProcessByName function finds a process ID by its name
func findProcessByName(name string) (int32, error) {
	// Getting a list of all the running processes
	ps, err := process.Processes()
	if err != nil {
		// If any error occurred while fetching the processes, return with the error
		return -1, fmt.Errorf("error fetching processes: %v", err)
	}
	// Iterating over all the processes to find a match with the provided name
	for _, p := range ps {
		processName, err := p.Name()
		// If any error occurred while fetching the process name, log the error and continue to the next process
		if err != nil {
			log.WithFields(logrus.Fields{
				"pid": p.Pid,
			}).Error("Error fetching process name: ", err)
			continue
		}
		// If the process name matches with the provided name, return the process ID
		if processName == name {
			return p.Pid, nil
		}
	}
	// If no matching process was found, return with an error
	return -1, fmt.Errorf("process not found")
}
