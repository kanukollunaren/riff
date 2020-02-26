/*
Copyright 2019 the original author or authors.

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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"

	v1alpha1 "github.com/projectriff/riff/system/pkg/apis/streaming/v1alpha1"
)

// FakeKafkaGateways implements KafkaGatewayInterface
type FakeKafkaGateways struct {
	Fake *FakeStreamingV1alpha1
	ns   string
}

var kafkagatewaysResource = schema.GroupVersionResource{Group: "streaming.projectriff.io", Version: "v1alpha1", Resource: "kafkagatewaies"}

var kafkagatewaysKind = schema.GroupVersionKind{Group: "streaming.projectriff.io", Version: "v1alpha1", Kind: "KafkaGateway"}

// Get takes name of the kafkaGateway, and returns the corresponding kafkaGateway object, and an error if there is any.
func (c *FakeKafkaGateways) Get(name string, options v1.GetOptions) (result *v1alpha1.KafkaGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kafkagatewaysResource, c.ns, name), &v1alpha1.KafkaGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KafkaGateway), err
}

// List takes label and field selectors, and returns the list of KafkaGateways that match those selectors.
func (c *FakeKafkaGateways) List(opts v1.ListOptions) (result *v1alpha1.KafkaGatewayList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(kafkagatewaysResource, kafkagatewaysKind, c.ns, opts), &v1alpha1.KafkaGatewayList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KafkaGatewayList{ListMeta: obj.(*v1alpha1.KafkaGatewayList).ListMeta}
	for _, item := range obj.(*v1alpha1.KafkaGatewayList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kafkaGateways.
func (c *FakeKafkaGateways) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kafkagatewaysResource, c.ns, opts))

}

// Create takes the representation of a kafkaGateway and creates it.  Returns the server's representation of the kafkaGateway, and an error, if there is any.
func (c *FakeKafkaGateways) Create(kafkaGateway *v1alpha1.KafkaGateway) (result *v1alpha1.KafkaGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(kafkagatewaysResource, c.ns, kafkaGateway), &v1alpha1.KafkaGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KafkaGateway), err
}

// Update takes the representation of a kafkaGateway and updates it. Returns the server's representation of the kafkaGateway, and an error, if there is any.
func (c *FakeKafkaGateways) Update(kafkaGateway *v1alpha1.KafkaGateway) (result *v1alpha1.KafkaGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(kafkagatewaysResource, c.ns, kafkaGateway), &v1alpha1.KafkaGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KafkaGateway), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKafkaGateways) UpdateStatus(kafkaGateway *v1alpha1.KafkaGateway) (*v1alpha1.KafkaGateway, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(kafkagatewaysResource, "status", c.ns, kafkaGateway), &v1alpha1.KafkaGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KafkaGateway), err
}

// Delete takes name of the kafkaGateway and deletes it. Returns an error if one occurs.
func (c *FakeKafkaGateways) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(kafkagatewaysResource, c.ns, name), &v1alpha1.KafkaGateway{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKafkaGateways) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(kafkagatewaysResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.KafkaGatewayList{})
	return err
}

// Patch applies the patch and returns the patched kafkaGateway.
func (c *FakeKafkaGateways) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KafkaGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(kafkagatewaysResource, c.ns, name, pt, data, subresources...), &v1alpha1.KafkaGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KafkaGateway), err
}
