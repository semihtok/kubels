package main

import "fmt"

func Run(args []string) {

	firstArg := "p"

	if len(args) > 0 {
		firstArg = args[0]

		if firstArg == "-p" || firstArg == "pods" {
			getPods(args)
		}

		if firstArg == "-n" || firstArg == "ns" || firstArg == "namespaces" {
			getNamespaces()
		}

		if firstArg == "-s" || firstArg == "svc" || firstArg == "services" {
			getServices(args)
		}

		if firstArg == "-d" || firstArg == "dp" || firstArg == "deployments" {
			getDeployments(args)
		}
		if firstArg == "-sec" || firstArg == "sec" || firstArg == "secrets" {
			getSecrets(args)
		}

		if firstArg == "-h" {
			helpText := `
Commands for kubels:

 kubels                               : list of pods in current namespace
 kubels -p               	      : list of pods in current namespace
 kubels -p -n {namespace}             : list of pods in specified namespace
 kubels -p w         		      : list of pods in current namespace with watch
 kubels -p o                          : list of pods in current namespace with wide output
 kubels -p m                          : list of pods in current namespace with metrics
 kubels -d                            : list of deployments in current namespace
 kubels -d -n {namespace}             : list of deployments in specified namespace
 kubels -s                            : list of services in current namespace
 kubels -s (or svc) -n {namespace}    : list of services in specified namespace
 kubels -sec                          : list of secrets in current namespace
 kubels -s (or svc) -n {namespace}    : list of secrets in specified namespace`
			fmt.Println(helpText)
		}
	} else {
		getPods(args)
	}
}
