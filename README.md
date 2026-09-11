The Bogus Project
=================

Bogus is an open-source command-line tool created to generate random, collision-free names for git branches.

Ever wonder about Docker's default container names? They’re cool enough to have caught attention during the first hour of playing with Docker. When you create a new Docker container and don’t provide a custom name, Docker generates one for you (like `fervent_lamarr` or `adoring_curie`).

Bogus brings that same naming convention to your personal git code development, enhanced with a unique random suffix so you never run into branch name collisions with existing or historical branches in your repository.

Features
--------

- **Docker-style Memorable Names**: Combines adjectives and famous scientists/pioneers (`<adjective>_<surname>`).
- **Collision Prevention**: Appends a cryptographically secure 6-character alphanumeric suffix (`-[a-z0-9]{6}`) to guarantee branch uniqueness across repository history.
- **Zero External Dependencies**: Built entirely with Go standard libraries and an internal package—no third-party dependencies or external security vulnerabilities.

Installation & Building
-----------------------

Clone the repository and build the binary:

```bash
git clone https://github.com/bldmgr/bogus.git
cd bogus
go build -o bogus .
```

To install it directly to your `$GOPATH/bin`:

```bash
go install .
```

Example Usage
-------------

Run `bogus` within any git repository to create and check out a unique branch:

```bash
$ git status
On branch master
Your branch is up to date with 'origin/master'.

nothing to commit, working tree clean

$ bogus
output | Switched to a new branch 'fervent_lamarr-4k8x2z'

$ git status
On branch fervent_lamarr-4k8x2z
nothing to commit, working tree clean
```

Running Tests
-------------

Run the unit test suite across the project:

```bash
go test -v ./...
```

