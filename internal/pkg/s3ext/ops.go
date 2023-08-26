package s3ext

import (
	"context"
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// CreateBucket creates a bucket
// ex:
//
//	&s3.CreateBucketInput{
//	 Bucket: aws.String(bucketName),
//	 CreateBucketConfiguration: &types.CreateBucketConfiguration{
//	 	LocationConstraint: types.BucketLocationConstraint(region),
//	 },
func CreateBucket(params *s3.CreateBucketInput, client *s3.Client, ctx context.Context, optFns ...func(*s3.Options)) (*s3.CreateBucketOutput, error) {
	// Create the S3 Bucket
	return client.CreateBucket(ctx, params, optFns...)
}

// GetBucket determines whether we have this bucket
func GetBucket(bucketName string, client *s3.Client, ctx context.Context, optFns ...func(*s3.Options)) (*s3.HeadBucketOutput, error) {
	// Do we have this Bucket
	return client.HeadBucket(ctx, &s3.HeadBucketInput{Bucket: aws.String(bucketName)}, optFns...)
}

// PutObject puts object to s3
//
//	&s3.PutObjectInput{
//	 	Bucket: aws.String(bucketName),
//	 	Key:    aws.String(fileName),
//	 	Body:   file,
//	}
func PutObject(params *s3.PutObjectInput, client *s3.Client, ctx context.Context, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
	return client.PutObject(ctx, params, optFns...)
}

// PutObjectWG puts object to s3 and uses a waitgroup
//
//	&s3.PutObjectInput{
//	 	Bucket: aws.String(bucketName),
//	 	Key:    aws.String(fileName),
//	 	Body:   file,
//	}
func PutObjectWG(params *s3.PutObjectInput, wg *sync.WaitGroup, client *s3.Client, ctx context.Context, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
	defer func() {
		wg.Done()
	}()
	return client.PutObject(ctx, params, optFns...)
}

// PutObjectPresigned puts presigned object to s3
//
//	&s3.PutObjectInput{
//	 	Bucket: aws.String(bucketName),
//	 	Key:    aws.String(fileName),
//	 	Body:   file,
//	}
func PutObjectPresigned(params *s3.PutObjectInput, client *s3.PresignClient, ctx context.Context, optFns ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error) {
	return client.PresignPutObject(ctx, params, optFns...)
}

// GetObject retrieves object from s3
// ex:
//
//	&s3.GetObjectInput{
//			Bucket: aws.String(bucketName),
//			Key:    aws.String(objectKey),
//	}
func GetObject(params *s3.GetObjectInput, client *s3.Client, ctx context.Context, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
	return client.GetObject(ctx, params, optFns...)
}

// GetObjectPresigned retrieves signed url for object from s3
// ex:
//
//	optFns: func(opts *s3.PresignOptions) {
//				opts.Expires =   timeext.Duration(6 * int64(time.Second))
//			})
func GetObjectPresigned(params *s3.GetObjectInput, client *s3.PresignClient, ctx context.Context, optFns ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error) {
	return client.PresignGetObject(ctx, params, optFns...)
}

// BuildObjectURL builds object url
func BuildObjectURL(region, bucketName, objectKey string) string {
	// https://s3.<region>.amazonaws.com/<bucket-name>/<key>
	// https://bucket-name.s3.region-code.amazonaws.com/key-name
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucketName, region, objectKey)
}

// BuildObjectURLPathStyle builds object url in url path style
func BuildObjectURLPathStyle(region, bucketName, objectKey string) string {
	// https://<bucket-name>.s3<region>.amazonaws.com/<key>
	// https://bucket-name.s3.region-code.amazonaws.com/key-name
	return fmt.Sprintf("https://s3.%s.amazonaws.com/%s/%s", region, bucketName, objectKey)
}
