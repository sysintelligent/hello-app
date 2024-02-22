package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Configuration holds the application version
type Configuration struct {
	Version string `json:"version"`
}

var (
	podStatus = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "hello_app_pod_status",
			Help: "Status of the hello-app pod (1 = running, 0 = not running)",
		},
	)
)

func main() {
	// Load the configuration file
	config, err := loadConfiguration("config.json")
	if err != nil {
		log.Fatal("Error loading configuration:", err)
	}

	// Initialize Kubernetes client
	clientset, err := getKubernetesClient()
	if err != nil {
		log.Fatal("Error creating Kubernetes client:", err)
	}

	// Set up metrics collection
	go func() {
		for {
			updatePodStatus(clientset)
			time.Sleep(30 * time.Second) // Update every 30 seconds
		}
	}()

	// Set up the HTTP handler for application endpoints
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := fmt.Sprintf("hello-app says hi! [version: %s]", config.Version)
		fmt.Fprintln(w, response)
	})

	// Register Prometheus metrics
	prometheus.MustRegister(podStatus)

	// Register Prometheus metrics handler
	http.Handle("/metrics", promhttp.Handler())

	// Start the server
	log.Println("Starting hello-app server...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting the server:", err)
	}
}

func loadConfiguration(filename string) (Configuration, error) {
	var config Configuration

	// Read the configuration file
	data, err := os.ReadFile(filename)
	if err != nil {
		return config, err
	}

	// Unmarshal the JSON data into the Configuration struct
	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

// Initialize Kubernetes client
func getKubernetesClient() (*kubernetes.Clientset, error) {
	// Load Kubernetes configuration
	config, err := rest.InClusterConfig()
	if err != nil {
		kubeconfig := os.Getenv("KUBECONFIG")
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			log.Fatalf("Error loading Kubernetes configuration: %v", err)
		}
	}

	// Create Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}

// Update the status of the hello-app pod
func updatePodStatus(clientset *kubernetes.Clientset) {
	// List pods in the "hello-app" namespace with the label selector app=hello-app
	pods, err := clientset.CoreV1().Pods("hello-app").List(context.TODO(), metav1.ListOptions{
		LabelSelector: "app=hello-app",
	})
	if err != nil {
		log.Fatalf("Error listing pods: %v", err)
	}

	// Iterate over each pod and update its status
	for _, pod := range pods.Items {
		// Update Prometheus metric based on pod status
		if pod.Status.Phase == corev1.PodRunning {
			// podStatus.Set(1)
			log.Printf("Pod %s is running\n", pod.Name)
		} else {
			// podStatus.Set(0)
			log.Printf("Pod %s is not running\n", pod.Name)
		}
	}
}
