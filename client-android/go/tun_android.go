package mobile

import (
    "os"
    "golang.org/x/mobile/asset"
)

type AndroidTun struct {
    fd *os.File
}

func NewAndroidTun(fd int) *AndroidTun {
    return &AndroidTun{
        fd: os.NewFile(uintptr(fd), "tun"),
    }
}

func (t *AndroidTun) ReadPacket() ([]byte, error) {
    buf := make([]byte, 4096)
    n, err := t.fd.Read(buf)
    if err != nil {
        return nil, err
    }
    return buf[:n], nil
}

func (t *AndroidTun) WritePacket(packet []byte) error {
    _, err := t.fd.Write(packet)
    return err
}