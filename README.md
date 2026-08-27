                         ┌─────────────┐
                         │    Client   │
                         │     Code    │
                         └──────┬──────┘
                                │
                                ▼
                         ┌─────────────┐
                         │ API Gateway │
                         │     Go      │
                         └──────┬──────┘
                                │
                  ┌─────────────┼──────────────┐
                  │             │              │
                  ▼             ▼              ▼
             Postgres         Redis          MinIO
                                │              │
                                │              │
                                ▼              ▼
                              Kafka ◄── Upload Event
                                │
                ┌───────────────┼───────────────┐
                │               │               │
                ▼               ▼               ▼
          Video Worker    Thumbnail Worker   Metadata Worker
                │               │               │
                └───────────────┼───────────────┘
                                ▼
                              MinIO
                                │
                                ▼
                        ┌───────────────┐
                        │ CDN Router    │
                        └───────┬───────┘
                   ┌────────────┼────────────┐
                   ▼            ▼            ▼
              Mumbai Edge   Tokyo Edge   Frankfurt Edge