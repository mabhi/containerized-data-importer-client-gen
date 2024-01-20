// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	"time"

	scheme "github.com/mabhi/containerized-data-importer-client-gen/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1beta1 "kubevirt.io/containerized-data-importer-api/pkg/apis/core/v1beta1"
)

// StorageProfilesGetter has a method to return a StorageProfileInterface.
// A group's client should implement this interface.
type StorageProfilesGetter interface {
	StorageProfiles() StorageProfileInterface
}

// StorageProfileInterface has methods to work with StorageProfile resources.
type StorageProfileInterface interface {
	Create(ctx context.Context, storageProfile *v1beta1.StorageProfile, opts v1.CreateOptions) (*v1beta1.StorageProfile, error)
	Update(ctx context.Context, storageProfile *v1beta1.StorageProfile, opts v1.UpdateOptions) (*v1beta1.StorageProfile, error)
	UpdateStatus(ctx context.Context, storageProfile *v1beta1.StorageProfile, opts v1.UpdateOptions) (*v1beta1.StorageProfile, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.StorageProfile, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.StorageProfileList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.StorageProfile, err error)
	StorageProfileExpansion
}

// storageProfiles implements StorageProfileInterface
type storageProfiles struct {
	client rest.Interface
}

// newStorageProfiles returns a StorageProfiles
func newStorageProfiles(c *CdiV1beta1Client) *storageProfiles {
	return &storageProfiles{
		client: c.RESTClient(),
	}
}

// Get takes name of the storageProfile, and returns the corresponding storageProfile object, and an error if there is any.
func (c *storageProfiles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.StorageProfile, err error) {
	result = &v1beta1.StorageProfile{}
	err = c.client.Get().
		Resource("storageprofiles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of StorageProfiles that match those selectors.
func (c *storageProfiles) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.StorageProfileList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.StorageProfileList{}
	err = c.client.Get().
		Resource("storageprofiles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested storageProfiles.
func (c *storageProfiles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("storageprofiles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a storageProfile and creates it.  Returns the server's representation of the storageProfile, and an error, if there is any.
func (c *storageProfiles) Create(ctx context.Context, storageProfile *v1beta1.StorageProfile, opts v1.CreateOptions) (result *v1beta1.StorageProfile, err error) {
	result = &v1beta1.StorageProfile{}
	err = c.client.Post().
		Resource("storageprofiles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(storageProfile).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a storageProfile and updates it. Returns the server's representation of the storageProfile, and an error, if there is any.
func (c *storageProfiles) Update(ctx context.Context, storageProfile *v1beta1.StorageProfile, opts v1.UpdateOptions) (result *v1beta1.StorageProfile, err error) {
	result = &v1beta1.StorageProfile{}
	err = c.client.Put().
		Resource("storageprofiles").
		Name(storageProfile.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(storageProfile).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *storageProfiles) UpdateStatus(ctx context.Context, storageProfile *v1beta1.StorageProfile, opts v1.UpdateOptions) (result *v1beta1.StorageProfile, err error) {
	result = &v1beta1.StorageProfile{}
	err = c.client.Put().
		Resource("storageprofiles").
		Name(storageProfile.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(storageProfile).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the storageProfile and deletes it. Returns an error if one occurs.
func (c *storageProfiles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("storageprofiles").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *storageProfiles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("storageprofiles").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched storageProfile.
func (c *storageProfiles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.StorageProfile, err error) {
	result = &v1beta1.StorageProfile{}
	err = c.client.Patch(pt).
		Resource("storageprofiles").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
