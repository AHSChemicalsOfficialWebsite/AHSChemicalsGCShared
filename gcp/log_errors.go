package gcp

import (
	"context"
	"log"
	"os"
	"sync"

	"cloud.google.com/go/compute/metadata"
	"cloud.google.com/go/logging"
)

var (
	logClient         *logging.Client
	initGCPLoggerOnce sync.Once
)

func InitLogger(ctx context.Context) {
	initGCPLoggerOnce.Do(func() {
		projectID, err := metadata.ProjectIDWithContext(ctx)
		if err != nil {
			log.Fatalf("Failed to get project ID: %v", err)
		}

		logClient, err = logging.NewClient(ctx, projectID)
		if err != nil {
			log.Fatalf("Failed to create logger: %v", err)
		}
	})
}

func CloseLogger() {
	if logClient != nil {
		_ = logClient.Close()
	}
}

func LogDebug(cloudFnName string, message string) {
	if os.Getenv("ENV") == "DEBUG"{
		log.Printf("DEBUG: %s: %s", cloudFnName, message)
	}
	if logClient == nil {
		return
	}
	logClient.Logger(cloudFnName).Log(logging.Entry{Payload: message, Severity: logging.Debug})
}

func LogError(cloudFnName string, message string) {
	if os.Getenv("ENV") == "DEBUG"{
		log.Printf("DEBUG: %s: %s", cloudFnName, message)
	}
	if logClient == nil {
		return
	}
	logClient.Logger(cloudFnName).Log(logging.Entry{Payload: message, Severity: logging.Error})
}

func LogInfo(cloudFnName string, message string) {
	if os.Getenv("ENV") == "DEBUG"{
		log.Printf("DEBUG: %s: %s", cloudFnName, message)
	}
	if logClient == nil {
		return
	}
	logClient.Logger(cloudFnName).Log(logging.Entry{Payload: message, Severity: logging.Info})
}

func LogWarning(cloudFnName string, message string) {
	if os.Getenv("ENV") == "DEBUG"{
		log.Printf("DEBUG: %s: %s", cloudFnName, message)
	}
	if logClient == nil {
		return
	}
	logClient.Logger(cloudFnName).Log(logging.Entry{Payload: message, Severity: logging.Warning})
}

func LogCritical(cloudFnName string, message string) {
	if os.Getenv("ENV") == "DEBUG"{
		log.Printf("DEBUG: %s: %s", cloudFnName, message)
	}	
	if logClient == nil {
		return
	}
	logClient.Logger(cloudFnName).Log(logging.Entry{Payload: message, Severity: logging.Critical})
}

func LogNotice(cloudFnName string, message string) {
	if os.Getenv("ENV") == "DEBUG"{
		log.Printf("DEBUG: %s: %s", cloudFnName, message)
	}	
	if logClient == nil {
		return
	}
	logClient.Logger(cloudFnName).Log(logging.Entry{Payload: message, Severity: logging.Notice})
}

func LogEmergency(cloudFnName string, message string) {
	if os.Getenv("ENV") == "DEBUG"{
		log.Printf("DEBUG: %s: %s", cloudFnName, message)
	}	
	if logClient == nil {
		return
	}
	logClient.Logger(cloudFnName).Log(logging.Entry{Payload: message, Severity: logging.Emergency})
}

func LogAlert(cloudFnName string, message string) {
	if os.Getenv("ENV") == "DEBUG"{
		log.Printf("DEBUG: %s: %s", cloudFnName, message)
	}	
	if logClient == nil {
		return
	}
	logClient.Logger(cloudFnName).Log(logging.Entry{Payload: message, Severity: logging.Alert})
}
