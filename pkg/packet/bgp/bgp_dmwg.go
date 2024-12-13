package bgp

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"net/netip"
	"strconv"
	"unsafe"
)

const (
	PathAttributeWireGuardPeerParamLen = int(unsafe.Sizeof(PathAttributeWireGuardPeerParam{}))
)

type PathAttributeWireGuardPeerParam struct {
	EndpointAddress     [16]byte
	EndpointPort        uint16
	PublicKey           [32]byte
	PersistentKeepalive uint16
}

type PathAttributeWireGuardPeer struct {
	PathAttribute
	Value PathAttributeWireGuardPeerParam
}

func (p *PathAttributeWireGuardPeer) DecodeFromBytes(data []byte, options ...*MarshallingOption) error {
	value, err := p.PathAttribute.DecodeFromBytes(data, options...)
	if err != nil {
		return err
	}
	if int(p.Length) != PathAttributeWireGuardPeerParamLen {
		return NewMessageError(BGP_ERROR_UPDATE_MESSAGE_ERROR, BGP_ERROR_SUB_ATTRIBUTE_LENGTH_ERROR, nil, "wireguard peer length isn't correct")
	}

	addr, _ := netip.AddrFromSlice(value[0:16])
	p.Value.EndpointAddress = addr.As16()

	p.Value.EndpointPort = binary.BigEndian.Uint16(value[16:18])

	copy(p.Value.PublicKey[:], value[18:50])

	p.Value.PersistentKeepalive = binary.BigEndian.Uint16(value[50:52])

	return nil
}

func (p *PathAttributeWireGuardPeer) Serialize(options ...*MarshallingOption) ([]byte, error) {
	buf := make([]byte, PathAttributeWireGuardPeerParamLen)

	copy(buf[0:16], p.Value.EndpointAddress[:])

	binary.BigEndian.PutUint16(buf[16:18], p.Value.EndpointPort)

	copy(buf[18:50], p.Value.PublicKey[:])

	binary.BigEndian.PutUint16(buf[50:52], p.Value.PersistentKeepalive)

	return p.PathAttribute.Serialize(buf, options...)
}

func (p *PathAttributeWireGuardPeer) String() string {
	return "{WireGuardPeer: {EndpointAddress: " + netip.AddrFrom16(p.Value.EndpointAddress).String() +
		", EndpointPort: " + strconv.Itoa(int(p.Value.EndpointPort)) +
		", PublicKey: " + base64.StdEncoding.EncodeToString(p.Value.PublicKey[:]) +
		", PersistentKeepalive: " + strconv.Itoa(int(p.Value.PersistentKeepalive)) + "}}"
}

func (p *PathAttributeWireGuardPeer) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type                BGPAttrType `json:"type"`
		EndpointAddress     string      `json:"endpoint_address"`
		EndpointPort        uint16      `json:"endpoint_port"`
		PublicKey           string      `json:"public_key"`
		PersistentKeepalive uint16      `json:"persistent_keepalive"`
	}{
		Type:                p.GetType(),
		EndpointAddress:     netip.AddrFrom16(p.Value.EndpointAddress).String(),
		EndpointPort:        p.Value.EndpointPort,
		PublicKey:           base64.StdEncoding.EncodeToString(p.Value.PublicKey[:]),
		PersistentKeepalive: p.Value.PersistentKeepalive,
	})
}

func NewPathAttributeWireGuardPeer(addr string, port uint32, publicKey string, persistentKeepalive uint32) *PathAttributeWireGuardPeer {
	t := BGP_ATTR_TYPE_WIREGUARD_PEER

	a, _ := netip.ParseAddr(addr)

	b, _ := base64.StdEncoding.DecodeString(publicKey)
	var pk [32]byte
	copy(pk[:], b)

	return &PathAttributeWireGuardPeer{
		PathAttribute: PathAttribute{
			Flags:  PathAttrFlags[t],
			Type:   t,
			Length: uint16(PathAttributeWireGuardPeerParamLen),
		},
		Value: PathAttributeWireGuardPeerParam{
			EndpointAddress:     a.As16(),
			EndpointPort:        uint16(port),
			PublicKey:           pk,
			PersistentKeepalive: uint16(persistentKeepalive),
		},
	}
}
