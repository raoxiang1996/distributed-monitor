package sampled_data

type NetAdapterInfo struct {
	NetAdapterName string `json:"net_adapter"`
	Ipv4           string `json:"i_pv_4"`
	Ipv4Mask       string `json:"i_pv_4_mask"`
	Ipv6           string `json:"i_pv_6"`
	Ipv6Mask       string `json:"i_pv_6_mask"`
}

func (*NetAdapterInfo) GetNetAdapterInfo() {

}

var netTicks int64 = 0
