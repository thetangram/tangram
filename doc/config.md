# The Tangram configuration

All configuration attributes are optional. If not defined the default value wiil be used. 


| Attribute        | Type     | Default    | Description                                    |
| -----------------|:--------:|:-----------|:-----------------------------------------------|
| address          | string   | ":2018"    | The HTTP server.                               |
| read timeout     | duration | 200ms      | The total time to load incoming request.       |
| write timeout    | duration | 2s         | The total time to response to client.          |
| shutdown timeout | duration | 5s         | Is the time for application graceful shutdown. |
| routes           |          | empty list | List of routes...                              |
| routes.path      | string   | "" (*)     | The path to map.                               |
| routes.url       | string   | "" (*)     | The URL of the target service to route.        |
| routes.timeout   | duration | 1s         | URL request timeout.                           |

(*) If a route is defined this two fields should have value. If no path or url is defined,
the server will not handled this route.


## Configuration file

An example of configuration file, with all fields configured, and two routes.

```yaml
address: ":2018"
read-timeout: 200ms
write-timeout: 2s    
shutdown-timeout: 5s

routes:
- path: /component-1
  url: http://target.com/component-1
  timeout: 1s
- path: /component-2
  url: http://target.com/component-2
  timeout: 2s
```
