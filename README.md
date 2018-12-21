# Kif

[![Build Status](https://travis-ci.org/pierreboissinot/kif.svg?branch=master)](https://travis-ci.org/pierreboissinot/kif)

## Usage
### Import a Wrike task into Gitlab issue
```
kif import https://www.wrike.com/open.htm\?id\=83740872A
```
[demo](https://asciinema.org/a/7j5KmkLIGYPK0Kp4CmgVhndKm)

## Install

[See releases page](https://github.com/pierreboissinot/kif/releases)


### Build from source

```
git clone git@github.com:pierreboissinot/kif.git
cd kif
go build
```

## Configuration
Init configuration with `kif init`

You can edit `~/.kif.toml` like:
```toml
wrikeApiToken="my.wrike.token"
gitlabApiToken="my.gitlab.token"
```

## Motivation

At work, we use Wrike to manage tasks. It's good enough to be used by
tech(dev) and non tech people(customer, project manager).
To track code changes, we use Gitlab issues.
When a task arrives in developer Inbox, he creates an associated Gitlab issue.
If the task is well described, he copies paste the title and the description;
if not, the developer has to adapt the title and description to correct
infos or split a Wrike task into multiple issues.

## Credits

- [issues-helper](https://www.clever-cloud.com/blog/features/2018/02/13/issues-helper/)
- [cobra](https://github.com/spf13/cobra)

<p align="center">
  <img alt="Hail" src="media/kif-hail.gif">
</p>
