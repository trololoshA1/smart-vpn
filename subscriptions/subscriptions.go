package subscriptions

import (
    "encoding/json"
    "os"
    "sort"
    "strings"
    "time"

    "smart-vpn/core"
)

type Subscription struct {
    Name     string `json:"name"`
    Address  string `json:"address"`  // example: "1.2.3.4:443"
    Region   string `json:"region"`   // RU, EU, US, ASIA
    LastPing int64  `json:"last_ping"` // ms
    Alive    bool   `json:"alive"`
}

type Manager struct {
    Subs []Subscription
}

func NewManager() *Manager {
    return &Manager{Subs: []Subscription{}}
}

// Load subscriptions from file
func (m *Manager) Load(path string) error {
    data, err := os.ReadFile(path)
    if err != nil {
        return err
    }
    return json.Unmarshal(data, &m.Subs)
}

// Save subscriptions to file
func (m *Manager) Save(path string) error {
    data, err := json.MarshalIndent(m.Subs, "", "  ")
    if err != nil {
        return err
    }
    return os.WriteFile(path, data, 0644)
}

// Add new subscription
func (m *Manager) Add(name, address, region string) {
    m.Subs = append(m.Subs, Subscription{
        Name:    name,
        Address: address,
        Region:  strings.ToUpper(region),
        Alive:   false,
        LastPing: -1,
    })
}

// Remove subscription by region
func (m *Manager) RemoveRegion(region string) {
    region = strings.ToUpper(region)
    filtered := []Subscription{}
    for _, s := range m.Subs {
        if s.Region != region {
            filtered = append(filtered, s)
        }
    }
    m.Subs = filtered
}

// Check all subscriptions via TCP ping
func (m *Manager) CheckAll() {
    for i := range m.Subs {
        ping := core.TcpPing(m.Subs[i].Address)
        if ping < 0 {
            m.Subs[i].Alive = false
            m.Subs[i].LastPing = -1
        } else {
            m.Subs[i].Alive = true
            m.Subs[i].LastPing = ping.Milliseconds()
        }
    }

    // Sort by ping (best first)
    sort.Slice(m.Subs, func(i, j int) bool {
        return m.Subs[i].LastPing < m.Subs[j].LastPing
    })
}

// Get only alive subscriptions
func (m *Manager) Alive() []Subscription {
    alive := []Subscription{}
    for _, s := range m.Subs {
        if s.Alive {
            alive = append(alive, s)
        }
    }
    return alive
}

// Get best subscription (lowest ping)
func (m *Manager) Best() *Subscription {
    for _, s := range m.Subs {
        if s.Alive {
            return &s
        }
    }
    return nil
}