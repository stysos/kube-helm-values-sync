package value_retrieval

import (
	"helm.sh/helm/v3/pkg/action"

	"fmt"

)

type HelmRetriever struct {
}

func (h *HelmRetriever) GetValues(name string) {
	getVals := action.NewGetValues()

	fmt.Println(getVals.Run(name))
}
