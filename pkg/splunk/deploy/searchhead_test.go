// Copyright (c) 2018-2020 Splunk Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package deploy

import (
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	enterprisev1 "github.com/splunk/splunk-operator/pkg/apis/enterprise/v1alpha2"
)

func TestReconcileSearchHead(t *testing.T) {
	funcCalls := []mockFuncCall{
		{metaName: "*v1.Secret-test-splunk-stack1-secrets"},
		{metaName: "*v1.Service-test-splunk-stack1-search-head-headless"},
		{metaName: "*v1.Service-test-splunk-stack1-search-head-service"},
		{metaName: "*v1.Service-test-splunk-stack1-deployer-service"},
		{metaName: "*v1.StatefulSet-test-splunk-stack1-deployer"},
		{metaName: "*v1.StatefulSet-test-splunk-stack1-search-head"},
	}
	createCalls := map[string][]mockFuncCall{"Get": funcCalls, "Create": funcCalls}
	updateCalls := map[string][]mockFuncCall{"Get": funcCalls, "Update": []mockFuncCall{funcCalls[4], funcCalls[5]}}
	current := enterprisev1.SearchHead{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "stack1",
			Namespace: "test",
		},
	}
	revised := current.DeepCopy()
	revised.Spec.Image = "splunk/test"
	reconcile := func(c *mockClient, cr interface{}) error {
		return ReconcileSearchHead(c, cr.(*enterprisev1.SearchHead))
	}
	reconcileTester(t, "TestReconcileSearchHead", &current, revised, createCalls, updateCalls, reconcile)
}

func TestApplySearchHead(t *testing.T) {
	funcCalls := []mockFuncCall{{metaName: "*v1alpha2.SearchHead-test-stack1"}}
	createCalls := map[string][]mockFuncCall{"Get": funcCalls, "Create": funcCalls}
	updateCalls := map[string][]mockFuncCall{"Get": funcCalls, "Update": funcCalls}
	current := enterprisev1.SplunkEnterprise{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "stack1",
			Namespace: "test",
		},
	}
	revised := current.DeepCopy()
	revised.Spec.SplunkImage = "splunk/test"
	reconcile := func(c *mockClient, cr interface{}) error {
		return applySearchHead(c, cr.(*enterprisev1.SplunkEnterprise))
	}
	reconcileTester(t, "TestApplySearchHead", &current, revised, createCalls, updateCalls, reconcile)
}