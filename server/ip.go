package server

func ProcessIPPacket(packet []byte) int {
    if len(packet) < 20 {
        return 0
    }
    return int(packet[9]) // 6 = TCP, 17 = UDP
}