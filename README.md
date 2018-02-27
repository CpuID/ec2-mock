*ec2-mock*

[![Build Status](https://travis-ci.org/CpuID/ec2-mock.svg?branch=master)](https://travis-ci.org/CpuID/ec2-mock) [![Coverage Status](https://coveralls.io/repos/github/CpuID/ec2-mock/badge.svg?branch=master)](https://coveralls.io/github/CpuID/ec2-mock?branch=master)

# Summary

Initially this just mocks out the [DescribeTags](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeTags.html) API action.

Other endpoints may follow in future.

Motivation for this was testing "infrastructure as code" builds (eg. Packer image provisioners) that talk to the EC2 API and fetch tags to perform other tasks, in a locally mocked environment.

# Usage

Any environment variables named with a prefix of `TAG__` will be propagated through, with the environment naming syntax as follows:

`TAG__instanceid__tagname`

With instanceid in the format of `i_asdfasdf` (`i-asdfasdf` with `-` substituted for `_`).

Example:

```
docker run -it -e "TAG__i_asdfasdf__BLAH=asdf" -e "TAG__i_aaaabbbb__aaaa=zzzz" -p 33333:33333 --rm cpuid/ec2-mock:latest
```

Then, to query:

```
$ aws --endpoint http://localhost:33333 ec2 describe-tags [--filters "..."]
{
    "Tags": [
        {
            "Key": "BLAH",
            "ResourceId": "i-asdfasdf",
            "ResourceType": "instance",
            "Value": "asdf"
        },
        {
            "Key": "aaaa",
            "ResourceId": "i-aaaabbbb",
            "ResourceType": "instance",
            "Value": "zzzz"
        }
    ]
}
```

# Docker Images

Git SHA tagged builds on the [Docker Hub](https://hub.docker.com/r/cpuid/ec2-mock/tags/)

# License

MIT
