/*
Copyright The KCP Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"

	v1alpha1 "github.com/kcp-dev/kcp/pkg/apis/tenancy/v1alpha1"
)

// FakeClusterWorkspaceTypes implements ClusterWorkspaceTypeInterface
type FakeClusterWorkspaceTypes struct {
	Fake *FakeTenancyV1alpha1
}

var clusterworkspacetypesResource = schema.GroupVersionResource{Group: "tenancy.kcp.dev", Version: "v1alpha1", Resource: "clusterworkspacetypes"}

var clusterworkspacetypesKind = schema.GroupVersionKind{Group: "tenancy.kcp.dev", Version: "v1alpha1", Kind: "ClusterWorkspaceType"}

// Get takes name of the clusterWorkspaceType, and returns the corresponding clusterWorkspaceType object, and an error if there is any.
func (c *FakeClusterWorkspaceTypes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterWorkspaceType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterworkspacetypesResource, name), &v1alpha1.ClusterWorkspaceType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterWorkspaceType), err
}

// List takes label and field selectors, and returns the list of ClusterWorkspaceTypes that match those selectors.
func (c *FakeClusterWorkspaceTypes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterWorkspaceTypeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterworkspacetypesResource, clusterworkspacetypesKind, opts), &v1alpha1.ClusterWorkspaceTypeList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterWorkspaceTypeList{ListMeta: obj.(*v1alpha1.ClusterWorkspaceTypeList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterWorkspaceTypeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterWorkspaceTypes.
func (c *FakeClusterWorkspaceTypes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterworkspacetypesResource, opts))
}

// Create takes the representation of a clusterWorkspaceType and creates it.  Returns the server's representation of the clusterWorkspaceType, and an error, if there is any.
func (c *FakeClusterWorkspaceTypes) Create(ctx context.Context, clusterWorkspaceType *v1alpha1.ClusterWorkspaceType, opts v1.CreateOptions) (result *v1alpha1.ClusterWorkspaceType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterworkspacetypesResource, clusterWorkspaceType), &v1alpha1.ClusterWorkspaceType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterWorkspaceType), err
}

// Update takes the representation of a clusterWorkspaceType and updates it. Returns the server's representation of the clusterWorkspaceType, and an error, if there is any.
func (c *FakeClusterWorkspaceTypes) Update(ctx context.Context, clusterWorkspaceType *v1alpha1.ClusterWorkspaceType, opts v1.UpdateOptions) (result *v1alpha1.ClusterWorkspaceType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterworkspacetypesResource, clusterWorkspaceType), &v1alpha1.ClusterWorkspaceType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterWorkspaceType), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterWorkspaceTypes) UpdateStatus(ctx context.Context, clusterWorkspaceType *v1alpha1.ClusterWorkspaceType, opts v1.UpdateOptions) (*v1alpha1.ClusterWorkspaceType, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clusterworkspacetypesResource, "status", clusterWorkspaceType), &v1alpha1.ClusterWorkspaceType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterWorkspaceType), err
}

// Delete takes name of the clusterWorkspaceType and deletes it. Returns an error if one occurs.
func (c *FakeClusterWorkspaceTypes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(clusterworkspacetypesResource, name, opts), &v1alpha1.ClusterWorkspaceType{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterWorkspaceTypes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterworkspacetypesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterWorkspaceTypeList{})
	return err
}

// Patch applies the patch and returns the patched clusterWorkspaceType.
func (c *FakeClusterWorkspaceTypes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterWorkspaceType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterworkspacetypesResource, name, pt, data, subresources...), &v1alpha1.ClusterWorkspaceType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterWorkspaceType), err
}