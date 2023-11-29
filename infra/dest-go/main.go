package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		myBucket, err := s3.NewBucket(ctx, "my-bucket", nil)
		if err != nil {
			return err
		}
		ctx.Export("bucketName", myBucket.ID())
		return nil
	})
}
