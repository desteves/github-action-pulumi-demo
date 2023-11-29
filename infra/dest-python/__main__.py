import pulumi
import pulumi_aws as aws

my_bucket = aws.s3.Bucket("my-bucket")
pulumi.export("bucketName", my_bucket.id)
