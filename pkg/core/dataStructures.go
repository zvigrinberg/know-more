package core

import "time"

type TraceRecord struct {
	TimeOfInvocation  time.Time
	TimeIntercepted   time.Time
	Command           string
	ProcessId         string
	CommandWorkingDir string
	ReturnCode        int8
	EnvVars           map[string]string
}
