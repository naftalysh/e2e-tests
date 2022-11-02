export DOCKER_CONFIG_JSON=$MY_DOCKER_CONFIG_JSON
export GITHUB_TOKEN=$MY_GITHUB_TOKEN
# export GITHUB_E2E_ORGANIZATION=app-studio-test
export GITHUB_E2E_ORGANIZATION=hacbs-release-tests


if [ -z ${DOCKER_CONFIG_JSON+x} ]; then echo "env DOCKER_CONFIG_JSON need to be defined"; exit 1;  else echo "DOCKER_CONFIG_JSON is set"; fi


# go run loadtest.go --username testuser --users 50 --batch 10 -w && ./clear.sh
go run loadtest.go --username testuser --users 10 --batch 5

