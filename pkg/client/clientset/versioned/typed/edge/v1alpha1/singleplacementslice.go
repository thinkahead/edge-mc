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

package v1alpha1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	v1alpha1 "github.com/kcp-dev/edge-mc/pkg/apis/edge/v1alpha1"
	scheme "github.com/kcp-dev/edge-mc/pkg/client/clientset/versioned/scheme"
)

// SinglePlacementSlicesGetter has a method to return a SinglePlacementSliceInterface.
// A group's client should implement this interface.
type SinglePlacementSlicesGetter interface {
	SinglePlacementSlices() SinglePlacementSliceInterface
}

// SinglePlacementSliceInterface has methods to work with SinglePlacementSlice resources.
type SinglePlacementSliceInterface interface {
	Create(ctx context.Context, singlePlacementSlice *v1alpha1.SinglePlacementSlice, opts v1.CreateOptions) (*v1alpha1.SinglePlacementSlice, error)
	Update(ctx context.Context, singlePlacementSlice *v1alpha1.SinglePlacementSlice, opts v1.UpdateOptions) (*v1alpha1.SinglePlacementSlice, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.SinglePlacementSlice, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.SinglePlacementSliceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SinglePlacementSlice, err error)
	SinglePlacementSliceExpansion
}

// singlePlacementSlices implements SinglePlacementSliceInterface
type singlePlacementSlices struct {
	client rest.Interface
}

// newSinglePlacementSlices returns a SinglePlacementSlices
func newSinglePlacementSlices(c *EdgeV1alpha1Client) *singlePlacementSlices {
	return &singlePlacementSlices{
		client: c.RESTClient(),
	}
}

// Get takes name of the singlePlacementSlice, and returns the corresponding singlePlacementSlice object, and an error if there is any.
func (c *singlePlacementSlices) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SinglePlacementSlice, err error) {
	result = &v1alpha1.SinglePlacementSlice{}
	err = c.client.Get().
		Resource("singleplacementslices").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SinglePlacementSlices that match those selectors.
func (c *singlePlacementSlices) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SinglePlacementSliceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SinglePlacementSliceList{}
	err = c.client.Get().
		Resource("singleplacementslices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested singlePlacementSlices.
func (c *singlePlacementSlices) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("singleplacementslices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a singlePlacementSlice and creates it.  Returns the server's representation of the singlePlacementSlice, and an error, if there is any.
func (c *singlePlacementSlices) Create(ctx context.Context, singlePlacementSlice *v1alpha1.SinglePlacementSlice, opts v1.CreateOptions) (result *v1alpha1.SinglePlacementSlice, err error) {
	result = &v1alpha1.SinglePlacementSlice{}
	err = c.client.Post().
		Resource("singleplacementslices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(singlePlacementSlice).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a singlePlacementSlice and updates it. Returns the server's representation of the singlePlacementSlice, and an error, if there is any.
func (c *singlePlacementSlices) Update(ctx context.Context, singlePlacementSlice *v1alpha1.SinglePlacementSlice, opts v1.UpdateOptions) (result *v1alpha1.SinglePlacementSlice, err error) {
	result = &v1alpha1.SinglePlacementSlice{}
	err = c.client.Put().
		Resource("singleplacementslices").
		Name(singlePlacementSlice.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(singlePlacementSlice).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the singlePlacementSlice and deletes it. Returns an error if one occurs.
func (c *singlePlacementSlices) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("singleplacementslices").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *singlePlacementSlices) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("singleplacementslices").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched singlePlacementSlice.
func (c *singlePlacementSlices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SinglePlacementSlice, err error) {
	result = &v1alpha1.SinglePlacementSlice{}
	err = c.client.Patch(pt).
		Resource("singleplacementslices").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}