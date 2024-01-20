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

// VolumeImportSourcesGetter has a method to return a VolumeImportSourceInterface.
// A group's client should implement this interface.
type VolumeImportSourcesGetter interface {
	VolumeImportSources(namespace string) VolumeImportSourceInterface
}

// VolumeImportSourceInterface has methods to work with VolumeImportSource resources.
type VolumeImportSourceInterface interface {
	Create(ctx context.Context, volumeImportSource *v1beta1.VolumeImportSource, opts v1.CreateOptions) (*v1beta1.VolumeImportSource, error)
	Update(ctx context.Context, volumeImportSource *v1beta1.VolumeImportSource, opts v1.UpdateOptions) (*v1beta1.VolumeImportSource, error)
	UpdateStatus(ctx context.Context, volumeImportSource *v1beta1.VolumeImportSource, opts v1.UpdateOptions) (*v1beta1.VolumeImportSource, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.VolumeImportSource, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.VolumeImportSourceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VolumeImportSource, err error)
	VolumeImportSourceExpansion
}

// volumeImportSources implements VolumeImportSourceInterface
type volumeImportSources struct {
	client rest.Interface
	ns     string
}

// newVolumeImportSources returns a VolumeImportSources
func newVolumeImportSources(c *CdiV1beta1Client, namespace string) *volumeImportSources {
	return &volumeImportSources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the volumeImportSource, and returns the corresponding volumeImportSource object, and an error if there is any.
func (c *volumeImportSources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.VolumeImportSource, err error) {
	result = &v1beta1.VolumeImportSource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("volumeimportsources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VolumeImportSources that match those selectors.
func (c *volumeImportSources) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VolumeImportSourceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.VolumeImportSourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("volumeimportsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested volumeImportSources.
func (c *volumeImportSources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("volumeimportsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a volumeImportSource and creates it.  Returns the server's representation of the volumeImportSource, and an error, if there is any.
func (c *volumeImportSources) Create(ctx context.Context, volumeImportSource *v1beta1.VolumeImportSource, opts v1.CreateOptions) (result *v1beta1.VolumeImportSource, err error) {
	result = &v1beta1.VolumeImportSource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("volumeimportsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(volumeImportSource).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a volumeImportSource and updates it. Returns the server's representation of the volumeImportSource, and an error, if there is any.
func (c *volumeImportSources) Update(ctx context.Context, volumeImportSource *v1beta1.VolumeImportSource, opts v1.UpdateOptions) (result *v1beta1.VolumeImportSource, err error) {
	result = &v1beta1.VolumeImportSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("volumeimportsources").
		Name(volumeImportSource.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(volumeImportSource).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *volumeImportSources) UpdateStatus(ctx context.Context, volumeImportSource *v1beta1.VolumeImportSource, opts v1.UpdateOptions) (result *v1beta1.VolumeImportSource, err error) {
	result = &v1beta1.VolumeImportSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("volumeimportsources").
		Name(volumeImportSource.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(volumeImportSource).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the volumeImportSource and deletes it. Returns an error if one occurs.
func (c *volumeImportSources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("volumeimportsources").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *volumeImportSources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("volumeimportsources").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched volumeImportSource.
func (c *volumeImportSources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VolumeImportSource, err error) {
	result = &v1beta1.VolumeImportSource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("volumeimportsources").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
