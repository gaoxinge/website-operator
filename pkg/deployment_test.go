package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func TestDeployment(t *testing.T) {
	config := rest.Config{
		Host: "http://127.0.0.1:8001",
	}

	clientSet, err := kubernetes.NewForConfig(&config)
	assert.NoError(t, err)

	deployment := NewDeployment(
		"kubia",
		apiv1.NamespaceDefault,
		"https://github.com.cnpmjs.org/luksa/kubia-website-example.git",
	)

	CreateDeployment(clientSet, deployment)
	DeleteDeployment(clientSet, deployment)
}
