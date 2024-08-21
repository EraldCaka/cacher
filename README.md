# Cacher
Cacher is a lightweight in-memory key-value cache library for Go. It provides a simple interface for storing, retrieving, and managing cached data with optional expiration times. The library is designed for concurrent use, ensuring thread-safety with minimal locking overhead.

## Features
<li><strong>Thread-Safe:</strong> Cacher uses sync.RWMutex to handle concurrent read and write operations safely.
<li><strong>TTL Support:</strong> You can set expiration times for each cache entry. Expired entries are automatically removed.
<li><strong>Basic Operations:</strong> Supports basic cache operations such as Get, Set, Has, and Delete.

## Installation
To install Cacher, use go get:
```shell
go get github.com/EraldCaka/cacher
```