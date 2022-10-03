package main

import (
	"context"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/metrics/pkg/apis/metrics/v1beta1"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
)

func newClient() *kubernetes.Clientset {
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, &clientcmd.ConfigOverrides{})
	config, err := kubeConfig.ClientConfig()

	if err != nil {
		panic(err)
	}

	client := kubernetes.NewForConfigOrDie(config)
	return client
}

func getNamespace() string {
	clientCfg, _ := clientcmd.NewDefaultClientConfigLoadingRules().Load()
	namespace := clientCfg.Contexts[clientCfg.CurrentContext].Namespace
	return namespace
}

func getMetricsClient(namespace string) (*v1beta1.PodMetricsList, error) {
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, &clientcmd.ConfigOverrides{})

	config, _ := kubeConfig.ClientConfig()
	mc, _ := metrics.NewForConfig(config)
	podMetrics, err := mc.MetricsV1beta1().PodMetricses(namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return &v1beta1.PodMetricsList{}, err
	}
	return podMetrics, nil
}
