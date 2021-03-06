package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PodNodeConstraintsConfig is the configuration for the pod node name
// and node selector constraint plug-in. For accounts, serviceaccounts
// and groups which lack the "pods/binding" permission, Loading this
// plugin will prevent setting NodeName on pod specs and will prevent
// setting NodeSelectors whose labels appear in the blacklist field
// "NodeSelectorLabelBlacklist"
type PodNodeConstraintsConfig struct {
	metav1.TypeMeta `json:",inline"`

	// NodeSelectorLabelBlacklist specifies a list of labels which cannot be set by entities without the "pods/binding" permission
	NodeSelectorLabelBlacklist []string `json:"nodeSelectorLabelBlacklist",description:"list of labels which cannot be set by entities without the 'pods/binding' permission"`
}
