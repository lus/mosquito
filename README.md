# mosquito
Mosquito is a simple and lightweight HTTP web server to serve and cache static files like assets.

## Deployment
I recommend deploying mosquito using Docker. The image is named `ksebrt/mosquito` and you can find a docker-compose example in the `docker-compose.yml` file.

## Configuration
Configuration is achieved using the following environment variables:
| Environment Variable    | Description                                         | Default value |
|-------------------------|-----------------------------------------------------|---------------|
| `MOS_ADDRESS`           | The address mosquito listens to                     | `:8080`       |
| `MOS_CACHE_DURATION`    | The duration mosquito caches files for              | `10m`         |
| `MOS_DIRECTORY_INDEXES` | Whether or not mosquito generates directory indexes | `false`       |