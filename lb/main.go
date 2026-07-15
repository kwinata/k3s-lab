package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{
		LabelSelector: "app=go-server",
	})
	if err != nil {
		panic(err.Error())
	}

	idx := rand.Intn(len(pods.Items))
	pod := pods.Items[idx]
	addr := pod.Status.PodIP + ":8080"

	target, err := url.Parse("http://" + addr)
	if err != nil {
		panic(err)
	}
	httputil.NewSingleHostReverseProxy(target).ServeHTTP(w, r)
}

func main() {
	fmt.Printf("Load balancer started\n")
	err := http.ListenAndServe(":8080", &handler{})
	if err != nil {
		panic(err)
	}
}
