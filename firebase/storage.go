package firebase

import (
	"context"
	"fmt"
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