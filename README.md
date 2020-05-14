# Serverless Scopa

A traditional cards game, serverless style.

## Infrastructure

The AWS infrastructure can be generated through the `simon-says` script that accepts two parameters:

* **command**: `create-stack`|`update-stack`|`delete-stack`
* **environment**: `test`|`prod`

A valid example is the following:

```
./simon-says create-stack test
```

### Please Note

The scripts use by default an AWS profile named `sideprojects`. To create such profile execute:

```
aws configure --profile sideprojects
```

To test it, run:

```
aws s3 ls --profile sideprojects
```