Little distributed systems simulation, I'm working on.

Full hand-written codebase, AIを一つも使わずに.

MIT License.

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
             Postgres         Redis         Storage
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
                             Storage
                                │
                                ▼
                        ┌───────────────┐
                        │ CDN Router    │
                        └───────┬───────┘
                   ┌────────────┼────────────┐
                   ▼            ▼            ▼
              Mumbai Edge   Tokyo Edge   Frankfurt Edge
