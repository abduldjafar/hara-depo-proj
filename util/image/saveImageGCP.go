package image

import (
	"cloud.google.com/go/storage"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"hara-depo-proj/config"
	"io"
	"log"
	"net/http"
	"os"
)

func getFiles(r *http.Request, key string) string {
	file, handler, err := r.FormFile(key)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	defer file.Close()

	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, _ = io.Copy(f, file)
	return handler.Filename

}
func UploadPhoto(r *http.Request, file, kodeuser, jenis, path string) string {
	fto := getFiles(r, file)
	url := UploadFileGCP(fto, kodeuser+"/"+jenis+"/"+path+"/"+fto)
	var err = os.Remove(fto)
	if err != nil {
		panic(err)
	}
	return url
}

func upload(ctx context.Context, r io.Reader, conf, projectID, bucket, name, path string, public bool) (*storage.ObjectHandle, *storage.ObjectAttrs, error) {

	client, err := storage.NewClient(ctx, option.WithCredentialsFile(conf))
	if err != nil {
		log.Fatal(err)
	}

	bh := client.Bucket(bucket)

	// Next check if the bucket exists
	if _, err = bh.Attrs(ctx); err != nil {
		return nil, nil, err
	}

	obj := bh.Object(name)
	w := obj.NewWriter(ctx)
	if _, err := io.Copy(w, r); err != nil {
		return nil, nil, err
	}
	if err := w.Close(); err != nil {
		return nil, nil, err
	}

	if public {
		if err := obj.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
			return nil, nil, err
		}
	}

	attrs, err := obj.Attrs(ctx)
	return obj, attrs, err
}

func objectURL(objAttrs *storage.ObjectAttrs) string {
	return fmt.Sprintf("https://storage.googleapis.com/%s/%s", objAttrs.Bucket, objAttrs.Name)
}
func UploadFileGCP(filename string, path string) string {
	baseConfig := &config.Configuration{}
	config.GetConfig(baseConfig)

	ctx := context.Background()

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, objAttrs, err := upload(ctx, f, baseConfig.Gcp.Filename, baseConfig.Gcp.Projectid, baseConfig.Gcp.Bucket, path, filename, true)
	if err != nil {
		switch err {
		case storage.ErrBucketNotExist:
			log.Fatal("Please create the bucket first e.g. with `gsutil mb`")
		default:
			log.Fatal(err)
		}
	}

	log.Printf("URL: %s", objectURL(objAttrs))
	log.Printf("Size: %d", objAttrs.Size)
	log.Printf("MD5: %x", objAttrs.MD5)
	log.Printf("objAttrs: %+v", objAttrs)
	return objectURL(objAttrs)
}

// [END storage_quickstart]
