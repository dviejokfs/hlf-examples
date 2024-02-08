package main

import (
	"context"
	operatorv1 "github.com/kfsoftware/hlf-operator/pkg/client/clientset/versioned"
	"github.com/sirupsen/logrus"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path"
)

func getHLFClientByKubeConfig(kubeConfig string) (*operatorv1.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		return nil, err
	}
	hlfClient, err := operatorv1.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return hlfClient, nil
}
func main() {
	var kubeConfig string
	kubeconfigEnv := os.Getenv("KUBECONFIG")
	if kubeconfigEnv != "" {
		logrus.Infof("Using kubeconfig from env: %s", kubeconfigEnv)
		kubeConfig = kubeconfigEnv
	} else {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			panic(err.Error())
		}
		kubeConfig = path.Join(homeDir, ".kube", "config")
		logrus.Infof("Using default kubeconfig: %s", kubeConfig)
	}
	hlfClient, err := getHLFClientByKubeConfig(kubeConfig)
	if err != nil {
		panic(err.Error())
	}
	peerCuratorList, err := hlfClient.HlfV1alpha1().FabricPeers("").List(context.Background(), v1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	logrus.Infof("List of peers:")
	for _, peer := range peerCuratorList.Items {
		logrus.Infof("Peer: %s/%s mspID: %s", peer.Name, peer.Namespace, peer.Spec.MspID)
	}
}
