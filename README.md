# forkutil
forkutil is a command line tool that is used to with help with different tasks that relate
to GitHub. forkutil is able to search, clone, and fork repositories, submit pull requests, and read readmes from a repository.

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
5. The fork command uses the config.yaml to retreive the personal access token in order to fork repositories
`forkutil fork itsHabib/art`
6. The pullrequest command also uses the config.yaml file for the personal access token
`forkutil pullrequest -d itsHabib/google-home-sms:master -t 'my new pull request' -m 'message for pull request' -s itsHabib:mychanges`
