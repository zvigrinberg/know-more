package core

type EnvironmentParser interface {
	ParseCommandLIne(block string) string
	ParseRC(block string) int8
	ParseEnvironmentVariables(block string) map[string]string
	ParseWorkingDirectory(block string) string
}
