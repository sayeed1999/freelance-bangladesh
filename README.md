# freelance-bangladesh

To run the docker containers locally, follow these steps:

- Open up a terminal from project root directory

- Allow permission to execute the .sh script files by running

```bash
chmod +x ./scripts/setenv.sh
chmod +x ./scripts/start-docker.local.sh
chmod +x ./scripts/stop-docker.local.sh
```

- Run the script for setting up envs

(Note: - a shell cannot modify the env of its parent, using a dot (.) fixes the problem)

```bash
. ./scripts/setenv.sh
```

- Then run the docker start script by running

```bash
./scripts/start-docker.local.sh
```

- Navigate to `http://localhost:8080` and check keycloak is running properly!!

- To stop the containers, run the stop script by running

```bash
./scripts/stop-docker.local.sh
```

## Important Notes

- When registering server in pgadmin, you have to use the postgres port where the docker container is really running, not where it is exposed at.

e.g for `5433:5432` in compose file, you have to use '5432' in port to be able to register server in pgadmin.
