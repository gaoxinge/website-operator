package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Metadata struct {
	Name      string
	Namespace string
}

type Spec struct {
	GitRepo string
}

type Website struct {
	Metadata Metadata
	Spec     Spec
}

func NewWebsite(name string, namespace string, gitRepo string) *Website {
	return &Website{
		Metadata: Metadata{
			Name:      name,
			Namespace: namespace,
		},
		Spec: Spec{
			GitRepo: gitRepo,
		},
	}
}

func (website *Website) Deployment() *Deployment {
	return NewDeployment(website.Metadata.Name, website.Metadata.Namespace, website.Spec.GitRepo)
}

func (website *Website) Service() *Service {
	return NewService(website.Metadata.Name, website.Metadata.Namespace)
}

func CreateWebsite(clientSet kubernetes.Interface, website *Website) {
	CreateDeployment(clientSet, website.Deployment())
	CreateService(clientSet, website.Service())
}

func DeleteWebsite(clientSet kubernetes.Interface, website *Website) {
	DeleteDeployment(clientSet, website.Deployment())
	DeleteService(clientSet, website.Service())
}

type WebsiteEvent struct {
	Type   string
	Object Website
}

type WebsiteController struct {
	Host      string
	ClientSet kubernetes.Interface
}

func NewWebsiteController(host string) (*WebsiteController, error) {
	config := rest.Config{
		Host: host,
	}

	clientSet, err := kubernetes.NewForConfig(&config)
	if err != nil {
		return nil, err
	}

	return &WebsiteController{
		Host:      host,
		ClientSet: clientSet,
	}, nil
}

func (websiteController *WebsiteController) Run() {
	for {
		websiteController.RunOnce()
	}
}

func (websiteController *WebsiteController) RunOnce() {
	response, err := http.Get(fmt.Sprintf("%s/apis/extensions.example.com/v1/websites?watch=true", websiteController.Host))
	if err != nil {
		log.Printf("run controller with error %v", err)
		return
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	for {
		var websiteEvent WebsiteEvent
		err := decoder.Decode(&websiteEvent)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("run controller %v with error %v", websiteEvent, err)
			return
		}

		switch websiteEvent.Type {
		case "ADDED":
			websiteController.CreateWebsite(&websiteEvent.Object)
		case "DELETED":
			websiteController.DeleteWebsite(&websiteEvent.Object)
		}
	}
}

func (websiteController *WebsiteController) CreateWebsite(website *Website) {
	CreateWebsite(websiteController.ClientSet, website)
}

func (websiteController *WebsiteController) DeleteWebsite(website *Website) {
	DeleteWebsite(websiteController.ClientSet, website)
}
