### Using aws cli
Create a user with admin access from AWS IAM

Install aws cli: `pip install awscli`

To configure aws console
`aws configure --profile aws-myaccount`

To list lambda functions
`aws lambda list-functions --profile aws-myaccount`

`aws iam list-roles --profile aws-myaccount`


```
aws lambda create-function --profile aws-myaccount \
--region us-west-2 \
--function-name hello-cli \
--zip-file fileb://index.zip \
--role=arn:aws:iam::<iam_id>:role/lambda_basic_execution \
--handler index.handler \
--runtime nodejs \
--description "this is my first cli command"
```

In the above command we are creating only a lambda function, we haven't
yet configured s3 to push events.

Zipped file can be uploaded to s3 and update command to take file from
S3

Delete a function
```
aws lambda delete-function --function-name hello-cli --profile aws-myaccount
```

Lambda functions are private by default. Only the AWS accounts that
created the account can invoke it or get configuration information from
it.

push model - service invokes lambda functions
pull model - lambda itself pulls for events

To get help
`aws lambda add-permission help`

Get User
`aws iam get-user --profile aws-myaccount`

Resource policy:

```
aws lambda add-permission --profile aws-myaccount \
--function-name pythonObjectMetadata \
--statement-id 1 \
--principal s3.amazonaws.com \
--action lambda:pythonObjectMetadata \
--source-arn arn:aws:s3:::<bucket-name> \
--source-account <account-id>
```

```
aws lambda get-policy --profile aws-myaccount \
--function-name pythonObjectMetadata
```

source-arn, principal changes for other services

giving permission to cross account

get s3 previois notifications
aws s3api get-bucket-notification-configuration --profile aws-myaccount --bucket <bucket-name>

to delete old notification configuration
aws s3api put-bucket-notification-configuration --profile aws-myaccount --bucket <bucket-name> --notification-configuration {}

aws s3api put-bucket-notification-configuration --profile aws-myaccount --generate-cli-skeleton


