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

The AWS infrastructure can be managed through the `simon-says` script that accepts two parameters:

* **command**: `create-stack`|`update-stack`|`delete-stack`
* **environment**: `test`|`prod`

A valid example is the following:

```
./simon-says create-stack test
```
