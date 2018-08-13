# Serverless Emailer

Work In Progress


# TODO
- [ ] Validate BCC and CC duplicates and return warning
- [ ] Internationalization (i18n)
- [ ] Authentication
- [ ] Analytics
- [ ] A/B Testing, Experimentation and Feature Flagging
- [ ] Vault Secrets Codified instead of ENV variables
- [ ] Take advantage of SendGrid's personalizations, right now only globally setting subjects and content

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

## 4. Run Serverless Deployment

```bash
    serverless deploy -v --aws-profile TEMPSESSION
```

# Testing The EndPoint


# Running Locally

```bash
    STAGE=local \ 
    SENDGRID_API_KEY='key' \
    MAILGUN_API_KEY='key' \
    go run email.go main.go
```

# Clean Up

```bash
    serverless remove -v --aws-profile TEMPSESSION
```

# References
[Serverless Framework](https://serverless.com/) - the main framework used for writing and deploying lambda functions

[AWS Amplify](https://github.com/aws/aws-amplify) - for the UI SDK to sign up and sign in users.

[Splunk Log Forwarder Blueprint](https://ap-southeast-2.console.aws.amazon.com/lambda/home?region=ap-southeast-2#/create/new?bp=splunk-logging) - log aggregation