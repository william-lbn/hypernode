/*
Copyright The Kubernetes Authors.

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

	v1alpha1 "github.com/william-lbn/hypernode/pkg/apis/pml/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHyperNodes implements HyperNodeInterface
type FakeHyperNodes struct {
	Fake *FakePmlV1alpha1
	ns   string
}

var hypernodesResource = schema.GroupVersionResource{Group: "pml", Version: "v1alpha1", Resource: "hypernodes"}

var hypernodesKind = schema.GroupVersionKind{Group: "pml", Version: "v1alpha1", Kind: "HyperNode"}

// Get takes name of the hyperNode, and returns the corresponding hyperNode object, and an error if there is any.
func (c *FakeHyperNodes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.HyperNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(hypernodesResource, c.ns, name), &v1alpha1.HyperNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HyperNode), err
}

// List takes label and field selectors, and returns the list of HyperNodes that match those selectors.
func (c *FakeHyperNodes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.HyperNodeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(hypernodesResource, hypernodesKind, c.ns, opts), &v1alpha1.HyperNodeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.HyperNodeList{ListMeta: obj.(*v1alpha1.HyperNodeList).ListMeta}
	for _, item := range obj.(*v1alpha1.HyperNodeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested hyperNodes.
func (c *FakeHyperNodes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(hypernodesResource, c.ns, opts))

}

// Create takes the representation of a hyperNode and creates it.  Returns the server's representation of the hyperNode, and an error, if there is any.
func (c *FakeHyperNodes) Create(ctx context.Context, hyperNode *v1alpha1.HyperNode, opts v1.CreateOptions) (result *v1alpha1.HyperNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(hypernodesResource, c.ns, hyperNode), &v1alpha1.HyperNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HyperNode), err
}

// Update takes the representation of a hyperNode and updates it. Returns the server's representation of the hyperNode, and an error, if there is any.
func (c *FakeHyperNodes) Update(ctx context.Context, hyperNode *v1alpha1.HyperNode, opts v1.UpdateOptions) (result *v1alpha1.HyperNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(hypernodesResource, c.ns, hyperNode), &v1alpha1.HyperNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HyperNode), err
}

// Delete takes name of the hyperNode and deletes it. Returns an error if one occurs.
func (c *FakeHyperNodes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(hypernodesResource, c.ns, name, opts), &v1alpha1.HyperNode{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHyperNodes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(hypernodesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.HyperNodeList{})
	return err
}

// Patch applies the patch and returns the patched hyperNode.
func (c *FakeHyperNodes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.HyperNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(hypernodesResource, c.ns, name, pt, data, subresources...), &v1alpha1.HyperNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HyperNode), err
}
