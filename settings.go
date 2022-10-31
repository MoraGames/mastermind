package main

type Path string
type Settings struct {
	//Environment Settings
	settingsFile Path
	recordsFile Path

	//Game Settings
	language string
	playMode string
	maxAttempts int
	codeDomain string
	codeRepeats bool

	//User Settings
	defaultUserName string
}