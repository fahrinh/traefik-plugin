# traefik-plugin
Traefik (Go) plugin samples. Enhance Traefik with your custom matcher &amp; middleware (WIP)


## Usage

1. Create a dir where we dump traefik binary, our plugin (`mymatcher.so`) & sample config (`traefik.toml`)
   ```
   mkdir traefik-playground
   cd traefik-playground
   ``` 
2. Build Traefik (based on v1.7.2) with go plugin support

    ```
    git clone -b goplugin-matcher https://github.com/fahrinh/traefik $GOPATH/src/github.com/containous/traefik
    GOGC=off go build github.com/containous/traefik/cmd/traefik
    ```
3. Build custom matcher plugin

    ```
    git clone https://github.com/tiket-libre/traefik-plugin $GOPATH/src/github.com/tiket-libre/traefik-plugin
    go build -buildmode=plugin -o mymatcher.so $GOPATH/src/github.com/tiket-libre/traefik-plugin/matcher/mymatcher.go
    
   # modify mymatcher.so path (plugin.matchers.mymatch.path) in sample-traefik-config.toml
    ```
    
4. Run traefik (with sample config)
   
   ```
   # copy sample config to our playground dir
   cp $GOPATH/src/github.com/tiket-libre/traefik-plugin/traefik.toml .
   
   # (console #1) for the demo purpose, run a simple local http server as a Traefik backend
   python -m http.server
   
   # (console #2) run traefik
   ./traefik
   ```

5. If we hit GET to `http://localhost:7080/` (Traefik http entrypoint) with request body `hello`, http request will be forwarded to python backend.

## Plugin Support

### Matcher
A matcher plugin defines a struct that implements these two methods:
```go
Load() interface{}
MatcherFunc (req *http.Request) bool
```

See [this sample matcher plugin](https://github.com/tiket-libre/traefik-plugin/blob/master/matcher/mymatcher.go) for a full example.

In a frontend rule, use our custom matcher by using `Matcher` keyword which refers to matcher var in the matcher plugin declaration.

```toml
[plugin]
    [plugin.matchers.mymatch]
    path = "./mymatcher.so"

# ...

[frontends]
  [frontends.frontend1]
  backend = "backend1"
    [frontends.frontend1.routes]
      [frontends.frontend1.routes.route0]
      rule = "Matcher: mymatch"
```

## Caveats

- Plugin hot reloading is not yet supported.





