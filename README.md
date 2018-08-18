# Serverless Emailer

[![Build Status](https://travis-ci.org/serinth/serverless-emailer.svg?branch=master)](https://travis-ci.org/serinth/serverless-emailer)

Serverless implementation of an emailer. It uses a primary emailer e.g. SendGrid then falls back to a secondary emailer.

# TODO
- [x] SendGrid Integration (plain/text only)
- [x] Mailgun Integration (plain/text only)
- [ ] Validate BCC and CC duplicates and return warning
- [ ] Authentication w/ API Key Throttling
- [ ] Analytics
- [ ] Take advantage of SendGrid's personalizations, right now only globally setting subjects and content
- [ ] Store Circuit Breaker state in a persistent store
- [ ] Read config files from S3 instead of file (filesystem wouldn't work as is but code is there for kubernetes type deployments with an accessible file system)
- [ ] Fallback to a queue if primary and secondary emailers fail.

# Environment Variables

Some of these environment variables are actually pulled in through SSM parameter store and decrypted.
See the `serverless.yml` file to see a full list and examples of how they work per *stage*.

|Variable|Description|Default|SSM Parameter|
|---|---|---:|---:|
|REGION|AWS Region|ap-southeast-2|No|
|STAGE|The running environment. Available options are: local, dev and prod|dev|No|
|SENDGRID_URL|The endpoint to call for sending messages. Should only override for environment specific mocks|https://api.sendgrid.com/v3/mail/send|No|
|SENDGRID_API_KEY|The API key issued by SendGrid|-|Yes|
|MAILGUN_URL|The endpoint to call for sending messages. Mailgun has sandbox specific URLs that will only send emails to verified emails unless the domain is verified.|-|No|
|MAILGUN_API_KEY|The API key issued by Mailgun|-|Yes|
|METRICS_COMMAND_NAME|The Hystrix context command name. Hystrix is mostly used for the timeout and code portability. State still needs to be stored to work across lambdas but that's backlogged in TODO.|EmailerAPICall|No|
|HYSTRIX_TIMEOUT|The timeout for the HTTP calls.|20000|No|
|HYSTRIX_MAX_CONCURRENT_REQUESTS|Maximum concurrent requests.|64|No|
|HYSTRIX_ERROR_THRESHOLD|% of errors before the circuit breaker trips|32|No|

# Deployment

## 1. Get an AWS Temporary Session Token:
With MFA enabled on the AWS account, we need to first grab a temporary session token and use it. See [AWS CLI Config Files](https://docs.aws.amazon.com/cli/latest/userguide/cli-config-files.html) on how to set up the credentials file.
```bash
    pip install awsmfa
    awsmfa -i <Profile>
```

## 1b. Optional Log Forwarder

If using the log forwarder, you can first implement and deploy using this template:
[Serverless Log Forwarder Template](https://github.com/serinth/serverless-log-forwarder)

This will aggregate all the lambda function logs and executions into one stream which we can then use with Splunk, Logz IO etc. for metrics and alerts. See the references below for Splunk's lambda forwarder to use with the template above.

## 2. Install Serverless Framework
```bash
    npm install -g serverless
```

## 3. Store API Secrets and Configs in Encrypted AWS SSM Parameter Store
This application uses [mailgun](https://www.mailgun.com) and [SendGrid](https://sendgrid.com) for sending emails.
Sign up and acquire the necessary API Keys.

Securely store the keys and ensure that the region is the same one as where the lambda functions are deployed.
e.g.
```bash
    # replace "dev" with the appropriate environment. See serverless.yml.
    aws --region <REGION> ssm put-parameter --name '/serverless-emailer/thirdparty/sendgrid/dev/apikey' --type "SecureString" --value '<API KEY>'
    aws --region <REGION> ssm put-parameter --name '/serverless-emailer/thirdparty/mailgun/dev/apikey' --type "SecureString" --value '<API KEY>'
```
**Note**
Not having the SSM keys must exist if you're referencing them in the `serverless.yaml` file. Otherwise deployments will fail.

## 4. Build the lambda binaries

This will build Linux ELF binaries with the addition of the configs folder so the proper environment configs can be taken.
**Note** configs in `util/config.go` pull the defaults from the configs/<env>.toml files then if the environment variable is set, it will override the defaults.

```bash
    make build
```

## 5. Run Serverless Deployment

```bash
    serverless deploy -v --aws-profile <Profile>
```

By default the `stage` is set to `dev` if no option is specified. To override it to `local` or `prod` pass it in the serverless deploy command e.g.

```bash
    serverless deploy --stage dev -v --aws-profile <Profile>
```

# Testing The EndPoint

Example simple email below with one recipient. Add more objects of the same type to the array. 
`cc` and `bcc` fields can additionally be added with the same array type as the `to` field.

```bash
curl -H "Content-Type: application/json" -X POST -d '{"to": [{"name":"Tony", "email":"serinth@gmail.com"}], "from":{"name":"FromTony", "email":"truong.tony@live.com"},"subject":"test subject","content":"some content"}' https://<function url>/v1/email/send
```

# Running Locally

Override the file defaults with these required environmental variables. Replace accordingly.
Also modify the `local()` function to implement what to test. Serverless framework does not support Golang lambdas local invocation at the moment.

```bash
    STAGE=local SENDGRID_API_KEY='key' MAILGUN_API_KEY='key' MAILGUN_URL='mysandbox' go run functions/email/*.go
```

# Clean Up

```bash
    serverless remove -v --aws-profile <Profile>
```

# References
[Serverless Framework](https://serverless.com/) - the main framework used for writing and deploying lambda functions

[AWS Amplify](https://github.com/aws/aws-amplify) - for the UI SDK to sign up and sign in users.

[Splunk Log Forwarder Blueprint](https://ap-southeast-2.console.aws.amazon.com/lambda/home?region=ap-southeast-2#/create/new?bp=splunk-logging) - log aggregation

[SendGrid](https://sendgrid.com)

[mailgun](https://www.mailgun.com)