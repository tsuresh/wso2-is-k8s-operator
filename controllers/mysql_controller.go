/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	wso2v1 "github.com/tsuresh/wso2-is-k8s-operator/api/v1"
	mysql "kubernetes-charts.storage.googleapis.co"
)

// MySQLReconciler reconciles a MySQL object
type MySQLReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=wso2.wso2.com,resources=mysqls,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=wso2.wso2.com,resources=mysqls/status,verbs=get;update;patch

func (r *MySQLReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("mysql", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *MySQLReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&wso2v1.MySQL{}).
		Complete(r)
}

// addConfigMap adds a new ConfigMap
func (r *Wso2IsReconciler) addMySQLResource(m wso2v1.Wso2Is) *corev1.ConfigMap {
	configMap := &corev1.MySQL{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "identity-server-conf",
			Namespace: m.Namespace,
		},
		Data: map[string]string{
			"deployment.toml": m.Spec.Configurations,
		},
	}
	ctrl.SetControllerReference(&m, configMap, r.Scheme)
	return configMap
}
