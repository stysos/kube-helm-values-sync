package value_retrieval

import (
	"helm.sh/helm/v3/pkg/action"

	"fmt"
	"log"
	"os"

	"helm.sh/helm/v3/pkg/cli"
)

type HelmRetriever struct {
}

func (h *HelmRetriever) GetValues(name string) {

	settings := cli.New()

	actionConfig := new(action.Configuration)
	// You can pass an empty string instead of settings.Namespace() to list
	// all namespaces
	if err := actionConfig.Init(settings.RESTClientGetter(), settings.Namespace(), os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
		log.Printf("%+v", err)
		os.Exit(1)
	}

	client := action.NewList(actionConfig)
	// Only list deployed
	client.Deployed = true
	results, err := client.Run()
	if err != nil {
		log.Printf("%+v", err)
		os.Exit(1)
	}

	for _, rel := range results {
		log.Printf("%+v", rel)
	}
	getVals := action.NewGetValues(actionConfig)

	fmt.Println(getVals.Run(name))
}
