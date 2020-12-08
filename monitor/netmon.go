package monitor

type NetData struct {
	NetAdapter string `json:"net_adapter"`
	IPv4       string `json:"i_pv_4"`
	IPv4_mask  string `json:"i_pv_4_mask"`
	IPv6       string `json:"i_pv_6"`
	IPv6_mask  string `json:"i_pv_6_mask"`

	sendrate float32 `json:"sendrate"`
	recvrate float32 `json:"recvrate"`
}

func (this *NetData) Parse() {

}
