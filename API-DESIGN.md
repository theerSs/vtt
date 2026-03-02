# Virtual Table Top

## Tech

- **Go**
- **PostgreSQL**
- **Redis**
- **WebSockets**
- **sqlc**
- **Goose**

## Structure

App will be built using modular monolith architecture where each module will contain its application, domain and communication logic. For now we'll have modules like ***Auth***, ***Tabletop***, ***Rooms***, ***Media***.

Modules will represent domain entities that together will build working VTT.

- **Auth** - Authorizes user and provides session using http only cookie and csrf token to prevent request forgery.
- **Rooms** - Handles room creation, deleting, room specific account creation etc. Mostly metadata related stuff.
- **Media** - Handles media uploads and serving (maps, tokens, handouts). Stores metadata in Postgres and files in object storage / filesystem (implementation detail), provides access control per room.
- **Tabletop** - Handles whole active play such as data syncing between users, tokens moving, rolling etc. Majority of data is stored inside Redis as latest Room snapshot. Every 5 minutes snapshot is saved to Postgres.

### Tabletop submodules

To avoid one giant Tabletop module, it is split internally into submodules (still owned by Tabletop):

- **Gateway (Transport)**  
  WebSocket connection lifecycle, authentication handshake, subscribe/unsubscribe to rooms, message encoding/decoding, ping/pong, backpressure.

- **Application (Commands / Use-cases)**  
  Accepts client intents (e.g. move token, roll dice), validates input, performs rate limiting, checks permissions via Rooms/Auth interfaces, produces domain events.

- **Domain (Events + State Engine)**  
  Canonical room state model, event definitions, reducer that applies events → new state, snapshot schema + versioning/migrations.

- **Realtime (Broadcast / Presence)**  
  Fanout of authoritative events to connected clients, presence/cursors/typing as best-effort streams, reconnection support (seq-based).

- **Persistence (Redis + Postgres)**  
  Redis state store (snapshot, seq, last N events, presence TTL), Postgres checkpointing (every 5 minutes) and optional append-only event log, restore/rebuild logic when a room becomes active.

### Communication between modules

- Modules communicate through explicit interfaces (ports) to keep boundaries clean.
- **Tabletop** depends on **Auth** (identity/session) and **Rooms** (membership/roles/permissions) only through interfaces.
- **Media** enforces access control using room membership/roles (via Rooms/Auth), not by direct DB reads from other modules.
