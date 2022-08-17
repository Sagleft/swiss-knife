package swissknife

import (
	"errors"
	"strconv"
	"time"

	simplecron "github.com/sagleft/simple-cron"
)

// Ternary operator. conditional operator
// usage example: var res = ternary(val > 0, "positive", "negative")
func Ternary(statement bool, a, b interface{}) interface{} {
	if statement {
		return a
	}
	return b
}

// RunInBackground - blocking method with no exit
func RunInBackground() {
	forever := make(chan bool)
	// background work
	<-forever
}

// ConnFunc - some func
type ConnFunc func() error

type ReconnectTask struct {
	ConnectionDescription      string // connection name
	ReconnectionAttemptsNumber int
	ConnectionTimeout          time.Duration
	ReconnectAfterTimeout      time.Duration
	WaitingBetweenAttempts     time.Duration

	// callbacks
	ConnCallback  ConnFunc
	ErrorCallback func(error)
	LogCallback   func(string)
}

// Reconnect - setup reconnect
func Reconnect(task ReconnectTask) {
	isConnected := false
	for !isConnected {
		for i := 0; i < task.ReconnectionAttemptsNumber; i++ {

			isTimeIsUP := simplecron.NewRuntimeLimitHandler(
				task.ConnectionTimeout*time.Second,
				func() {
					err := task.ConnCallback()
					if err == nil {
						// connection established
						isConnected = true
						return
					}
					task.LogCallback(err.Error())
					time.Sleep(task.ReconnectAfterTimeout)
				},
			).Run()
			if isTimeIsUP {
				task.LogCallback(task.ConnectionDescription + " connection went into timeout")
			}

		}

		if isConnected {
			task.LogCallback("`" + task.ConnectionDescription + "` connection established")
			return
		}

		task.ErrorCallback(errors.New("failed to connect to " +
			task.ConnectionDescription + " after " +
			strconv.Itoa(task.ReconnectionAttemptsNumber) + " attempts"))

		task.LogCallback("wait " + task.WaitingBetweenAttempts.String() + " between attempts...")
		time.Sleep(task.WaitingBetweenAttempts)
	}
}
