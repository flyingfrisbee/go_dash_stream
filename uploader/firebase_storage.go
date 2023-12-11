package uploader

import (
	env "GithubRepository/go_dash_stream/environment"
	"context"
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

func (fbc *firebaseConn) uploadFile(fileInfo *fileInformation) {
	file, err := os.Open(fileInfo.path)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	remotePath := fbc.generateRemotePath(fileInfo)
	wr := fbc.bucket.Object(remotePath).NewWriter(fbc.ctx)
	defer wr.Close()

	contentType := fbc.getContentTypeBasedOnFileFormat(fileInfo.path)

	wr.ContentType = contentType
	wr.Metadata = map[string]string{
		env.MetadataKey: uuid.NewString(),
	}

	_, err = io.Copy(wr, file)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func (fbc *firebaseConn) generateRemotePath(fileInfo *fileInformation) string {
	return fmt.Sprintf("%s/%s", env.FirebaseDir, fileInfo.name)
}

func (fbc *firebaseConn) getContentTypeBasedOnFileFormat(filePath string) string {
	length := len(filePath)
	fileFormat := filePath[(length - 3):]

	switch fileFormat {
	case "vtt":
		return env.SubtitleExt
	case "mp4":
		return env.Mp4Ext
	default:
		return env.GenericExt
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
		StorageBucket: env.BucketURL,
	}
	opt := option.WithCredentialsFile(env.FirebaseJSONPath)
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
