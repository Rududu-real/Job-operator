/*
Copyright 2022.

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

package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var taskv2log = logf.Log.WithName("taskv2-resource")

func (r *TaskV2) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-batch-platform-upstage-ai-v1-taskv2,mutating=true,failurePolicy=fail,sideEffects=None,groups=batch.platform.upstage.ai,resources=taskv2s,verbs=create;update,versions=v1,name=mtaskv2.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &TaskV2{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *TaskV2) Default() {
	taskv2log.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-batch-platform-upstage-ai-v1-taskv2,mutating=false,failurePolicy=fail,sideEffects=None,groups=batch.platform.upstage.ai,resources=taskv2s,verbs=create;update,versions=v1,name=vtaskv2.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &TaskV2{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *TaskV2) ValidateCreate() error {
	taskv2log.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *TaskV2) ValidateUpdate(old runtime.Object) error {
	taskv2log.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *TaskV2) ValidateDelete() error {
	taskv2log.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
