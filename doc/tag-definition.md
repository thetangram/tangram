# Tag definition


# Component holder tag

```html
<tag data-src=[url] 
     data-name=[component-name] 
     data-timeout=[timeout] 
     data-headers-filter=[string-list] 
     data-cookie-filter=[string-list] 
     data-ttl=[timeout]>
    <!-- Fallback content -->
</tag> 
```

| Attribute          | Optional | Default | description                                   |
| -------------------|:--------:|:--------|:----------------------------------------------|
| data-src           | no       |         | URL of the service.                           |
| data-name          | yes      | root    | The component name from `src` to be included. |
| data-timeout       | yes      | 1s      | The service request timeout.                  |
| data-header-filter | yes      | ""      | List of request header fields to be filtered. |
| data-cookie-filter | yes      | ""      | List of request header fields to be filtered. |
| data-ttl           | yes      | 1h      | Time to live in cache. Default 1 hour.        |


Examples:

A simple use case, with default values and no fallback content

```html
<section data-src=https://github.com/thetangram/sample-components/component1>
</section> 


A more complex use case, defining the timeout, header filters and ttl, and with 
a default fallback

```html
<section data-src=https://github.com/thetangram/sample-components/component1
         data-timeout=1s
         data-headers-filter=[Authorization] 
         data-cookie-filter=[jsessionid] 
         data-ttl=1d>
    <h1>Default content</h1>
    <p>This is the default content, if remote couldn't be retrieved.</p>
</section> 
```


# Component definition tag

<div data-name=[string]>
    <!-- Here the component content -->
</div> 
