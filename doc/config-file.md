# The Tangram configuration

All configuration attributes are optional. If not defined the default value wiil be used. 


| Attribute              | Type     | Default | Description                                   |
| -----------------------|:--------:|:--------|:----------------------------------------------|
| address                | string   | ":2018" | The http address the servide is listening.    |
| system shutdownTimeout | duration | 5s      | Is the time for graceful shutdown.            |
| http readTimeout       | duration | 200ms   | The total time to load request from client.   |
| http writeTimeout      | duration | 2s      | The total time to response to client.         |


## Configuration file

```yaml
address: ":2018"
system:
    shutdownTimeout: 10s
http:
    readTimeout: 200ms
    writeTimeout: 2s    
```
