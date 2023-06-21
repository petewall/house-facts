# House Facts

This system is about gathering house facts and offering them in structured way.

## Architecture

Webserver:
* `GET /` - Return all facts as JSON
* `GET /metrics` - Return facts as Prometheus-style metrics

In the future, more methods for organizing, changing, adding, and removing facts

## TODO

* Facts file path should come from an environment variable
* Set a file watch on the facts file (to set it up to work with a ConfigMap)
* Build a docker image
* Add testing
