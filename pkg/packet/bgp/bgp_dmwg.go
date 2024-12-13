package bgp

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"net/netip"
	"strconv"
	"unsafe"
)

// type PathAttributeAggregatorParam struct {
// 	AS      uint32
// 	Askind  reflect.Kind
// 	Address net.IP
// }

const (
	PathAttributeWireGuardPeerParamLen = int(unsafe.Sizeof(PathAttributeWireGuardPeerParam{}))
)

type PathAttributeWireGuardPeerParam struct {
	EndpointAddress     [16]byte // 16 bytes
	EndpointPort        uint16   // 2 bytes
	PublicKey           [32]byte // 32 bytes
	PersistentKeepalive uint16   // 2 bytes
}

// type PathAttributeAggregator struct {
// 	PathAttribute
// 	Value PathAttributeAggregatorParam
// }

type PathAttributeWireGuardPeer struct {
	PathAttribute
	Value PathAttributeWireGuardPeerParam
}

// func (p *PathAttributeAggregator) DecodeFromBytes(data []byte, options ...*MarshallingOption) error {
// 	value, err := p.PathAttribute.DecodeFromBytes(data, options...)
// 	if err != nil {
// 		return err
// 	}
// 	switch p.Length {
// 	case 6:
// 		p.Value.Askind = reflect.Uint16
// 		p.Value.AS = uint32(binary.BigEndian.Uint16(value[0:2]))
// 		p.Value.Address = value[2:]
// 	case 8:
// 		p.Value.Askind = reflect.Uint32
// 		p.Value.AS = binary.BigEndian.Uint32(value[0:4])
// 		p.Value.Address = value[4:]
// 	default:
// 		eCode := uint8(BGP_ERROR_UPDATE_MESSAGE_ERROR)
// 		eSubCode := uint8(BGP_ERROR_SUB_ATTRIBUTE_LENGTH_ERROR)
// 		return NewMessageError(eCode, eSubCode, nil, "aggregator length isn't correct")
// 	}
// 	return nil
// }

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

// func (p *PathAttributeAggregator) Serialize(options ...*MarshallingOption) ([]byte, error) {
// 	var buf []byte
// 	switch p.Value.Askind {
// 	case reflect.Uint16:
// 		buf = make([]byte, 6)
// 		binary.BigEndian.PutUint16(buf, uint16(p.Value.AS))
// 		copy(buf[2:], p.Value.Address)
// 	case reflect.Uint32:
// 		buf = make([]byte, 8)
// 		binary.BigEndian.PutUint32(buf, p.Value.AS)
// 		copy(buf[4:], p.Value.Address)
// 	}
// 	return p.PathAttribute.Serialize(buf, options...)
// }

func (p *PathAttributeWireGuardPeer) Serialize(options ...*MarshallingOption) ([]byte, error) {
	buf := make([]byte, PathAttributeWireGuardPeerParamLen)

	copy(buf[0:16], p.Value.EndpointAddress[:])

	binary.BigEndian.PutUint16(buf[16:18], p.Value.EndpointPort)

	copy(buf[18:50], p.Value.PublicKey[:])

	binary.BigEndian.PutUint16(buf[50:52], p.Value.PersistentKeepalive)

	return p.PathAttribute.Serialize(buf, options...)
}

// func (p *PathAttributeAggregator) String() string {
// 	return "{Aggregate: {AS: " + strconv.FormatUint(uint64(p.Value.AS), 10) +
// 		", Address: " + p.Value.Address.String() + "}}"
// }

func (p *PathAttributeWireGuardPeer) String() string {
	return "{WireGuardPeer: {EndpointAddress: " + netip.AddrFrom16(p.Value.EndpointAddress).String() +
		", EndpointPort: " + strconv.Itoa(int(p.Value.EndpointPort)) +
		", PublicKey: " + base64.StdEncoding.EncodeToString(p.Value.PublicKey[:]) +
		", PersistentKeepalive: " + strconv.Itoa(int(p.Value.PersistentKeepalive)) + "}}"
}

// func (p *PathAttributeAggregator) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(struct {
// 		Type    BGPAttrType `json:"type"`
// 		AS      uint32      `json:"as"`
// 		Address string      `json:"address"`
// 	}{
// 		Type:    p.GetType(),
// 		AS:      p.Value.AS,
// 		Address: p.Value.Address.String(),
// 	})
// }

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

// func NewPathAttributeAggregator(as interface{}, address string) *PathAttributeAggregator {
// 	v := reflect.ValueOf(as)
// 	asKind := v.Kind()
// 	var l uint16
// 	switch asKind {
// 	case reflect.Uint16:
// 		l = 6
// 	case reflect.Uint32:
// 		l = 8
// 	default:
// 		// Invalid type
// 		return nil
// 	}
// 	t := BGP_ATTR_TYPE_AGGREGATOR
// 	return &PathAttributeAggregator{
// 		PathAttribute: PathAttribute{
// 			Flags:  PathAttrFlags[t],
// 			Type:   t,
// 			Length: l,
// 		},
// 		Value: PathAttributeAggregatorParam{
// 			AS:      uint32(v.Uint()),
// 			Askind:  asKind,
// 			Address: net.ParseIP(address).To4(),
// 		},
// 	}
// }

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
