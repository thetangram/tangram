# The Tangram configuration

All configuration attributes are optional. If not defined the default value wiil be used. 


| Attribute        | Type     | Default | Description                                    |
| -----------------|:--------:|:--------|:-----------------------------------------------|
| address          | string   | ":2018" | The http address the servide is listening.     |
| read timeout     | duration | 200ms   | The total time to load request from client.    |
| write timeout    | duration | 2s      | The total time to response to client.          |
| shutdown timeout | duration | 5s      | Is the time for application graceful shutdown. |
| routes           |          |         | List of routes.                                |
| routes.path      | string   |         | The path to route.                             |
| routes.url       | string   |         | The URL of the target service to route.        |
| routes.timeout   | duration | 1s      | Service request timeout.                       |


## Configuration file

```yaml
address: ":2018"
readTimeout: 200ms
writeTimeout: 2s    
shutdownTimeout: 5s

routes:
- path: /component-1
  url: http://target.com/component-1
  timeout: 1s
- path: /component-2
  url: http://target.com/component-2
  timeout: 2s
```
