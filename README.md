The Bogus Project
================

Bogus is an open-source project created to generates random names for github branches.

Ever wonder about the Docker default container names. They’re cool enough to have caught my attention during the first hour of playing with Docker.  When you create a new Docker container and don’t give a custom name, Docker generates a name for you. Very cool and smart.

Of course, I wanted to know how Docker generates these container names because I wanted to create random names for my personal code development. I found the code in their GitHub repository – and learned that they use Go!

**Example**:

    $ git clone https://github.com/bldmgr/bogus.git
    $ git status
    On branch master
    Your branch is up to date with 'origin/master'.

    nothing to commit, working tree clean
    $ bogus
    $ git status
    On branch fervent_lamarr
    nothing to commit, working tree clean
