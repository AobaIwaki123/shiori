// Package memo provides data structures, storage routines, and terminal
// formatters for managing time-decaying personal markdown memos.
//
// # Architecture & Time-Decaying Lifecycle
//
// Memos managed by this package follow the Personal Time-Decaying Memo Protocol:
//
//  1. Timestamp-First Naming:
//     Files are named `YYYY-MM-DD-HHMM_<slug>.md` so they sort chronologically
//     in reverse order (newest first).
//
//  2. Lifecycle States (Freshness & Priority):
//     - StatusHot  ("hot"):  0 to 48 hours old. Active, top priority.
//     - StatusWarm ("warm"): 2 to 7 days old. In progress or pending review.
//     - StatusCold ("cold"): > 7 days old. Inactive archive logs.
//     - StatusDone ("done"): Explicitly marked completed.
//
//  3. Frontmatter Specification:
//     Every memo begins with a YAML frontmatter block:
//
//     ---
//     created_at: 2026-10-18T11:22:00+09:00
//     expires_at: 2026-10-20T00:00:00+09:00
//     status: hot
//     title: Memo Title
//     tags: [schedule, handover]
//     ---
//
// # Storage Discovery
//
// The Store type locates the memo directory using the following hierarchy:
//  1. Explicit directory passed via CLI flag `--dir`.
//  2. Environment variable `MEMO_DIR`.
//  3. Relative candidates: `private/memo`, `../private/memo`, `../../private/memo`.
//  4. Git root search upwards from the working directory looking for `private/memo`.
//  5. Default fallback to `$HOME/ai-workspace/private/memo`.
package memo
