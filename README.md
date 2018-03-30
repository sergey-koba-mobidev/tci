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
Use next commands to install latest `tci` on your mac or linux
```
curl -L https://github.com/sergey-koba-mobidev/tci/releases/download/latest/tci-`uname -s`-`uname -m` -o /usr/local/bin/tci
chmod +x /usr/local/bin/tci
```

Change `latest` to a specific version to install it.

## Usage
- Create `tci.yml` file. Let's consider deploying stack to Docker Swarm as Example
```yml
steps:
  # Perform command
  - command: docker service create --name registry --publish published=5000,target=5000 registry:2

  # A defer command will be executed in the end of deploy script eve if some step failed
  - command: docker service rm registry
    mode: defer

  - command: docker-compose -f docker-compose.prod.yml build
  - command: docker-compose -f docker-compose.prod.yml push
  - command: docker stack deploy --compose-file docker-compose.prod.yml --resolve-image=never MYSTACK

  # Wait until db service is up. Check output of command every 5 seconds to contain 1/1. Retry 36 times ~ 3 minutes
  - command: docker service ls | grep db | awk '/ / { print $4 }'
    mode: until
    contains: 1/1
    retries: 36
    delay: 5000
    shell: true

  # Execute command in a shell to allow substitution
  - command: docker exec $(docker ps -q -f name=MYSTACK_db) /database/bin/db_migrate.sh
    shell: true

  - command: docker stack services MYSTACK
```
- Run `tci deploy` or `tci d` to run `tci.yml` file
- To run another file use `--file` or `-f` flag. Example: `tci -f deploy.yml d`

## Environment variables
Don't forget to set required env vars before running `tci d`.
If you try to run any `export` s via tci they won't persist.

## Help
- Run `tci` or `tci help`

## Roadmap
- tests