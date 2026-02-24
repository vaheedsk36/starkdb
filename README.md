# MiniDB

A minimal key-value database built from scratch in Go to explore core Computer Science fundamentals and database internals.

## 🚀 Project Goal

Understand how real databases like Redis and PostgreSQL work internally by implementing core concepts step-by-step.

---

## ✅ Features Implemented (Week 1–3)

### 1️⃣ In-Memory Key-Value Store

Supported commands:

- `SET key value`
- `GET key`
- `DELETE key`
- `KEYS`
- `EXIT`

Concepts learned:
- Hash maps
- Command parsing
- Time complexity basics
- REPL design

---

### 2️⃣ Append-Only Log (Durability)

- All write operations appended to `data.log`
- `file.Sync()` used for crash safety
- Log replay on startup

Concepts learned:
- File I/O
- Write-Ahead Logging (WAL) concept
- Crash recovery
- Durability (D in ACID)

---

### 3️⃣ Snapshots & Log Compaction

- `SNAPSHOT` command
- Full state stored in `snapshot.db`
- Log cleared after snapshot
- On startup:
  1. Load snapshot
  2. Replay remaining log

Concepts learned:
- Snapshotting
- Log compaction
- Faster recovery design
- State vs history separation

---

## 🧠 Architecture Overview

### Startup Flow
1. Load snapshot (if exists)
2. Replay append-only log
3. Start REPL

### Write Flow
1. Update in-memory map
2. Append command to log
3. Sync to disk

### Snapshot Flow
1. Serialize full state
2. Save to snapshot file
3. Clear log

---

## 📂 Project Structure

- `main.go` → Core database logic  
- `data.log` → Append-only log  
- `snapshot.db` → Compacted state snapshot  

---

## 🎯 Why This Project?

This project is built for learning:
- Systems design fundamentals
- Storage engine basics
- Database durability mechanics
- Real-world backend internals

---

## 🔮 Next Steps (Future Improvements)

- Transactions (BEGIN / COMMIT / ROLLBACK)
- B-Tree indexing
- Table + schema support
- Concurrency handling
- TCP server interface

---

Built for exploring the beauty of Computer Science ❤️