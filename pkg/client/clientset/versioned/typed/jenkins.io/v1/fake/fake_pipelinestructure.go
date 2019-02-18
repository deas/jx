// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	jenkins_io_v1 "github.com/jenkins-x/jx/pkg/apis/jenkins.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePipelineStructures implements PipelineStructureInterface
type FakePipelineStructures struct {
	Fake *FakeJenkinsV1
	ns   string
}

var pipelinestructuresResource = schema.GroupVersionResource{Group: "jenkins.io", Version: "v1", Resource: "pipelinestructures"}

var pipelinestructuresKind = schema.GroupVersionKind{Group: "jenkins.io", Version: "v1", Kind: "PipelineStructure"}

// Get takes name of the pipelineStructure, and returns the corresponding pipelineStructure object, and an error if there is any.
func (c *FakePipelineStructures) Get(name string, options v1.GetOptions) (result *jenkins_io_v1.PipelineStructure, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(pipelinestructuresResource, c.ns, name), &jenkins_io_v1.PipelineStructure{})

	if obj == nil {
		return nil, err
	}
	return obj.(*jenkins_io_v1.PipelineStructure), err
}

// List takes label and field selectors, and returns the list of PipelineStructures that match those selectors.
func (c *FakePipelineStructures) List(opts v1.ListOptions) (result *jenkins_io_v1.PipelineStructureList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(pipelinestructuresResource, pipelinestructuresKind, c.ns, opts), &jenkins_io_v1.PipelineStructureList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &jenkins_io_v1.PipelineStructureList{ListMeta: obj.(*jenkins_io_v1.PipelineStructureList).ListMeta}
	for _, item := range obj.(*jenkins_io_v1.PipelineStructureList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pipelineStructures.
func (c *FakePipelineStructures) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(pipelinestructuresResource, c.ns, opts))

}

// Create takes the representation of a pipelineStructure and creates it.  Returns the server's representation of the pipelineStructure, and an error, if there is any.
func (c *FakePipelineStructures) Create(pipelineStructure *jenkins_io_v1.PipelineStructure) (result *jenkins_io_v1.PipelineStructure, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(pipelinestructuresResource, c.ns, pipelineStructure), &jenkins_io_v1.PipelineStructure{})

	if obj == nil {
		return nil, err
	}
	return obj.(*jenkins_io_v1.PipelineStructure), err
}

// Update takes the representation of a pipelineStructure and updates it. Returns the server's representation of the pipelineStructure, and an error, if there is any.
func (c *FakePipelineStructures) Update(pipelineStructure *jenkins_io_v1.PipelineStructure) (result *jenkins_io_v1.PipelineStructure, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(pipelinestructuresResource, c.ns, pipelineStructure), &jenkins_io_v1.PipelineStructure{})

	if obj == nil {
		return nil, err
	}
	return obj.(*jenkins_io_v1.PipelineStructure), err
}

// Delete takes name of the pipelineStructure and deletes it. Returns an error if one occurs.
func (c *FakePipelineStructures) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(pipelinestructuresResource, c.ns, name), &jenkins_io_v1.PipelineStructure{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePipelineStructures) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(pipelinestructuresResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &jenkins_io_v1.PipelineStructureList{})
	return err
}

// Patch applies the patch and returns the patched pipelineStructure.
func (c *FakePipelineStructures) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *jenkins_io_v1.PipelineStructure, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(pipelinestructuresResource, c.ns, name, data, subresources...), &jenkins_io_v1.PipelineStructure{})

	if obj == nil {
		return nil, err
	}
	return obj.(*jenkins_io_v1.PipelineStructure), err
}
