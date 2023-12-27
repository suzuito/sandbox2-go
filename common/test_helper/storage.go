package test_helper

import (
	"context"
	"io"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

func DeleteGoogleCloudStorageObjects(ctx context.Context, cli *storage.Client, bucketName string, prefix string) error {
	it := cli.Bucket(bucketName).Objects(ctx, &storage.Query{Prefix: prefix})
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		if err := cli.Bucket(bucketName).Object(attrs.Name).Delete(ctx); err != nil {
			return err
		}
	}
	return nil
}

func GetGoogleCloudStorageObject(ctx context.Context, cli *storage.Client, bucketName, path string) (*storage.ObjectAttrs, []byte, error) {
	oh := cli.Bucket(bucketName).Object(path)
	r, err := oh.NewReader(ctx)
	if err != nil {
		return nil, nil, err
	}
	defer r.Close()
	body, err := io.ReadAll(r)
	if err != nil {
		return nil, nil, err
	}
	attrs, err := oh.Attrs(ctx)
	if err != nil {
		return nil, nil, err
	}
	return attrs, body, nil
}
