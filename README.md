# golang-url-shortner
Simple url shortner made by go, just like tinyUrl or bitly


## What it does
This allows users to input a long URL and they get a shortened version of it


```mermaid
flowchart LR
    A["www.quiteverryverylongurlservice.com/long-path"]
    B["URL Shortener Service"]
    C["www.shorturlservice.com/id"]

    A -->|"Long URL"| B
    B -->|"Short URL"| C

    style B fill:#990000,stroke:#aab8c2,stroke-width:2px,color:#ffffff
```


## Key project features:
- Generates a unique short url id for every long url id
- Stores the mappings of the shorturls -> longurls
- Ridirects users from the short url to the original long url


## Testing:
curl -X  POST -d "url=https://www.google.com" http://localhost:8080/shorten