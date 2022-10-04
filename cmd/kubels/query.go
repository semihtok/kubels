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
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"golang.org/x/exp/slices"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	"k8s.io/metrics/pkg/apis/metrics/v1beta1"
	"os"
	"os/exec"
	"strconv"
)

// getPods is getting pods from kubernetes cluster
func getPods(args []string) error {
	c := color.New(color.FgWhite).Add(color.BgBlue)
	var err error
	var metrics *v1beta1.PodMetricsList
	var namespace string

	fmt.Printf("\n")
	_, _ = c.Printf("Listing Pods:")
	fmt.Printf("\n\n")

	client := newClient()

	if slices.Contains(args[:], "-n") {
		namespace = args[slices.Index(args[:], "-n")+1]
	} else {
		namespace = getNamespace()
	}

	pods, err := client.CoreV1().Pods(namespace).List(context.TODO(), v1.ListOptions{})
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	if slices.Contains(args[:], "-m") {
		metrics, err = getMetricsClient(namespace)
		if err != nil {
			return err
		}
		table.SetHeader([]string{"Name", "Status", "Restarts", "CPU", "Memory"})

	} else if slices.Contains(args[:], "-o") {
		table.SetHeader([]string{"Name", "Status", "Restarts", "CPU", "Memory", "Node", "IP", "Node IP"})
		for _, pod := range pods.Items {
			table.Append([]string{
				pod.Name,
				string(pod.Status.Phase),
				strconv.Itoa(int(pod.Status.ContainerStatuses[0].RestartCount)),
				pod.Status.PodIP,
				pod.Spec.NodeName,
				pod.Status.HostIP,
			})
		}
	} else {
		table.SetHeader([]string{"Name", "Status", "Restarts"})
	}

	for i, pod := range pods.Items {
		if metrics != nil && len(metrics.Items) > 0 {
			cpu := metrics.Items[i].Containers[0].Usage.Cpu().MilliValue()
			memory := metrics.Items[i].Containers[0].Usage.Memory().Value() / 1024 / 1024

			table.Append([]string{
				pod.Name,
				string(pod.Status.Phase),
				strconv.Itoa(int(pod.Status.ContainerStatuses[0].RestartCount)),
				strconv.Itoa(int(cpu)) + "m",
				strconv.Itoa(int(memory)) + "Mi",
			})
		} else {
			table.Append([]string{
				pod.Name,
				string(pod.Status.Phase),
				strconv.Itoa(int(pod.Status.ContainerStatuses[0].RestartCount)),
			})
		}
	}
	table.Render()
	return nil
}

// getNamespaces is getting namespaces from kubernetes cluster
func getNamespaceList() error {
	c := color.New(color.FgBlack).Add(color.BgYellow)
	_, _ = c.Printf("\nListing Namespaces:\n\n")

	client := newClient()
	list, err := client.CoreV1().Namespaces().List(context.TODO(), v1.ListOptions{})
	if err != nil {
		return err
	}

	var rows [][]string
	for i, ns := range list.Items {
		rows = append(rows, []string{strconv.Itoa(i + 1), ns.ObjectMeta.Name})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "Name"})

	for _, row := range rows {
		table.Append(row)
	}
	table.Render()
	return nil
}

// getServices is getting services from kubernetes cluster
func getServices(args []string) error {
	c := color.New(color.FgBlack).Add(color.BgYellow)
	_, _ = c.Printf("\nListing Services:\n\n")
	var namespace string

	if slices.Contains(args[:], "-n") {
		namespace = args[slices.Index(args[:], "-n")+1]
	} else {
		namespace = getNamespace()
	}

	client := newClient()
	list, err := client.CoreV1().Services(namespace).List(context.TODO(), v1.ListOptions{})
	if err != nil {
		return err
	}

	var rows [][]string
	for _, svc := range list.Items {
		rows = append(rows, []string{svc.ObjectMeta.Name, svc.Spec.ClusterIP})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Cluster IP"})
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	for _, row := range rows {
		table.Append(row)
	}

	table.Render()
	return nil
}

// getDeployments is getting deployments from kubernetes cluster
func getDeployments(args []string) error {
	c := color.New(color.FgBlack).Add(color.BgHiMagenta)
	_, _ = c.Printf("\nListing Deployments:\n\n")
	var namespace string

	if slices.Contains(args[:], "-n") {
		namespace = args[slices.Index(args[:], "-n")+1]
	} else {
		namespace = getNamespace()
	}

	client := newClient()
	list, err := client.AppsV1().Deployments(namespace).List(context.TODO(), v1.ListOptions{})
	if err != nil {
		return err
	}

	var rows [][]string
	for _, dep := range list.Items {
		rows = append(rows, []string{dep.ObjectMeta.Name, strconv.Itoa(int(dep.Status.AvailableReplicas))})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Available Replicas"})
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	for _, row := range rows {
		table.Append(row)
	}

	table.Render()
	return nil
}

// getSecrets is getting secrets from kubernetes clusters
func getSecrets(args []string) error {
	c := color.New(color.FgWhite).Add(color.BgRed)
	_, _ = c.Printf("\nListing Secrets:")
	var namespace string

	if slices.Contains(args[:], "-n") {
		namespace = args[slices.Index(args[:], "-n")+1]
	} else {
		namespace = getNamespace()
	}

	client := newClient()
	list, err := client.CoreV1().Secrets(namespace).List(context.TODO(), v1.ListOptions{})
	if err != nil {
		return err
	}

	var rows [][]string
	for _, sec := range list.Items {
		rows = append(rows, []string{sec.ObjectMeta.Name})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name"})
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	for _, row := range rows {
		table.Append(row)
	}

	table.Render()
	return nil
}

func switchNamespace(namespace string) error {
	c := color.New(color.FgWhite).Add(color.BgMagenta)
	var cmd = "kubectl config set-context --current --namespace=" + namespace

	_, _ = c.Printf("\nSwitching Namespace to:" + namespace)

	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return err
	}
	fmt.Printf("\n\n" + string(out))
	return nil
}
