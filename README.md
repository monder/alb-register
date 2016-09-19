# alb-register

[![Go Report Card](https://goreportcard.com/badge/github.com/monder/alb-register)](https://goreportcard.com/report/github.com/monder/alb-register)
[![license](https://img.shields.io/github/license/monder/alb-register.svg?maxAge=2592000&style=flat-square)]()
[![GitHub tag](https://img.shields.io/github/tag/monder/alb-register.svg?style=flat-square)]()

An `ACI` tthat register current instance in AWS [Application Load Balancer](https://aws.amazon.com/elasticloadbalancing/applicationloadbalancer/)

## Images
Signed `ACI`s for `linux-arm64` are available at `monder.cc/alb-register` with versions matching git tags.

## Usage

```
rkt run \
   --insecure-options=image \
   docker://app \
   monder.cc/alb-register:v0.1.0 \
   --
   arn:aws:elasticloadbalancing:eu-west-1:123456789:targetgroup/name1/1234567890
   --port=444
   arn:aws:elasticloadbalancing:eu-west-1:123456789:targetgroup/name2/1234567890
   arn:aws:elasticloadbalancing:eu-west-1:123456789:targetgroup/name3/1234567890
   --port=33
```

The script above will launch app and a sidekick in the same pod.
The sidekick will register current instance in ALB in 3 target groups:

- `arn:aws:elasticloadbalancing:eu-west-1:123456789:targetgroup/name1/1234567890` on port `444`.
- `arn:aws:elasticloadbalancing:eu-west-1:123456789:targetgroup/name2/1234567890` on default port for that target group.
- `arn:aws:elasticloadbalancing:eu-west-1:123456789:targetgroup/name3/1234567890` on port `33`.

## License
MIT    
