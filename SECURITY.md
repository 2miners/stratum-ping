# Security Policy

## Security maintainer

Gleb Shirshov ([@belgracing](https://github.com/belgracing)), 2Miners.

## Supported versions

Only the latest release receives security fixes.

| Version | Supported |
|---------|-----------|
| 1.1.x   | Yes       |
| < 1.1   | No        |

## Reporting a vulnerability

Please do not open a public issue for security problems.

Report privately through GitHub: open the **Security** tab of this repository and click **Report a vulnerability**.

We aim to acknowledge reports within 3 business days and to ship a fix or a mitigation within 30 days, depending on severity. We coordinate disclosure with the reporter and credit them in the advisory unless they prefer to stay anonymous.

## Scope

In scope: the `stratum-ping` source code and the release binaries published on this repository.

Out of scope: 2Miners pool servers and services. Please report issues with them to the 2Miners team directly.

Note: with `-tls` the tool does not verify the server certificate on purpose, because it measures response time and many pools use self-signed certificates. This is expected behaviour, not a vulnerability.
