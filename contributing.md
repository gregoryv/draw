# Contributing

- Write readable code over comments
- Describe intent in your comments, don't rewrite the func signature
  in plain words
- Never lower test coverage, if necessary write a comment why it
  cannot be tested
- Keep changelog up to date
- Don't be afraid of refactoring, coverage keeps you safe

## Daily Routine

    go test ./...

## Release

- Update changelog with version and date, also make it readable for
  the end user
- Tag `git tag -a -m "Tagging v0.5.0" v0.5.0`
- Push `git push --tags`
- Add `[Unreleased]` section to changelog again
