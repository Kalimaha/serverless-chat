# Serverless Chat ☁️💬

Just another serverless chat built on AWS with Lambda, DynamoDB and Websocket API Gateway.

## Configuration

The scripts use by default an AWS profile. The name of the profile is set in `config/settings/AWS_PROFILE`. To create 
such profile execute:

```
aws configure --profile $AWS_PROFILE
```

To test it, run:

```
aws s3 ls --profile $AWS_PROFILE
```

## Manage the infrastructure

The AWS infrastructure can be managed through the `simon-says` script.

### Create stack

`./simon-says create-stack [test|prod]`

### Update stack

`./simon-says update-stack [test|prod]`

### Delete stack

`./simon-says delete-stack [test|prod]`

### Update Lambda function

`./simon-says update-function [test|prod] [CreateMessage|RegisterUser]`
