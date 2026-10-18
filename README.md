# shiori (栞)

CLI tool to manage time-decaying personal markdown memos and handovers (`YYYY-MM-DD-HHMM_<slug>.md`) with automated freshness lifecycles (HOT, WARM, COLD).

---

## Installation

### Via `go install`
```bash
go install github.com/AobaIwaki123/shiori/cmd/shiori@latest
```

### From source
```bash
git clone https://github.com/AobaIwaki123/shiori.git
cd shiori
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
shiori list

# Show latest hot memo content
shiori show

# Create a new memo
shiori new "sunday-reflection" -title="Sunday Reflection" -tags="reflection,weekend"

# Auto-decay expired memos (48h -> WARM, 7d -> COLD)
shiori decay

# Print file path for editor integration
code $(shiori path)
```

---

## Environment Variables

- `MEMO_DIR`: Directory containing memos. Defaults to auto-discovering `private/memo` in the current or parent git repositories, or `$HOME/ai-workspace/private/memo`.

---

## Verification

```bash
make verify
```
