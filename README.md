# tci
Deploys in your terminal. Describe deployment steps in YAML file and run them from terminal using `tci deploy` command.

Features:
- describe deployment steps in YAML file
- enjoy colored output :)

Step types:
- terminal command
- cycles with conditions, retries and delays
- defer commands, which will run in the end of deploy even if it failed

## Install


## Usage
- Create `tci.yml` file. Example:
```yml
steps:
  # run `ls -a`
  - mode: cmd
    command: ls -a

  # default mode is `cmd` so you can omit it
  - command: ls

  # the defer step will be executed in the end of deploy no matter what
  - mode: defer
    command: echo q

  # Run `date +%s` maximum 5 times with 1s delay between retries until output contains `99`
  - mode: until
    command: date +%s
    contains: 99
    retries: 5 # default value is 3
    delay: 1000 # default value is 100
```
- Run `tci deploy` or `tci d` to run `tci.yml` file
- To run another file use `--file` or `-f` flag. Example: `tci -f deploy.yml d`

## Real Example

## Help
- Run `tci` or `tci help`

## Roadmap
- add real example to readme
- release