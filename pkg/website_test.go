package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func TestWebsite(t *testing.T) {
	config := rest.Config{
		Host: "http://127.0.0.1:8001",
	}

	clientSet, err := kubernetes.NewForConfig(&config)
	assert.NoError(t, err)

	website := NewWebsite(
		"kubia",
		apiv1.NamespaceDefault,
		"https://github.com.cnpmjs.org/luksa/kubia-website-example.git",
	)

	CreateWebsite(clientSet, website)
	DeleteWebsite(clientSet, website)
}
