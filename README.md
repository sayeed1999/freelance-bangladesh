# freelance-bangladesh

To run the docker containers locally, follow these steps:

- Open up a terminal from ./docker directory

- Allow permission to execute the .sh script files by running

```bash
chmod +x ./start-docker.local.sh
chmod +x ./stop-docker.local.sh
```

- Then run the start script by running

```bash
./start-docker.local.sh
```

- Navigate to `http://localhost:9990` and check keycloak is running properly!!

- To stop the containers, run the stop script by running

```bash
./stop-docker.local.sh
```
