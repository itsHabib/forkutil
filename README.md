# forkutil
forkutil is a command line tool that is used to with help with different tasks that relate
to GitHub. forkutil is able to search repositories by keyword and more, fork repositories, clone repositories, submit pull requests, and read readmes from repositories.

## Install 
```
go get github.com/itsHabib/forkutil
```
## Usage
1. Getting Help
`forkutil [command] --help` or `forkutil --help`
2. Searching Repositories
`forkutil search topic:go`
3. Getting Readmes
`forkutil docs itsHabib/google-home-sms`
4. To clone repositories the package uses viper and a config.yaml file to 
determine where to clone the repository to. By default it is cloned to the 
HOME env variable.
`forkutil clone itsHabib/forkutil`
5. Forking repositories also uses the config.yaml file for the personal access token
`forkutil fork itsHabib/art`
6. Pull Request also uses the config.yaml file for the personal access token
`forkutil pullrequest -d itsHabib/google-home-sms:master -t 'my new pull request' -m 'message for pull request' -s itsHabib:mychanges`