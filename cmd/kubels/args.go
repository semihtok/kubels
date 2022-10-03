/*
 * Copyright 2022 Semih Tok
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"
)

func Run(args []string) {
	firstArg := "p"

	if len(args) > 0 {
		firstArg = args[0]

		if firstArg == "-p" || firstArg == "pods" {
			err := getPods(args)
			if err != nil {
				panic(err)
			}
		}

		if (firstArg == "-n" || firstArg == "ns" || firstArg == "namespaces") && len(args) == 1 {
			err := getNamespaceList()
			if err != nil {
				panic(err)
			}
		}

		if (firstArg == "-n" || firstArg == "ns" || firstArg == "namespaces") && len(args) == 2 {
			err := switchNamespace(args[1])
			if err != nil {
				panic(err)
			}
		}

		if firstArg == "-s" || firstArg == "svc" || firstArg == "services" {
			err := getServices(args)
			if err != nil {
				panic(err)
			}
		}

		if firstArg == "-d" || firstArg == "dp" || firstArg == "deployments" {
			err := getDeployments(args)
			if err != nil {
				panic(err)
			}
		}
		if firstArg == "-sec" || firstArg == "sec" || firstArg == "secrets" {
			err := getSecrets(args)
			if err != nil {
				panic(err)
			}
		}

		if firstArg == "-h" || firstArg == "-help" {
			helpText := `
Usage:	
  kubels [command]

Available Commands:
 kubels                               : list of pods in current namespace
 kubels -p               	      : list of pods in current namespace
 kubels -p -n {namespace}             : list of pods in specified namespace
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

		err := getPods(args)
		if err != nil {
			return
		}

		err = getServices(args)
		if err != nil {
			return
		}

		err = getDeployments(args)
		if err != nil {
			return
		}

		if err != nil {
			panic(err)
		}
	}
}
