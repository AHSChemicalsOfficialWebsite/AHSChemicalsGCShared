package firebase

import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
)

func UploadDataToStorage(ctx context.Context, path string, data []byte, metaData map[string]string) error {
	bucket, err := StorageClient.Bucket(StorageBucket)
	if err != nil {
		return err
	}
	writer := bucket.Object(path).NewWriter(ctx)
		
	if metaData != nil {
		writer.Metadata = metaData
	}

	if _, err := writer.Write(data); err != nil {
		return fmt.Errorf("failed to write to Cloud Storage: %w", err)
	}

	if err := writer.Close(); err != nil {
		return fmt.Errorf("failed to close writer: %w", err)
	}

	return nil
}

func DownloadDataFromStorage(ctx context.Context, path string) ([]byte, error) {
	bucket, err := StorageClient.Bucket(StorageBucket)
	if err != nil {
		return nil, err
	}
	reader, err := bucket.Object(path).NewReader(ctx);
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read from Cloud Storage: %w", err)
	}

	return data, nil
}

func GetStorageObjectMetadata(ctx context.Context, path string) (*storage.ObjectAttrs, error) {
    bucket, err := StorageClient.Bucket(StorageBucket)
    if err != nil {
        return nil, err
    }
    attrs, err := bucket.Object(path).Attrs(ctx)
    if err != nil {
        return nil, err
    }
    return attrs, nil
}