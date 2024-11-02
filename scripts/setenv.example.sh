## commented out envs are not necessary to run all containers in a single host server

## keycloak envs
export KEYCLOAK_DATABASE_USER=xxx
export KEYCLOAK_DATABASE_PASSWORD=xxx
export KEYCLOAK_DATABASE_NAME=keycloak-db
export KEYCLOAK_ADMIN=xxx
export KEYCLOAK_ADMIN_PASSWORD=xxx

## database envs
export PGADMIN_EMAIL=xxx
export PGADMIN_PASSWORD=xxx

# export API_KEYCLOAK_BASEURL=http://127.0.0.1:8080 # only for running locally
# export API_KEYCLOAK_REALM=freelance-bangladesh # only for running locally
export API_KEYCLOAK_RESTAPI_CLIENTID=backend-api
export API_KEYCLOAK_RESTAPI_CLIENTSECRET=xxx

# export API_DATABASE_HOST=127.0.0.1 # only for running locally
# export API_DATABASE_PORT=5432 # only for running locally
export API_DATABASE_USER=admin
export API_DATABASE_PASSWORD=xxx
export API_DATABASE_NAME=freelance-db

## frontend envs
# export API_URL=http://localhost:5000 # only for running locally
# export NEXTAUTH_URL=http://localhost:3000 # only for running locally
export NEXTAUTH_SECRET=any-random-long-string-123
export KEYCLOAK_CLIENT_ID=frontend-client
export KEYCLOAK_CLIENT_SECRET=JJxzZNg0aPNDbftzYHu7bjAsWeN6Mbov
# export KEYCLOAK_ISSUER=http://localhost:8080/realms/freelance-bangladesh # only for running locally
