package adblock

import (
    "bufio"
    "os"
    "strings"
)

type AdBlock struct {
    blocked map[string]bool
}

// Load hosts file (adblock list)
func NewAdBlock(path string) (*AdBlock, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    ab := &AdBlock{blocked: make(map[string]bool)}
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()

        // Skip comments
        if strings.HasPrefix(line, "#") {
            continue
        }

        parts := strings.Fields(line)
        if len(parts) < 2 {
            continue
        }

        domain := strings.ToLower(parts[1])
        ab.blocked[domain] = true
    }

    return ab, nil
}

// Check if domain is blocked
func (ab *AdBlock) IsBlocked(domain string) bool {
    domain = strings.ToLower(domain)

    // Exact match
    if ab.blocked[domain] {
        return true
    }

    // Check subdomains
    for blocked := range ab.blocked {
        if strings.HasSuffix(domain, blocked) {
            return true
        }
    }

    return false
}