package main

import (
	"context"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var clientset *kubernetes.Clientset

func main() {
	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic("构建配置失败: " + err.Error())
	}

	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic("创建客户端失败: " + err.Error())
	}

	router := gin.Default()

	router.GET("/pods", getPods)
	router.GET("/nodes", getNodes)
	router.GET("/namespaces", getNamespaces)

	router.Run(":8080")
}

func getPods(c *gin.Context) {
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var result []gin.H
	for _, pod := range pods.Items {
		result = append(result, gin.H{
			"namespace": pod.Namespace,
			"name":      pod.Name,
			"status":    string(pod.Status.Phase),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"count": len(result),
		"pods":  result,
	})
}

func getNodes(c *gin.Context) {
	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var result []gin.H
	for _, node := range nodes.Items {
		result = append(result, gin.H{
			"name": node.Name,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"count": len(result),
		"nodes": result,
	})
}

func getNamespaces(c *gin.Context) {
	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var result []gin.H
	for _, ns := range namespaces.Items {
		result = append(result, gin.H{
			"name": ns.Name,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"count":      len(result),
		"namespaces": result,
	})
}
