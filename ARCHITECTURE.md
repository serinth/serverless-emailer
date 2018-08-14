# Architectural Choices

## Why Serverless
### Pros
- fast to get started including development time, deployment and clean up
- cost per invocation (I'm no longer on any free tiers on all 3 public cloud providers)
- will automatically scale to number of requests, subject to account limits
- no infrastructure or OS security management
- writing it in Go allows for moving it to Kubernetes later with small container sizes and to keep most of the code. It's easier to assemble from smaller functions than it is to break a monolith
- versioning is intrinsic to lambda functions and there are also aliases to point to specific functions. It does not version lock from a consumer perspective.

### Cons
- better to use full running VMs (standalone or in a Kubernetes cluster) when functions are of high volume
- response time is subject to lambda warm up times and even then will not run as fast as a Go pod
- packaging size is large (relatively) if the solution grows too big


## Alternatives
Depending on team skills, size and microservices architectures in place, I could have used: 

1. https://github.com/serinth/generator-go-grpcgw-api GRPC Gateway Go microservice ready for Kubernetes

2. https://github.com/serinth/node-api-boilerplate - there are no db calls and no heavy processing. Multiplexing should be fast.

3. If given extremely heavy loads and systems that require multiple notifications of what happened, I would dump all the messages on a queue in binary format to minimize internal message sizing.
Then use a worker pull pattern so I know how much to scale out the workers based on queue size. All messaging and subscribers will need to implement protocol buffers so that processing can be language agnostic.
If we're really at this point, I wouldn't even use a third party service for emailing.

# Auditing
- Personally identifiable data should not be logged. Messages should be ephermeral, we don't need to see body content.
- If information is stored long term, there needs to be an anonymizing process or a deletion process. It doesn't have to be automated but documents should exist for a process for Australian Privacy laws.
- The above is also true for GDPR compliance for customizers residing in the EU. In which case the data also needs to sit in eu-west-2, eu-west-3 or eu-central-1 clusters.
- As with best practices, being OWASP compliant is a good idea.
- The entire flow from the end user should be traceable with sensitive data masked. A correlation/trace ID should exist for all sessions.
- Log aggregation should be implemented e.g. Splunk, ElasticSearch, Logz.io, Sumo Logic etc. So that monitoring can be implemented for odd behaviour and alerts can be made for critical system failures.
- Support plans for each product should be agreed upon to ensure we do not break any SLAs and have a formal process for rollback procedures

# Additional Nice To Haves Given More Time
The bullet points below plus the TODOs on the README.md

- service discovery for queue and service names
- analytics platform to push statistics and data to which would power;
    - A/B testing to minimize risk when rolling out and testing major changes
    - Experimentation for a subset of the population for fine grained feature testing. Also a risk mitigation exercise.
    - Feature Flagging (can be combined with the above)
- SendGrid and MailGun provide a bounces API which we can use to collect data on

# Domain Verification
- ensure to follow the integration parties' documentation to verify the domain so that DKIM, SPF works and emails are less likely to land in the spam box

# Email Regex
There is a trade off between complex regex for input validation and simple ones. The upkeep larger regex patterns and the ever changing domain names is difficult and not necessary. The only real way to test is to just send the email and deal with repeatedly failed ones differently.
