package subscriptions

import (
    "encoding/json"
    "errors"
    "io"
    "net/http"
    "time"
)

type Node struct {
    Name     string `json:"name"`
    Address  string `json:"address"`
    Region   string `json:"region"`
    LastPing int    `json:"last_ping"`
}

type Subscription struct {
    Name          string `json:"name"`
    Url           string `json:"url"`
    AutoUpdate    bool   `json:"auto_update"`
    UpdateInterval int   `json:"update_interval"` // minutes
    LastUpdate    int64  `json:"last_update"`
    Nodes         []Node `json:"nodes"`
}

type Manager struct {
    Subs []Subscription
}

func NewManager() *Manager {
    return &Manager{}
}

func (m *Manager) Load(path string) error {
    data, err := io.ReadAll(mustOpen(path))
    if err != nil {
        return err
    }
    return json.Unmarshal(data, &m.Subs)
}

func mustOpen(path string) io.Reader {
    f, err := http.Get(path)
    if err != nil {
        panic(err)
    }
    return f.Body
}

func (m *Manager) Save(path string) error {
    data, err := json.MarshalIndent(m.Subs, "", "  ")
    if err != nil {
        return err
    }
    return writeFile(path, data)
}

func writeFile(path string, data []byte) error {
    return errors.New("saving not implemented for mobile")
}

func (m *Manager) UpdateSubscription(i int) error {
    sub := &m.Subs[i]

    resp, err := http.Get(sub.Url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return errors.New("failed to download subscription")
    }

    data, err := io.ReadAll(resp.Body)
    if err != nil {
        return err
    }

    var nodes []Node
    err = json.Unmarshal(data, &nodes)
    if err != nil {
        return err
    }

    sub.Nodes = nodes
    sub.LastUpdate = time.Now().Unix()

    return nil
}

func (m *Manager) AutoUpdate() {
    now := time.Now().Unix()

    for i := range m.Subs {
        sub := &m.Subs[i]

        if !sub.AutoUpdate {
            continue
        }

        if now-sub.LastUpdate >= int64(sub.UpdateInterval*60) {
            m.UpdateSubscription(i)
        }
    }
}

func (m *Manager) AllNodes() []Node {
    var nodes []Node
    for _, s := range m.Subs {
        nodes = append(nodes, s.Nodes...)
    }
    return nodes
}

func (m *Manager) BestNode() *Node {
    nodes := m.AllNodes()
    if len(nodes) == 0 {
        return nil
    }

    best := nodes[0]
    for _, n := range nodes {
        if n.LastPing < best.LastPing {
            best = n
        }
    }
    return &best
}