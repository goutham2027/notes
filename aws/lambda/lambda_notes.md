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
