package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	cli "github.com/ansalamdaniel/hadepl/pkg/client/clientset/versioned"
	hinformer "github.com/ansalamdaniel/hadepl/pkg/client/informers/externalversions/ansimatt.dev/v1alpha1"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/home/ansalam/.kube/config", "(optional) absolute path to the kubeconfig file")

	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Printf("error %s when loading the config file\n", err.Error())
		config, err = rest.InClusterConfig()
		if err != nil {
			fmt.Printf("error %s when trying to load the cofig from cluster\n", err.Error())
		}
	}

	clientset, err := cli.NewForConfig(config)
	if err != nil {
		fmt.Printf("error %s", err.Error())
	}

	fmt.Println(clientset)

	had, err := clientset.AnsimattV1alpha1().HADeployments("default").List(context.Background(), v1.ListOptions{})
	if err != nil {
		fmt.Printf("getting error %s in had retrieval\n", err.Error())
	}

	fmt.Printf("length of HADeployments %d\n", len(had.Items))

	inFactory, err := hinformer.NewHADeploymentInformer(clientset, 10*time.Minute)
	if err != nil {
		fmt.Println(err.Error())
	}
}
