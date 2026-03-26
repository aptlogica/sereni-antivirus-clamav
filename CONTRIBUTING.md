# Contributing


Thank you for considering contributing to this project! Please follow these guidelines to help us maintain a high-quality, welcoming, and consistent open-source ecosystem.

---

# Table of Contents
1. [Development setup](#development-setup)
2. [Branch naming](#branch-naming)
3. [Commit conventions](#commit-conventions)
4. [Pull request checklist](#pull-request-checklist)
5. [Testing requirements](#testing-requirements)
6. [Documentation requirements](#documentation-requirements)
7. [Review process](#review-process)
8. [Code of Conduct](#code-of-conduct)
9. [Issue Reporting](#issue-reporting)
10. [Feature Requests](#feature-requests)
11. [Style Guide](#style-guide)
12. [License](#license)

---


## Development setup
- Follow the instructions in the README to set up your environment.
- Use the provided .env.example as a template for environment variables.
- Install all dependencies using `go mod tidy`.
- Run `make build` to build the project.
- Run `make test-coverage` to ensure all tests pass and coverage is reported.
- Use a linter (e.g., golangci-lint) before submitting code.


## Branch naming
- Use `feature/<name>` for new features (e.g., `feature/add-logging`).
- Use `fix/<name>` for bug fixes (e.g., `fix/null-pointer`).
- Use `release/<version>` for release preparation (e.g., `release/v1.0.0`).


## Commit conventions
- Use [Conventional Commits](https://www.conventionalcommits.org/):
  - `feat`: new features
  - `fix`: bug fixes
  - `docs`: documentation changes
  - `refactor`: code refactoring
  - `test`: adding or updating tests
  - `build/ci/chore`: build or CI changes

**Example commit messages:**
- feat: add user authentication
- fix: correct typo in README
- docs: update contributing guidelines


## Pull request checklist
- Ensure all tests pass (`make test-coverage`).
- Lint your code (`golangci-lint run`).
- Update documentation if needed.
- Reference related issues (e.g., `Closes #123`).
- Ensure your branch is up to date with `main`.
- Complete the PR template.


## Testing requirements
- Add or update tests for new/changed code.
- Aim for at least 90% code coverage.
- Use table-driven tests where possible.
- Run tests with race detection: `go test -race ./...`


## Documentation requirements
- Update README and docs for any user-facing changes.
- Add or update ADRs in `docs/adr/` for architectural decisions.


## Review process
- At least one maintainer must approve before merging.
- Address all review comments.
- Be responsive to feedback.

---

## Code of Conduct
This project follows the [Code of Conduct](CODE_OF_CONDUCT.md). Please be respectful and inclusive.

## Issue Reporting
- Search existing issues before opening a new one.
- Provide clear steps to reproduce bugs.
- Include logs, screenshots, or error messages if possible.

## Feature Requests
- Explain the motivation and use case.
- Suggest possible implementation ideas.

## Style Guide
- Follow Go best practices.
- Use `gofmt` for formatting.
- Keep functions small and focused.
- Write clear comments and documentation.

## License
By contributing, you agree that your contributions will be licensed under the project's license.

---

# Example PR Template

## Description
<!-- Please include a summary of the change and which issue is fixed. -->

## Checklist
- [ ] Tests added/updated
- [ ] Documentation updated
- [ ] Ready for review

---

# Example Issue Template

## Bug report
- **Title:**
- **Description:**
- **Steps to reproduce:**
- **Expected behavior:**
- **Actual behavior:**
- **Environment:**

## Feature request
- **Title:**
- **Description:**
- **Motivation:**
- **Alternatives:**

---

# FAQ

**Q: How do I run tests?**
A: Run `make test-coverage` or `go test -race ./...`.

**Q: How do I check code coverage?**
A: Coverage is generated in `coverage/coverage.out`.

**Q: How do I lint my code?**
A: Use `golangci-lint run`.

**Q: Who do I contact for help?**
A: Open an issue or contact a maintainer listed in the README.
