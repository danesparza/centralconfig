# centralconfig [![Circle CI](https://circleci.com/gh/cagedtornado/centralconfig.svg?style=svg)](https://circleci.com/gh/cagedtornado/centralconfig)
A simple REST based service for managing application configuration across a cluster using a SQL back-end.  Runs on Linux/Windows/OSX/FreeBSD/Raspberry Pi.

Back-ends supported:
- [BoltDB](https://github.com/boltdb/bolt) (default)
- [MySQL](https://www.mysql.com/)
- [Microsoft SQL server (MSSQL)](https://www.microsoft.com/en-us/server-cloud/products/sql-server/)

### Quick start
To get up and running, [grab the latest release](https://github.com/danesparza/centralconfig/releases/latest) for your platform

Start the server:
```
centralconfig serve
```
Then visit the url [http://localhost:3000](http://localhost:3000) and you can add/edit your configuration through the built-in web interface.  

If no other configuration is specified, BoltDB will be used to store your config items in a file called 'config.db' in the working directory.

#### Configuration
To customize the config, first generate a default config file (with the name centralconfig.yaml):
```
centralconfig defaults > centralconfig.yaml
```
