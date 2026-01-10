## Contributing

Contributions are welcome, and they are greatly appreciated! Every
little bit helps, and credit will always be given. There a just a few small
guidelines you need to follow.


## Submitting a patch

  1. It's generally best to start by opening a new issue describing the bug or
     feature you're intending to fix.  Even if you think it's relatively minor,
     it's helpful to know what people are working on.  Mention in the initial
     issue that you are planning to work on that bug or feature so that it can
     be assigned to you.

  2. Follow the normal process of [forking][] the project, and setup a new
     branch to work in.  It's important that each group of changes be done in
     separate branches in order to ensure that a pull request only includes the
     commits related to that bug or feature.

  3. Go makes it very simple to ensure properly formatted code, so always run
     `go fmt` on your code before committing it.

  4. Any significant changes should almost always be accompanied by tests.  The
     project already has good test coverage, so look at some of the existing
     tests if you're unsure how to go about it.

  5. Do your best to have well-formed commit messages for each change.
     This project follows [Conventional Commits](https://www.conventionalcommits.org/).
     Examples:
     - `feat: add new rule DL3028`
     - `fix: correct port validation in DL3011`
     - `docs: update installation guide`
     - `refactor: simplify analyzer logic`

  6. Finally, push the commits to your fork and submit a [pull request][].

[forking]: https://help.github.com/articles/fork-a-repo
[pull request]: https://help.github.com/articles/creating-a-pull-request
