package swissknife

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	simplecron "github.com/sagleft/simple-cron"
	"gopkg.in/yaml.v3"
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

func WaitForAppFinish() {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	<-exit
	fmt.Println()
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

func ParseStructFromJSON(jsonBytes []byte, destinationPointer interface{}) error {
	return json.Unmarshal(jsonBytes, destinationPointer)
}

func ParseStructFromYaml(jsonBytes []byte, destinationPointer interface{}) error {
	return yaml.Unmarshal(jsonBytes, destinationPointer)
}

func ParseStructFromJSONFile(filepath string, destinationPointer interface{}) error {
	dataBytes, err := ReadFileToBytes(filepath)
	if err != nil {
		return err
	}

	return ParseStructFromJSON(dataBytes, destinationPointer)
}

func ParseStructFromYamlFile(filepath string, destinationPointer interface{}) error {
	dataBytes, err := ReadFileToBytes(filepath)
	if err != nil {
		return err
	}

	return ParseStructFromYaml(dataBytes, destinationPointer)
}

// MD5 - calc MD5 checksum
func MD5(val []byte) string {
	hash := md5.Sum(val)
	return hex.EncodeToString(hash[:])
}
