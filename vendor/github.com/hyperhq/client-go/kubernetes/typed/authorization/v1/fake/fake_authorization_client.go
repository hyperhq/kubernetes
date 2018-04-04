/*
Copyright 2018 The Kubernetes Authors.

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
	v1 "github.com/hyperhq/client-go/kubernetes/typed/authorization/v1"
	rest "github.com/hyperhq/client-go/rest"
	testing "github.com/hyperhq/client-go/testing"
)

type FakeAuthorizationV1 struct {
	*testing.Fake
}

func (c *FakeAuthorizationV1) LocalSubjectAccessReviews(namespace string) v1.LocalSubjectAccessReviewInterface {
	return &FakeLocalSubjectAccessReviews{c, namespace}
}

func (c *FakeAuthorizationV1) SelfSubjectAccessReviews() v1.SelfSubjectAccessReviewInterface {
	return &FakeSelfSubjectAccessReviews{c}
}

func (c *FakeAuthorizationV1) SelfSubjectRulesReviews() v1.SelfSubjectRulesReviewInterface {
	return &FakeSelfSubjectRulesReviews{c}
}

func (c *FakeAuthorizationV1) SubjectAccessReviews() v1.SubjectAccessReviewInterface {
	return &FakeSubjectAccessReviews{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeAuthorizationV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}