# Set Up
```
minikube start
make init
```

# Tutorial using cli
```
# simple pub/sub
$ nats sub cli.demo
$ nats pub cli.demo "message {{.Count}} @ {{.TimeStamp}}" --count=10

# more interesting we can interact with the wttr.in web service
$ nats reply 'cli.weather.>' --command "curl -s wttr.in/{{2}}?format=3" 
$ nats request "cli.weather.{beijin,tokyo}" ""
23:57:50 Sending request on "cli.weather.{beijin,tokyo}"
23:57:52 Received with rtt 1.5063365s
beijin: ⛅️  +24°C
tokyo: ⛅️  +28°C
```

# Tutorial using nats.go
```
$ pushd hello-world && go mod download && go run main.go && pulld
```