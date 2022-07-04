package gcp

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/rs/zerolog/log"

	video "cloud.google.com/go/videointelligence/apiv1"
	videopb "google.golang.org/genproto/googleapis/cloud/videointelligence/v1"
)

func HandleVideoIntelligence(ctx context.Context, e GCSEvent) error {

	client, err := video.NewClient(ctx)
	if err != nil {
		log.Error().Msgf("video.NewClient: %s", err.Error())
	}

	defer client.Close()

	inputUri := fmt.Sprintf("gs://%s/%s", e.Bucket, e.Name)

	ext := path.Ext(e.Name)
	resultFile := e.Name[0:len(e.Name)-len(ext)] + ".json"

	bucketResult, exist := os.LookupEnv("BUCKET_RESULT")
	if !exist {
		return fmt.Errorf("BUCKET_RESULT not found")
	}

	outputUri := fmt.Sprintf("gs://%s/%s", bucketResult, resultFile)

	transcriptConfig := &videopb.SpeechTranscriptionConfig{
		LanguageCode: "pt-BR",
	}

	textDetectionConfig := &videopb.TextDetectionConfig{
		LanguageHints: []string{"pt-BR"},
	}

	//personConfig := &videopb.PersonDetectionConfig{
	//	IncludeBoundingBoxes: true,
	//	IncludeAttributes:    true,
	//	IncludePoseLandmarks: true,
	//}

	//faceConfig := &videopb.FaceDetectionConfig{
	//	IncludeBoundingBoxes: true,
	//	IncludeAttributes:    true,
	//}

	op, err := client.AnnotateVideo(ctx, &videopb.AnnotateVideoRequest{
		InputUri: inputUri,
		Features: []videopb.Feature{
			//videopb.Feature_OBJECT_TRACKING,
			videopb.Feature_LABEL_DETECTION,
			videopb.Feature_SHOT_CHANGE_DETECTION,
			videopb.Feature_SPEECH_TRANSCRIPTION,
			videopb.Feature_LOGO_RECOGNITION,
			//videopb.Feature_EXPLICIT_CONTENT_DETECTION,
			videopb.Feature_TEXT_DETECTION,
			//videopb.Feature_FACE_DETECTION,
			//videopb.Feature_PERSON_DETECTION,
		},
		OutputUri: outputUri,
		VideoContext: &videopb.VideoContext{
			SpeechTranscriptionConfig: transcriptConfig,
			TextDetectionConfig:       textDetectionConfig,

			//PersonDetectionConfig:     personConfig,
			//FaceDetectionConfig:       faceConfig,
		},
	})

	if err != nil {
		return fmt.Errorf("AnnotateVideo: %s", err.Error())
	}

	log.Info().Msgf("Job started: %s", e.Name)

	_, err = op.Poll(ctx)
	if err != nil {
		log.Fatal().Msgf("Failed to annotate: %s", err.Error())
	}

	log.Info().Msgf("Job completed: %s", e.Name)

	return nil
}
