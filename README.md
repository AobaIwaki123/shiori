# memo-cli

CLI tool to manage time-decaying personal markdown memos (`YYYY-MM-DD-HHMM_<slug>.md`) with automated freshness lifecycles (HOT, WARM, COLD).

---

## Installation

### Via `go install`
```bash
go install github.com/AobaIwaki123/memo-cli/cmd/memo@latest
```

### From source
```bash
git clone https://github.com/AobaIwaki123/memo-cli.git
cd memo-cli
make install
```

Make sure `$(go env GOPATH)/bin` is in your `$PATH`:
```bash
export PATH="$(go env GOPATH)/bin:$PATH"
```

---

## Quick Start

```bash
# List active memos (HOT and WARM)
memo list

# Show latest hot memo content
memo show

# Create a new memo
memo new "sunday-reflection" -title="Sunday Reflection" -tags="reflection,weekend"

# Auto-decay expired memos (48h -> WARM, 7d -> COLD)
memo decay

# Print file path for editor integration
code $(memo path)
```

---

## Environment Variables

- `MEMO_DIR`: Directory containing memos. Defaults to auto-discovering `private/memo` in the current or parent git repositories, or `$HOME/ai-workspace/private/memo`.

---

## Verification

```bash
make verify
```
