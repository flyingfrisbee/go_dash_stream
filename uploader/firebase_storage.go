package uploader

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"github.com/google/uuid"
	"google.golang.org/api/option"
)

type firebaseConn struct {
	bucket *storage.BucketHandle
	ctx    context.Context
}

func (fbc *firebaseConn) uploadFile(movieTitle, episode string, fileInfo *fileInformation) {
	file, err := os.Open(fileInfo.path)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	remotePath := fbc.generateRemoteDirectoryBasedOnExtensionType(fileInfo, movieTitle, episode)
	wr := fbc.bucket.Object(remotePath).NewWriter(fbc.ctx)
	defer wr.Close()

	contentType := fbc.getContentTypeBasedOnFileFormat(fileInfo.path)
	isInvalidContentType := contentType == ""
	if isInvalidContentType {
		errMsg := fmt.Sprintf("cannot determine content type for file: %s", fileInfo.path)
		log.Fatal(errors.New(errMsg))
		return
	}

	wr.ContentType = contentType
	wr.Metadata = map[string]string{
		metadataKey: uuid.NewString(),
	}

	_, err = io.Copy(wr, file)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func (fbc *firebaseConn) generateRemoteDirectoryBasedOnExtensionType(fileInfo *fileInformation, movieTitle, episode string) string {
	switch fileInfo.extType {
	case audio:
		return fmt.Sprintf("%s/%s/audio/%s", movieTitle, episode, fileInfo.name)
	case video:
		return fmt.Sprintf("%s/%s/video/%s", movieTitle, episode, fileInfo.name)
	default:
		return fmt.Sprintf("%s/%s/%s", movieTitle, episode, fileInfo.name)
	}
}

func (fbc *firebaseConn) getContentTypeBasedOnFileFormat(filePath string) string {
	length := len(filePath)
	fileFormat := filePath[(length - 3):]

	switch fileFormat {
	case "mpd", "m4s":
		return manifestAndSegmentContentType
	case "mp4":
		return mp4ContentType
	default:
		return ""
	}
}

var (
	FBConn *firebaseConn
)

func InitFirebaseConn() error {
	conn, err := getConnectionToFirebaseBucket()
	if err != nil {
		return err
	}

	FBConn = conn
	return nil
}

func getConnectionToFirebaseBucket() (*firebaseConn, error) {
	ctx := context.Background()

	config := &firebase.Config{
		StorageBucket: bucketURL,
	}
	opt := option.WithCredentialsFile(privateKeyPath)
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		return nil, err
	}

	storage, err := app.Storage(ctx)
	if err != nil {
		return nil, err
	}

	bucketHandler, err := storage.DefaultBucket()
	if err != nil {
		return nil, err
	}

	conn := &firebaseConn{
		bucket: bucketHandler,
		ctx:    ctx,
	}

	return conn, nil
}
