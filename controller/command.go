package controller

import (
	"errors"
)

type CommandKey string

const (
	CommandStart    CommandKey = "start"
	CommandShutdown            = "shutdown"
	CommandAdd                 = "add"
	CommandRemove              = "remove"
	CommandList                = "list"
)

type CommandTarget string

const (
	CommandTargetMonitor   CommandTarget = "monitor"
	CommandTargetEmail                   = "email"
	CommandTargetLocations               = "locations"
	CommandTargetSites                   = "sites"
)

type CommandOption string

const (
	CommandOptionConfig            CommandOption = "config"
	CommandOptionPort                            = "port"
	CommandOptionTo                              = "to"
	CommandOptionFrom                            = "from"
	CommandOptionFor                             = "for"
	CommandOptionLocation                        = "location"
	CommandOptionSites                           = "sites"
	CommandOptionHighPrioritySites               = "hpsites"
)

// some default constants
const (
	defaultConfigFile = "reservations.json"
)

type CommandInfo struct {
	Command    CommandKey
	ConfigFile string
}

func ResolveCommandKeyString(key string) (CommandKey, error) {
	if key == (string)(CommandStart) {
		return CommandStart, nil
	} else if key == (string)(CommandShutdown) {
		return CommandShutdown, nil
	} else if key == (string)(CommandAdd) {
		return CommandAdd, nil
	} else if key == (string)(CommandRemove) {
		return CommandRemove, nil
	} else if key == (string)(CommandList) {
		return CommandList, nil
	} else {
		return (CommandKey)(""), errors.New("Cannot find command key from string [" + key + "]")
		//panic("Cannot find command key from string [" + key + "]")
	}
}

func ResolveCommandTargetString(key string) (CommandTarget, error) {
	if key == (string)(CommandTargetMonitor) {
		return CommandTargetMonitor, nil
	} else if key == (string)(CommandTargetEmail) {
		return CommandTargetEmail, nil
	} else if key == (string)(CommandTargetLocations) {
		return CommandTargetLocations, nil
	} else if key == (string)(CommandTargetSites) {
		return CommandTargetSites, nil
	} else {
		return (CommandTarget)(""), errors.New("Cannot find command target from string [" + key + "]")
	}
}
func ResolveCommandOptionString(key string) (CommandOption, error) {
	if key == (string)(CommandOptionConfig) {
		return CommandOptionConfig, nil
	} else if key == (string)(CommandOptionPort) {
		return CommandOptionPort, nil
	} else if key == (string)(CommandOptionTo) {
		return CommandOptionTo, nil
	} else if key == (string)(CommandOptionFrom) {
		return CommandOptionFrom, nil
	} else if key == (string)(CommandOptionFor) {
		return CommandOptionFor, nil
	} else if key == (string)(CommandOptionLocation) {
		return CommandOptionLocation, nil
	} else if key == (string)(CommandOptionSites) {
		return CommandOptionSites, nil
	} else if key == (string)(CommandOptionHighPrioritySites) {
		return CommandOptionHighPrioritySites, nil
	} else {
		return (CommandOption)(""), errors.New("Cannot find command option from string [" + key + "]")
	}
}

type ControlCommand struct {
	Command     CommandKey
	Target      CommandTarget
	Options     map[CommandOption]string
	TargetValue string
}
