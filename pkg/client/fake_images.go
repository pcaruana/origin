package client

import (
	"github.com/GoogleCloudPlatform/kubernetes/pkg/fields"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/labels"

	imageapi "github.com/openshift/origin/pkg/image/api"
)

// FakeImages implements ImageInterface. Meant to be embedded into a struct to
// get a default implementation. This makes faking out just the methods you
// want to test easier.
type FakeImages struct {
	Fake *Fake
}

var _ ImageInterface = &FakeImages{}

func (c *FakeImages) List(label labels.Selector, field fields.Selector) (*imageapi.ImageList, error) {
	obj, err := c.Fake.Invokes(FakeAction{Action: "list-images"}, &imageapi.ImageList{})
	return obj.(*imageapi.ImageList), err
}

func (c *FakeImages) Get(name string) (*imageapi.Image, error) {
	obj, err := c.Fake.Invokes(FakeAction{Action: "get-image", Value: name}, &imageapi.Image{})
	return obj.(*imageapi.Image), err
}

func (c *FakeImages) Create(image *imageapi.Image) (*imageapi.Image, error) {
	obj, err := c.Fake.Invokes(FakeAction{Action: "create-image"}, &imageapi.Image{})
	return obj.(*imageapi.Image), err
}

func (c *FakeImages) Delete(name string) error {
	c.Fake.Actions = append(c.Fake.Actions, FakeAction{Action: "delete-image", Value: name})
	_, err := c.Fake.Invokes(FakeAction{Action: "delete-image", Value: name}, nil)
	return err
}
