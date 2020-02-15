
Start first server npx http-server -p 5001 &
npx http-server -p 5002 &
npx http-server -p 5003 &
go run main.go &
watch curl localhost:3441

# miniLoadBalancer
miniLoadBalancer is sample for understand loadbalancer mechanism

based on https://github.com/ahmetb youtube channel video 

## Quick Start

```bash
git clone https://github.com/ysBayram/miniLoadBalancer.git
cd ./miniLoadBalancer
go run main.go
```

LB running on http://localhost:3441

## Testing app ;

First you need servers like behind LB we can simulate with http-server(npm package)

Start first server 
```bash
npx http-server -p 5001
```
and scond and third (for understand strategy like round robin)
```bash
npx http-server -p 5002
npx http-server -p 5003
```

Then start request like client 
```bash
curl localhost:3441
```

The Result is;

counter = 1 backend = localhost:5001
counter = 2 backend = localhost:5002
counter = 3 backend = localhost:5003
counter = 4 backend = localhost:5001


