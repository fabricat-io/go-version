# Security Policy

## Reporting Security Issues

We take security seriously and appreciate efforts to responsibly disclose vulnerabilities. If you discover a security issue, please report it **privately** via email or through Github's private vulnerability disclosure reporting mechanism.

- **Contact:** [security@fabricat.io](mailto:security@fabricat.io)
- **Github:** [Private Security Advisory](https://github.com/fabricat-io/semver/security/advisories)
- **Do not create a public issue** on GitHub or any other public channel.

We will acknowledge receipt of your report within **3 business days** and aim to provide an initial response within **7 business days**. Once the issue is resolved, we will coordinate a responsible disclosure timeline with you.

## Supported Versions

We provide security updates for the latest stable version of this project. Older versions **may not** receive security updates unless otherwise specified.

| Version | Supported?                      |
|---------|---------------------------------|
| Latest  | ✅ Yes                           |
| Older   | ❌ No (unless explicitly stated) |

## Security Best Practices

To maintain the security of this project, we follow these best practices:

- **Secure Coding Guidelines**: All contributors must follow secure coding practices and avoid common vulnerabilities such as SQL injection, cross-site scripting (XSS), and remote code execution.
- **Dependency Management**: We regularly update dependencies and monitor known vulnerabilities using automated tools.
- **Secrets Management**: API keys, passwords, and sensitive data **must not** be hardcoded in the repository. Use a secrets manager such as AWS Secrets Manager, HashiCorp Vault, or environment variables.
- **Code Reviews**: Security reviews are part of the code review process to prevent the introduction of vulnerabilities.

## Disclosure Policy

Once a security issue is reported and verified, we will follow this process:

1. **Assess Impact & Risk**: Determine severity and affected versions.
2. **Develop & Test a Fix**: Implement a patch and validate security fixes.
3. **Release a Security Update**: Deploy the fix and notify affected users if necessary.
4. **Coordinate Public Disclosure**: If applicable, disclose the issue responsibly after affected users have had a chance to update.

## Security Testing & Hardening

- **Use static analysis tools** (e.g., `npm audit`, `snyk test`, `bandit`) to detect vulnerabilities.
- **Enable secure configurations** by default (e.g., enforcing HTTPS, disabling weak cipher suites).
- **Regularly review dependencies** for known vulnerabilities.

## Questions?

For any security-related questions, please reach out to **[security@fabricat.io](mailto:security@fabricat.io)**.
