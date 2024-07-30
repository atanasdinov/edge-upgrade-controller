package upgrade

import (
	upgradecattlev1 "github.com/rancher/system-upgrade-controller/pkg/apis/upgrade.cattle.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

const (
	planNamespace     = "cattle-system"
	PlanAnnotation    = "lifecycle.suse.com/upgrade-plan"
	ReleaseAnnotation = "lifecycle.suse.com/release"
	controlPlaneKey   = "control-plane"
	workersKey        = "workers"

	ControlPlaneLabel = "node-role.kubernetes.io/control-plane"
)

func baseUpgradePlan(name string) *upgradecattlev1.Plan {
	const (
		kind               = "Plan"
		apiVersion         = "upgrade.cattle.io/v1"
		serviceAccountName = "system-upgrade-controller"
	)

	plan := &upgradecattlev1.Plan{
		TypeMeta: metav1.TypeMeta{
			Kind:       kind,
			APIVersion: apiVersion,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: planNamespace,
		},
		Spec: upgradecattlev1.PlanSpec{
			ServiceAccountName: serviceAccountName,
		},
	}

	return plan
}

func PlanNamespacedName(plan string) types.NamespacedName {
	return types.NamespacedName{
		Name:      plan,
		Namespace: planNamespace,
	}
}
