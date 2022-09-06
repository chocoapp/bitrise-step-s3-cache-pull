package main

import (
	"log"
	"os"
	"strings"
)

const (
	CACHE_AWS_ACCESS_KEY_ID     = "cache_aws_access_key_id"
	CACHE_AWS_SECRET_ACCESS_KEY = "cache_aws_secret_access_key"
	CACHE_AWS_REGION            = "cache_aws_region"
	CACHE_BUCKET_NAME           = "cache_bucket_name"
	CACHE_RESTORE_KEYS          = "cache_restore_keys"
	CACHE_PATH                  = "cache_path"
)

func parseRestoreKeysInput(keysString string) []string {
	var keys []string
	for _, keyString := range strings.Split(
		strings.TrimSpace(
			keysString,
		),
		"\n",
	) {
		keys = append(keys, strings.TrimSpace(keyString))
	}
	return keys
}

func main() {
	awsAccessKeyId := GetEnvOrExit(CACHE_AWS_ACCESS_KEY_ID)
	awsSecretAccessKey := GetEnvOrExit(CACHE_AWS_SECRET_ACCESS_KEY)
	awsRegion := GetEnvOrExit(CACHE_AWS_REGION)
	bucketName := GetEnvOrExit(CACHE_BUCKET_NAME)
	restoreKeys := GetEnvOrExit(CACHE_RESTORE_KEYS)
	cachePath := GetEnvOrExit(CACHE_PATH)

	s3 := NewAwsS3(
		awsRegion,
		awsAccessKeyId,
		awsSecretAccessKey,
		bucketName,
	)

	keys := parseRestoreKeysInput(restoreKeys)

	for _, key := range keys {
		log.Printf("Checking if cache exists for key '%s'\n", key)
		cacheExists, cacheKey := s3.CacheExists(key)
		if cacheExists {
			log.Println("Cache found! Downloading...")
			size, err := s3.Download(cacheKey, cachePath)

			if err != nil {
				log.Printf("Download failed with error: %s. Cancelling cache restore.\n", err.Error())
				return
			}

			log.Printf("Download was successful, file size: %d.", size)

			return
		}
	}
	log.Println("Cache not found.")

	os.Exit(0)
}
