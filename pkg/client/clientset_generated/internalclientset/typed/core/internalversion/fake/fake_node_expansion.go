/*
Copyright 2016 The Kubernetes Authors.

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

package fake

import (
	core "github.com/hyperhq/client-go/testing"
	api "k8s.io/kubernetes/pkg/apis/core"
)

func (c *FakeNodes) PatchStatus(nodeName string, data []byte) (*api.Node, error) {
	obj, err := c.Fake.Invokes(
		core.NewRootPatchSubresourceAction(nodesResource, nodeName, data, "status"), &api.Node{})
	if obj == nil {
		return nil, err
	}

	return obj.(*api.Node), err
}
