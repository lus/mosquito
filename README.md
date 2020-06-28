# mosquito
Mosquito is a simple and lightweight HTTP web server to serve and cache static files like assets.

## Deployment
I recommend deploying mosquito using Docker. The image is named `ksebrt/mosquito` and you can find a docker-compose example in the `docker-compose.yml` file.

## Configuration
Configuration is achieved using the following environment variables:
| Environment Variable    | Description                                         | Default value | Example                 |
|-------------------------|-----------------------------------------------------|---------------|-------------------------|
| `MOS_ADDRESS`           | The address mosquito listens to                     | `:8080`       | `127.0.0.1:2255`        |
| `MOS_CACHE_DURATION`    | The duration mosquito caches files for              | `10m`         | `20m10s`                |
| `MOS_DIRECTORY_INDEXES` | Whether or not mosquito generates directory indexes | `false`       | `true`                  |
| `MOS_INDEX_FILES`       | A list of index file names, separated by `;;`       | `<none>`      | `index.html;;index.htm` |