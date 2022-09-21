package main

import (
	"fmt"
	"github.com/fatih/color"
	"golang.org/x/exp/slices"
	"log"
	"os/exec"
)

const (
	cmdGetPods        = "kubectl get pods"
	cmdTopPods        = "kubectl top pods"
	cmdGetNamespaces  = "kubectl get namespaces"
	cmdGetDeployments = "kubectl get deployment"
	cmdGetServices    = "kubectl get service"
	cmdGetSecrets     = "kubectl get secret"
	watch             = "--watch"
	namespace         = "--namespace"
	wide              = "-o=wide"
)

// getPods is getting pods from kubernetes cluster
func getPods(args []string) {
	c := color.New(color.FgWhite).Add(color.BgBlue)
	var out []byte
	var err error
	var cmd = cmdGetPods

	_, _ = c.Printf("\nListing Pods:")

	if len(args) > 1 && args[1] == "w" {
		cmd = cmdGetPods + " " + watch
	}

	if slices.Contains(args[:], "o") {
		cmd = cmd + " " + wide
	}

	cmd = namespaceCheck(cmd, args)

	out, err = exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n\n" + string(out))

	if slices.Contains(args[:], "m") {
		c := color.New(color.FgBlack).Add(color.BgCyan)
		_, _ = c.Printf("\nListing Metrics:")
		cmd = namespaceCheck(cmdTopPods, args)

		out, err = exec.Command("bash", "-c", cmd).Output()

		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("\n\n" + string(out))
	}
}

// getNamespaces is getting namespaces from kubernetes cluster
func getNamespaces() {
	c := color.New(color.FgBlack).Add(color.BgYellow)

	_, _ = c.Printf("\nListing Namespaces:")

	out, err := exec.Command("bash", "-c", cmdGetNamespaces).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n\n" + string(out))
}

// getServices is getting services from kubernetes cluster
func getServices(args []string) {
	c := color.New(color.FgWhite).Add(color.BgMagenta)
	var cmd = cmdGetServices

	_, _ = c.Printf("\nListing Services:")

	cmd = namespaceCheck(cmd, args)
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n\n" + string(out))
}

// getDeployments is getting deployments from kubernetes cluster
func getDeployments(args []string) {
	c := color.New(color.FgBlack).Add(color.BgGreen)
	var cmd = cmdGetDeployments

	_, _ = c.Printf("\nListing Deployments:")

	cmd = namespaceCheck(cmd, args)
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n\n" + string(out))
}

// getSecrets is getting secrets from kubernetes clusters
func getSecrets(args []string) {
	c := color.New(color.FgWhite).Add(color.BgRed)
	var cmd = cmdGetSecrets

	_, _ = c.Printf("\nListing Secrets:")

	cmd = namespaceCheck(cmd, args)
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n\n" + string(out))
}

// namespaceCheck is adding namespace parameter if namespace parameter is present
func namespaceCheck(cmd string, args []string) string {
	if len(args) > 2 && slices.Contains(args[:], "-n") && args[len(args)-1] != "" {
		cmd = cmd + " " + namespace + "=" + args[len(args)-1]
	}
	return cmd
}
