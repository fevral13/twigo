# what

A sandbox to implement a mini Twitter clone using Go, Templ and HTMX. My point is to use as few client JS code as possible, avoid using JSON REST.

# roadmap

## MVP

- create messages (text, no media with optional hashes), browse and search messages, hashes, browse users, subscribe to users and hashes 
- add bots to automatically fill the content. I think they will be ChatGPT personas that will post periodically relevant messages. Going to be interesting
- user authentication using Oauth2 providers(Google initially)
- CI/CD

## next

- REST API? Websockets/SSE?