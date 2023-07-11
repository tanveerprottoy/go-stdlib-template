package s3pkg

import (
	"fmt"
	"log"
	"mime/multipart"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// CreateBucket creates a bucket
func CreateBucket(bucket string, s3Svc *s3.S3) error {
	// Create the S3 Bucket
	o, err := s3Svc.CreateBucket(&s3.CreateBucketInput{Bucket: aws.String(bucket)})
	fmt.Print(o)
	if err != nil {
		return err
	}
	// Wait until bucket is created before finishing
	err = s3Svc.WaitUntilBucketExists(&s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		return err
	}
	return nil
}

// GetBucket determines whether we have this bucket
func GetBucket(bucket string, s3Svc *s3.S3) error {
	// Do we have this Bucket
	o, err := s3Svc.HeadBucket(&s3.HeadBucketInput{Bucket: aws.String(bucket)})
	fmt.Print(o)
	if err != nil {
		return err
	}
	return nil
}

// UploadObject uploads to s3
func UploadObject(fileName string, file multipart.File, bucket, region string, uploader *s3manager.Uploader, s3Svc *s3.S3) (string, error) {
	// Upload the file's body to S3 bucket as an object with the key being the
	// same as the filename.
	o, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),

		// Can also use the `filepath` standard library package to modify the
		// filename as need for an S3 object key. Such as turning absolute path
		// to a relative path.
		Key: aws.String(fileName),

		// The file to be uploaded. io.ReadSeeker is preferred as the Uploader
		// will be able to optimize memory when uploading large content. io.Reader
		// is supported, but will require buffering of the reader's bytes for
		// each part.
		Body: file,
	})
	fmt.Print(o)
	if err != nil {
		// Print the error and exit.
		log.Println("unable to upload")
		return "", err
	}
	fmt.Printf("Successfully uploaded %q to %q\n", fileName, bucket)
	// https://<region>.amazonaws.com/<bucket-name>/<key>
	url := fmt.Sprintf(
		"https://%s.amazonaws.com/%s/%s",
		region,
		bucket,
		fileName,
	)
	return url, nil
}

// UploadObject uploads to s3, this function can be invoked
// with a goroutine and to sync with waitgroup
func UploadObjectWg(fileName string, file multipart.File, bucket, region string, uploader *s3manager.Uploader, s3Svc *s3.S3, wg *sync.WaitGroup) (string, error) {
	defer func() {
		wg.Done()
	}()
	// Upload the file's body to S3 bucket as an object with the key being the
	// same as the filename.
	o, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),

		// Can also use the `filepath` standard library package to modify the
		// filename as need for an S3 object key. Such as turning absolute path
		// to a relative path.
		Key: aws.String(fileName),

		// The file to be uploaded. io.ReadSeeker is preferred as the Uploader
		// will be able to optimize memory when uploading large content. io.Reader
		// is supported, but will require buffering of the reader's bytes for
		// each part.
		Body: file,
	})
	fmt.Print(o)
	if err != nil {
		// Print the error and exit.
		log.Println("unable to upload")
		return "", err
	}
	fmt.Printf("Successfully uploaded %q to %q\n", fileName, bucket)
	// https://<region>.amazonaws.com/<bucket-name>/<key>
	url := fmt.Sprintf(
		"https://%s.amazonaws.com/%s/%s",
		region,
		bucket,
		fileName,
	)
	return url, nil
}

// PutObject puts object to s3
func PutObject(fileName string, file multipart.File, contentLength int64, contentType, bucket, region string, s3Svc *s3.S3) (string, error) {
	// Upload the file's body to S3 bucket as an object with the key being the
	// same as the filename.
	i := &s3.PutObjectInput{
		Bucket: aws.String(bucket),

		// Can also use the `filepath` standard library package to modify the
		// filename as need for an S3 object key. Such as turning absolute path
		// to a relative path.
		Key: aws.String(fileName),

		// The file to be uploaded. io.ReadSeeker is preferred as the Uploader
		// will be able to optimize memory when uploading large content. io.Reader
		// is supported, but will require buffering of the reader's bytes for
		// each part.
		Body:          file,
		ContentLength: aws.Int64(contentLength),
		ContentType:   aws.String(contentType),
	}
	fmt.Print(i)
	o, err := s3Svc.PutObject(i)
	log.Println(o)
	if err != nil {
		// handle error
		// Print the error and exit.
		log.Println("unable to upload")
		return "", err
	}
	fmt.Printf("Successfully uploaded %q to %q\n", fileName, bucket)
	fmt.Printf("response %s", awsutil.StringValue(o)) 
	// https://<region>.amazonaws.com/<bucket-name>/<key>
	url := fmt.Sprintf(
		"https://%s.amazonaws.com/%s/%s",
		region,
		bucket,
		fileName,
	)
	return url, nil
}

// UploadObject uploads to s3
func GetObject(fileName string, file multipart.File, bucket, region string, s3Svc *s3.S3) (string, error) {
	// create the input
	i := &s3.GetObjectInput{
		
	}
	// get the object from s3
	o, err := s3Svc.GetObject()
		Bucket: aws.String(bucket),

		// Can also use the `filepath` standard library package to modify the
		// filename as need for an S3 object key. Such as turning absolute path
		// to a relative path.
		Key: aws.String(fileName),

		// The file to be uploaded. io.ReadSeeker is preferred as the Uploader
		// will be able to optimize memory when uploading large content. io.Reader
		// is supported, but will require buffering of the reader's bytes for
		// each part.
		Body: file,
	})
	fmt.Print(o)
	if err != nil {
		// Print the error and exit.
		log.Println("unable to upload")
		return "", err
	}
	fmt.Printf("Successfully uploaded %q to %q\n", fileName, bucket)
	// https://<region>.amazonaws.com/<bucket-name>/<key>
	url := fmt.Sprintf(
		"https://%s.amazonaws.com/%s/%s",
		region,
		bucket,
		fileName,
	)
	return url, nil
}
