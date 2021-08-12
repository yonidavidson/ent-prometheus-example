# ent-prometheus-example
The code for ent prometheus blog post

## Installation

```console
git@github.com:yonidavidson/ent-prometheus-example.git
cd ent-prometheus-example
go run cmd/main.go
```

Then, open [localhost:8080](http://localhost:8080) and query the server
On the second try you'll get an error.

Then, open [localhost:8080/metrics](http://localhost:8080) to see the metrics.

## Running Prometheus


```console
cd ent-prometheus-example/internal/prometheus
docker-compose up
```

then, open [localhost:9090](http://localhost:9090)

## Join the ent Community
In order to contribute to `ent`, see the [CONTRIBUTING](https://github.com/ent/ent/blob/master/CONTRIBUTING.md) file for how to go get started.  
If your company or your product is using `ent`, please let us know by adding yourself to the [ent users page](https://github.com/ent/ent/wiki/ent-users).

## License
This project is licensed under Apache 2.0 as found in the [LICENSE file](LICENSE).
