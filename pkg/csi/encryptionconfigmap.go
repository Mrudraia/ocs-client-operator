/*
Copyright 2022 Red Hat, Inc.

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

package csi

import (
	"context"

	"github.com/red-hat-storage/ocs-client-operator/pkg/templates"

	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

func CreateEncryptionConfigMap(ctx context.Context, c client.Client, log klog.Logger) error {
	monConfigMap := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: templates.EncryptionConfigMapName,
		},
		Data: map[string]string{
			"config.json": "[]",
		},
	}
	err := controllerutil.SetControllerReference(OperatorDeployment, monConfigMap, c.Scheme())
	if err != nil {
		return err
	}

	err = c.Create(ctx, monConfigMap)
	if err != nil && !k8serrors.IsAlreadyExists(err) {
		log.Error(err, "failed to create encryption configmap", "name", monConfigMap.Name)
		return err
	}

	log.Info("successfully created encryption configmap", "name", monConfigMap.Name)
	return nil
}