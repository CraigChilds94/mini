# Mini
Lightweight Message Queue written in Go, designed to a very simple, self-hosted solution using Redis as the backend.

## Roadmap
This is a list of things I need to add to this:
 - Dockerize
 - Add / remove items from Redis queues
 - Add removed items to temp "working" queue, with a timeout
 - Add ability to delete items from "working" queue by id
 - Add http interface for fetching queue items
 - Add http interface for marking queue items as processed
 - Add authentication methods (with scopes, producer/consumer/producer-consumer)
 - Clients to interface with Mini
    - Go
    - PHP & Laravel Queue adapter
    - Node
    - Python
