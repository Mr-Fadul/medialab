package gcp

import "time"

// GCSEvent is the payload of a GCS event.
type GCSEvent struct {
	Bucket                  string    `json:"bucket"`
	ContentType             string    `json:"contentType"`
	Crc32C                  string    `json:"crc32c"`
	Etag                    string    `json:"etag"`
	Generation              string    `json:"generation"`
	ID                      string    `json:"id"`
	Kind                    string    `json:"kind"`
	Md5Hash                 string    `json:"md5Hash"`
	MediaLink               string    `json:"mediaLink"`
	Metageneration          string    `json:"metageneration"`
	Name                    string    `json:"name"`
	SelfLink                string    `json:"selfLink"`
	Size                    string    `json:"size"`
	StorageClass            string    `json:"storageClass"`
	TimeCreated             time.Time `json:"timeCreated"`
	TimeStorageClassUpdated time.Time `json:"timeStorageClassUpdated"`
	Updated                 time.Time `json:"updated"`
}
