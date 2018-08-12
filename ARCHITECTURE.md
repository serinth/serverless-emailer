# Why Serverless
## Pros
- fast to get started including development time, deployment and clean up
- cost per invocation (I'm no longer on any free tiers on all 3 public cloud providers)
- will automatically scale to number of requests, subject to account limits
- no infrastructure or OS security management
- writing it in Go allows for moving it to Kubernetes later with small container sizes and to keep most of the code. It's easier to assemble from smaller functions than it is to break a monolith

## Cons
- better to use full running VMs (standalone or in a Kubernetes cluster) when functions are of high volume
- response time is subject to lambda warm up times and even then will not run as fast as a Go pod
- packaging size is large (relatively) if the solution grows too big

# Auditing
- Personally identifiable data should not be logged

# Email Regex
There is a trade off between complex regex for input validation and simple ones. The upkeep larger regex patterns and the ever changing domain names is difficult and not necessary. The only real way to test is to just send the email and deal with repeatedly failed ones differently.
