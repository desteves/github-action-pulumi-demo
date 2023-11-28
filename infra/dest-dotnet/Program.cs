using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var myBucket = new Aws.S3.Bucket("my-bucket");

    return new Dictionary<string, object?>
    {
        ["bucketName"] = myBucket.Id,
    };
});

