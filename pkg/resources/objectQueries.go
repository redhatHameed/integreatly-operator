package resources

import (
	"context"

	k8serr "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	pkgclient "sigs.k8s.io/controller-runtime/pkg/client"
)

func Exists(ctx context.Context, serverClient pkgclient.Client, obj runtime.Object) (bool, error) {
	metaobj, err := meta.Accessor(obj)
	if err != nil {
		return false, err
	}
	err = serverClient.Get(ctx, pkgclient.ObjectKey{Name: metaobj.GetName(), Namespace: metaobj.GetNamespace()}, obj)
	if err != nil {
		if k8serr.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
