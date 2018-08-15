# Serverless Emailer

Work In Progress


# TODO
- [x] SendGrid Integration (plain/text only)
- [x] Mailgun Integration (plain/text only)
- [ ] Validate BCC and CC duplicates and return warning
- [ ] Authentication w/ API Key Throttling
- [ ] Analytics
- [ ] A/B Testing, Experimentation and Feature Flagging
- [ ] Take advantage of SendGrid's personalizations, right now only globally setting subjects and content
- [ ] Store Circuit Breaker state in a persistent store


# Deployment

With MFA enabled on the AWS account, we need to first grab a temporary session token and use it. See [AWS CLI Config Files](https://docs.aws.amazon.com/cli/latest/userguide/cli-config-files.html) on how to set up the credentials file.

## 1. - Get an AWS Temporary Session Token:
```bash
    pip install awsmfa
    awsmfa -i <Profile>
```

## 1.a Optional Log Forwarder

If using the log forwarder, you can first implement and deploy using this template:
[Serverless Log Forwarder Template](https://github.com/serinth/serverless-log-forwarder)

This will aggregate all the lambda function logs and executions into one stream which we can then use with Splunk, Logz IO etc. for metrics and alerts. See the references below for Splunk's lambda forwarder to use with the template above.

## 2. Install Serverless Framework
```bash
    npm install -g serverless
```

## 3. Store API Secrets in Encrypted AWS SSM Parameter Store
This application uses [mailgun](https://www.mailgun.com) and [SendGrid](https://sendgrid.com) for sending emails.
Sign up and acquire the necessary API Keys.

e.g.
```bash
    # replace "dev" with the appropriate environment. See serverless.yml.
    aws ssm put-parameter --name '/serverless-emailer/thirdparty/sendgrid/dev/apikey' --type "SecureString" --value '<API KEY>'
    aws ssm put-parameter --name '/serverless-emailer/thirdparty/mailgun/dev/apikey' --type "SecureString" --value '<API KEY>'
```

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

# Testing The EndPoint


# Running Locally

Override the file defaults with these required environmental variables. Replace accordingly.

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