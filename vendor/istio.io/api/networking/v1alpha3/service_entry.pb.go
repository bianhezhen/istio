// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: networking/v1alpha3/service_entry.proto

package v1alpha3

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Location specifies whether the service is part of Istio mesh or
// outside the mesh.  Location determines the behavior of several
// features, such as service-to-service mTLS authentication, policy
// enforcement, etc. When communicating with services outside the mesh,
// Istio's mTLS authentication is disabled, and policy enforcement is
// performed on the client-side as opposed to server-side.
type ServiceEntry_Location int32

const (
	// Signifies that the service is external to the mesh. Typically used
	// to indicate external services consumed through APIs.
	ServiceEntry_MESH_EXTERNAL ServiceEntry_Location = 0
	// Signifies that the service is part of the mesh. Typically used to
	// indicate services added explicitly as part of expanding the service
	// mesh to include unmanaged infrastructure (e.g., VMs added to a
	// Kubernetes based service mesh).
	ServiceEntry_MESH_INTERNAL ServiceEntry_Location = 1
)

var ServiceEntry_Location_name = map[int32]string{
	0: "MESH_EXTERNAL",
	1: "MESH_INTERNAL",
}
var ServiceEntry_Location_value = map[string]int32{
	"MESH_EXTERNAL": 0,
	"MESH_INTERNAL": 1,
}

func (x ServiceEntry_Location) String() string {
	return proto.EnumName(ServiceEntry_Location_name, int32(x))
}
func (ServiceEntry_Location) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorServiceEntry, []int{0, 0}
}

// Resolution determines how the proxy will resolve the IP addresses of
// the network endpoints associated with the service, so that it can
// route to one of them. The resolution mode specified here has no impact
// on how the application resolves the IP address associated with the
// service. The application may still have to use DNS to resolve the
// service to an IP so that the outbound traffic can be captured by the
// Proxy. Alternatively, for HTTP services, the application could
// directly communicate with the proxy (e.g., by setting HTTP_PROXY) to
// talk to these services.
type ServiceEntry_Resolution int32

const (
	// Assume that incoming connections have already been resolved (to a
	// specific destination IP address). Such connections are typically
	// routed via the proxy using mechanisms such as IP table REDIRECT/
	// eBPF. After performing any routing related transformations, the
	// proxy will forward the connection to the IP address to which the
	// connection was bound.
	ServiceEntry_NONE ServiceEntry_Resolution = 0
	// Use the static IP addresses specified in endpoints (see below) as the
	// backing instances associated with the service.
	ServiceEntry_STATIC ServiceEntry_Resolution = 1
	// Attempt to resolve the IP address by querying the ambient DNS,
	// during request processing. If no endpoints are specified, the proxy
	// will resolve the DNS address specified in the hosts field, if
	// wildcards are not used. If endpoints are specified, the DNS
	// addresses specified in the endpoints will be resolved to determine
	// the destination IP address.  DNS resolution cannot be used with unix
	// domain socket endpoints.
	ServiceEntry_DNS ServiceEntry_Resolution = 2
)

var ServiceEntry_Resolution_name = map[int32]string{
	0: "NONE",
	1: "STATIC",
	2: "DNS",
}
var ServiceEntry_Resolution_value = map[string]int32{
	"NONE":   0,
	"STATIC": 1,
	"DNS":    2,
}

func (x ServiceEntry_Resolution) String() string {
	return proto.EnumName(ServiceEntry_Resolution_name, int32(x))
}
func (ServiceEntry_Resolution) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorServiceEntry, []int{0, 1}
}

// `ServiceEntry` enables adding additional entries into Istio's internal
// service registry, so that auto-discovered services in the mesh can
// access/route to these manually specified services. A service entry
// describes the properties of a service (DNS name, VIPs, ports, protocols,
// endpoints). These services could be external to the mesh (e.g., web
// APIs) or mesh-internal services that are not part of the platform's
// service registry (e.g., a set of VMs talking to services in Kubernetes).
//
// The following configuration adds a set of MongoDB instances running on
// unmanaged VMs to Istio's registry, so that these services can be treated
// as any other service in the mesh. The associated DestinationRule is used
// to initiate mTLS connections to the database instances.
//
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: ServiceEntry
// metadata:
//   name: external-svc-mongocluster
// spec:
//   hosts:
//   - mymongodb.somedomain # not used
//   addresses:
//   - 192.192.192.192/24 # VIPs
//   ports:
//   - number: 27018
//     name: mongodb
//     protocol: MONGO
//   location: MESH_INTERNAL
//   resolution: STATIC
//   endpoints:
//   - address: 2.2.2.2
//   - address: 3.3.3.3
// ```
//
// and the associated DestinationRule
//
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: DestinationRule
// metadata:
//   name: mtls-mongocluster
// spec:
//   host: mymongodb.somedomain
//   trafficPolicy:
//     tls:
//       mode: MUTUAL
//       clientCertificate: /etc/certs/myclientcert.pem
//       privateKey: /etc/certs/client_private_key.pem
//       caCertificates: /etc/certs/rootcacerts.pem
// ```
//
// The following example uses a combination of service entry and TLS
// routing in virtual service to demonstrate the use of SNI routing to
// forward unterminated TLS traffic from the application to external
// services via the sidecar. The sidecar inspects the SNI value in the
// ClientHello message to route to the appropriate external service.
//
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: ServiceEntry
// metadata:
//   name: external-svc-https
// spec:
//   hosts:
//   - api.dropboxapi.com
//   - www.googleapis.com
//   - api.facebook.com
//   location: MESH_EXTERNAL
//   ports:
//   - number: 443
//     name: https
//     protocol: HTTPS
//   resolution: DNS
// ```
//
// And the associated VirtualService to route based on the SNI value.
//
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: VirtualService
// metadata:
//   name: tls-routing
// spec:
//   hosts:
//   - api.dropboxapi.com
//   - www.googleapis.com
//   - api.facebook.com
//   tls:
//   - match:
//     - port: 443
//       sniHosts:
//       - api.dropboxapi.com
//     route:
//     - destination:
//         host: api.dropboxapi.com
//   - match:
//     - port: 443
//       sniHosts:
//       - www.googleapis.com
//     route:
//     - destination:
//         host: www.googleapis.com
//   - match:
//     - port: 443
//       sniHosts:
//       - api.facebook.com
//     route:
//     - destination:
//         host: api.facebook.com
//
// ```
//
// The following example demonstrates the use of a dedicated egress gateway
// through which all external service traffic is forwarded.
//
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: ServiceEntry
// metadata:
//   name: external-svc-httpbin
// spec:
//   hosts:
//   - httpbin.com
//   location: MESH_EXTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//   resolution: DNS
// ```
//
// Define a gateway to handle all egress traffic.
//
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: Gateway
// metadata:
//  name: istio-egressgateway
// spec:
//  selector:
//    istio: egressgateway
//  servers:
//  - port:
//      number: 80
//      name: http
//      protocol: HTTP
//    hosts:
//    - "*"
// ```
//
// And the associated VirtualService to route from the sidecar to the
// gateway service (istio-egressgateway.istio-system.svc.cluster.local), as
// well as route from the gateway to the external service.
//
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: VirtualService
// metadata:
//   name: gateway-routing
// spec:
//   hosts:
//   - httpbin.com
//   gateways:
//   - mesh
//   - istio-egressgateway
//   http:
//   - match:
//     - port: 80
//       gateways:
//       - mesh
//     route:
//     - destination:
//         host: istio-egressgateway.istio-system.svc.cluster.local
//   - match:
//     - port: 80
//       gateway:
//       - istio-egressgateway
//     route:
//     - destination:
//         host: httpbin.com
// ```
//
// The following example demonstrates the use of wildcards in the hosts for
// external services. If the connection has to be routed to the IP address
// requested by the application (i.e. application resolves DNS and attempts
// to connect to a specific IP), the discovery mode must be set to `NONE`.
//
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: ServiceEntry
// metadata:
//   name: external-svc-wildcard-example
// spec:
//   hosts:
//   - "*.bar.com"
//   location: MESH_EXTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//   resolution: NONE
// ```
//
// The following example demonstrates a service that is available via a
// Unix Domain Socket on the host of the client. The resolution must be
// set to STATIC to use unix address endpoints.
//
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: ServiceEntry
// metadata:
//   name: unix-domain-socket-example
// spec:
//   hosts:
//   - "example.unix.local"
//   location: MESH_EXTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//   resolution: STATIC
//   endpoints:
//   - address: unix:///var/run/example/socket
// ```
//
// For HTTP based services, it is possible to create a VirtualService
// backed by multiple DNS addressable endpoints. In such a scenario, the
// application can use the HTTP_PROXY environment variable to transparently
// reroute API calls for the VirtualService to a chosen backend. For
// example, the following configuration creates a non-existent external
// service called foo.bar.com backed by three domains: us.foo.bar.com:8080,
// uk.foo.bar.com:9080, and in.foo.bar.com:7080
//
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: ServiceEntry
// metadata:
//   name: external-svc-dns
// spec:
//   hosts:
//   - foo.bar.com
//   location: MESH_EXTERNAL
//   ports:
//   - number: 80
//     name: https
//     protocol: HTTP
//   resolution: DNS
//   endpoints:
//   - address: us.foo.bar.com
//     ports:
//       https: 8080
//   - address: uk.foo.bar.com
//     ports:
//       https: 9080
//   - address: in.foo.bar.com
//     ports:
//       https: 7080
// ```
//
// With HTTP_PROXY=http://localhost/, calls from the application to
// http://foo.bar.com will be load balanced across the three domains
// specified above. In other words, a call to http://foo.bar.com/baz would
// be translated to http://uk.foo.bar.com/baz.
//
type ServiceEntry struct {
	// REQUIRED. The hosts associated with the ServiceEntry. Could be a DNS
	// name with wildcard prefix (external services only). DNS names in hosts
	// will be ignored if the application accesses the service over non-HTTP
	// protocols such as mongo/opaque TCP/even HTTPS. In such scenarios, the
	// IP addresses specified in the Addresses field or the port will be used
	// to uniquely identify the destination.
	Hosts []string `protobuf:"bytes,1,rep,name=hosts" json:"hosts,omitempty"`
	// The virtual IP addresses associated with the service. Could be CIDR
	// prefix.  For HTTP services, the addresses field will be ignored and
	// the destination will be identified based on the HTTP Host/Authority
	// header. For non-HTTP protocols such as mongo/opaque TCP/even HTTPS,
	// the hosts will be ignored. If one or more IP addresses are specified,
	// the incoming traffic will be identified as belonging to this service
	// if the destination IP matches the IP/CIDRs specified in the addresses
	// field. If the Addresses field is empty, traffic will be identified
	// solely based on the destination port. In such scenarios, the port on
	// which the service is being accessed must not be shared by any other
	// service in the mesh. In other words, the sidecar will behave as a
	// simple TCP proxy, forwarding incoming traffic on a specified port to
	// the specified destination endpoint IP/host. Unix domain socket
	// addresses are not supported in this field.
	Addresses []string `protobuf:"bytes,2,rep,name=addresses" json:"addresses,omitempty"`
	// REQUIRED. The ports associated with the external service. If the
	// Endpoints are unix domain socket addresses, there must be exactly one
	// port.
	Ports []*Port `protobuf:"bytes,3,rep,name=ports" json:"ports,omitempty"`
	// Specify whether the service should be considered external to the mesh
	// or part of the mesh.
	Location ServiceEntry_Location `protobuf:"varint,4,opt,name=location,proto3,enum=istio.networking.v1alpha3.ServiceEntry_Location" json:"location,omitempty"`
	// REQUIRED: Service discovery mode for the hosts. Care must be taken
	// when setting the resolution mode to NONE for a TCP port without
	// accompanying IP addresses. In such cases, traffic to any IP on
	// said port will be allowed (i.e. 0.0.0.0:<port>).
	Resolution ServiceEntry_Resolution `protobuf:"varint,5,opt,name=resolution,proto3,enum=istio.networking.v1alpha3.ServiceEntry_Resolution" json:"resolution,omitempty"`
	// One or more endpoints associated with the service.
	Endpoints []*ServiceEntry_Endpoint `protobuf:"bytes,6,rep,name=endpoints" json:"endpoints,omitempty"`
}

func (m *ServiceEntry) Reset()                    { *m = ServiceEntry{} }
func (m *ServiceEntry) String() string            { return proto.CompactTextString(m) }
func (*ServiceEntry) ProtoMessage()               {}
func (*ServiceEntry) Descriptor() ([]byte, []int) { return fileDescriptorServiceEntry, []int{0} }

func (m *ServiceEntry) GetHosts() []string {
	if m != nil {
		return m.Hosts
	}
	return nil
}

func (m *ServiceEntry) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *ServiceEntry) GetPorts() []*Port {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *ServiceEntry) GetLocation() ServiceEntry_Location {
	if m != nil {
		return m.Location
	}
	return ServiceEntry_MESH_EXTERNAL
}

func (m *ServiceEntry) GetResolution() ServiceEntry_Resolution {
	if m != nil {
		return m.Resolution
	}
	return ServiceEntry_NONE
}

func (m *ServiceEntry) GetEndpoints() []*ServiceEntry_Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

// Endpoint defines a network address (IP or hostname) associated with
// the mesh service.
type ServiceEntry_Endpoint struct {
	// REQUIRED: Address associated with the network endpoint without the
	// port.  Domain names can be used if and only if the resolution is set
	// to DNS, and must be fully-qualified without wildcards. Use the form
	// unix:///absolute/path/to/socket for unix domain socket endpoints.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// Set of ports associated with the endpoint. The ports must be
	// associated with a port name that was declared as part of the
	// service. Do not use for unix:// addresses.
	Ports map[string]uint32 `protobuf:"bytes,2,rep,name=ports" json:"ports,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// One or more labels associated with the endpoint.
	Labels map[string]string `protobuf:"bytes,3,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *ServiceEntry_Endpoint) Reset()         { *m = ServiceEntry_Endpoint{} }
func (m *ServiceEntry_Endpoint) String() string { return proto.CompactTextString(m) }
func (*ServiceEntry_Endpoint) ProtoMessage()    {}
func (*ServiceEntry_Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptorServiceEntry, []int{0, 0}
}

func (m *ServiceEntry_Endpoint) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ServiceEntry_Endpoint) GetPorts() map[string]uint32 {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *ServiceEntry_Endpoint) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func init() {
	proto.RegisterType((*ServiceEntry)(nil), "istio.networking.v1alpha3.ServiceEntry")
	proto.RegisterType((*ServiceEntry_Endpoint)(nil), "istio.networking.v1alpha3.ServiceEntry.Endpoint")
	proto.RegisterEnum("istio.networking.v1alpha3.ServiceEntry_Location", ServiceEntry_Location_name, ServiceEntry_Location_value)
	proto.RegisterEnum("istio.networking.v1alpha3.ServiceEntry_Resolution", ServiceEntry_Resolution_name, ServiceEntry_Resolution_value)
}
func (m *ServiceEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ServiceEntry) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Hosts) > 0 {
		for _, s := range m.Hosts {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Addresses) > 0 {
		for _, s := range m.Addresses {
			dAtA[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Ports) > 0 {
		for _, msg := range m.Ports {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintServiceEntry(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Location != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintServiceEntry(dAtA, i, uint64(m.Location))
	}
	if m.Resolution != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintServiceEntry(dAtA, i, uint64(m.Resolution))
	}
	if len(m.Endpoints) > 0 {
		for _, msg := range m.Endpoints {
			dAtA[i] = 0x32
			i++
			i = encodeVarintServiceEntry(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *ServiceEntry_Endpoint) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ServiceEntry_Endpoint) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintServiceEntry(dAtA, i, uint64(len(m.Address)))
		i += copy(dAtA[i:], m.Address)
	}
	if len(m.Ports) > 0 {
		for k, _ := range m.Ports {
			dAtA[i] = 0x12
			i++
			v := m.Ports[k]
			mapSize := 1 + len(k) + sovServiceEntry(uint64(len(k))) + 1 + sovServiceEntry(uint64(v))
			i = encodeVarintServiceEntry(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintServiceEntry(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x10
			i++
			i = encodeVarintServiceEntry(dAtA, i, uint64(v))
		}
	}
	if len(m.Labels) > 0 {
		for k, _ := range m.Labels {
			dAtA[i] = 0x1a
			i++
			v := m.Labels[k]
			mapSize := 1 + len(k) + sovServiceEntry(uint64(len(k))) + 1 + len(v) + sovServiceEntry(uint64(len(v)))
			i = encodeVarintServiceEntry(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintServiceEntry(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintServiceEntry(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	return i, nil
}

func encodeVarintServiceEntry(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ServiceEntry) Size() (n int) {
	var l int
	_ = l
	if len(m.Hosts) > 0 {
		for _, s := range m.Hosts {
			l = len(s)
			n += 1 + l + sovServiceEntry(uint64(l))
		}
	}
	if len(m.Addresses) > 0 {
		for _, s := range m.Addresses {
			l = len(s)
			n += 1 + l + sovServiceEntry(uint64(l))
		}
	}
	if len(m.Ports) > 0 {
		for _, e := range m.Ports {
			l = e.Size()
			n += 1 + l + sovServiceEntry(uint64(l))
		}
	}
	if m.Location != 0 {
		n += 1 + sovServiceEntry(uint64(m.Location))
	}
	if m.Resolution != 0 {
		n += 1 + sovServiceEntry(uint64(m.Resolution))
	}
	if len(m.Endpoints) > 0 {
		for _, e := range m.Endpoints {
			l = e.Size()
			n += 1 + l + sovServiceEntry(uint64(l))
		}
	}
	return n
}

func (m *ServiceEntry_Endpoint) Size() (n int) {
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovServiceEntry(uint64(l))
	}
	if len(m.Ports) > 0 {
		for k, v := range m.Ports {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovServiceEntry(uint64(len(k))) + 1 + sovServiceEntry(uint64(v))
			n += mapEntrySize + 1 + sovServiceEntry(uint64(mapEntrySize))
		}
	}
	if len(m.Labels) > 0 {
		for k, v := range m.Labels {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovServiceEntry(uint64(len(k))) + 1 + len(v) + sovServiceEntry(uint64(len(v)))
			n += mapEntrySize + 1 + sovServiceEntry(uint64(mapEntrySize))
		}
	}
	return n
}

func sovServiceEntry(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozServiceEntry(x uint64) (n int) {
	return sovServiceEntry(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ServiceEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowServiceEntry
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ServiceEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServiceEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hosts", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthServiceEntry
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hosts = append(m.Hosts, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthServiceEntry
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addresses = append(m.Addresses, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ports", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthServiceEntry
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ports = append(m.Ports, &Port{})
			if err := m.Ports[len(m.Ports)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Location", wireType)
			}
			m.Location = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Location |= (ServiceEntry_Location(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resolution", wireType)
			}
			m.Resolution = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Resolution |= (ServiceEntry_Resolution(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Endpoints", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthServiceEntry
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Endpoints = append(m.Endpoints, &ServiceEntry_Endpoint{})
			if err := m.Endpoints[len(m.Endpoints)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipServiceEntry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthServiceEntry
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ServiceEntry_Endpoint) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowServiceEntry
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Endpoint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Endpoint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthServiceEntry
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ports", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthServiceEntry
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Ports == nil {
				m.Ports = make(map[string]uint32)
			}
			var mapkey string
			var mapvalue uint32
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowServiceEntry
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowServiceEntry
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthServiceEntry
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowServiceEntry
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvalue |= (uint32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipServiceEntry(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthServiceEntry
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Ports[mapkey] = mapvalue
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthServiceEntry
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Labels == nil {
				m.Labels = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowServiceEntry
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowServiceEntry
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthServiceEntry
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowServiceEntry
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthServiceEntry
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipServiceEntry(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthServiceEntry
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Labels[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipServiceEntry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthServiceEntry
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipServiceEntry(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowServiceEntry
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowServiceEntry
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowServiceEntry
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthServiceEntry
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowServiceEntry
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipServiceEntry(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthServiceEntry = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowServiceEntry   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("networking/v1alpha3/service_entry.proto", fileDescriptorServiceEntry) }

var fileDescriptorServiceEntry = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcf, 0x6e, 0xd4, 0x30,
	0x10, 0xc6, 0xeb, 0xa4, 0x9b, 0x6e, 0xa6, 0x14, 0x05, 0x8b, 0x83, 0x59, 0xa1, 0x25, 0xec, 0x85,
	0x48, 0x48, 0xd9, 0xb2, 0x15, 0x52, 0xf9, 0x73, 0x29, 0x10, 0x89, 0x4a, 0x4b, 0x00, 0x6f, 0x0e,
	0x88, 0x4b, 0xe5, 0x76, 0xad, 0xd6, 0x6a, 0x14, 0x47, 0xb6, 0xbb, 0x55, 0x9f, 0x82, 0x77, 0xe2,
	0xc4, 0x91, 0x47, 0x40, 0xfb, 0x24, 0x68, 0x9d, 0x64, 0x93, 0x43, 0xa1, 0xed, 0x2d, 0x33, 0x99,
	0xef, 0x37, 0xe3, 0x6f, 0x6c, 0x78, 0x56, 0x70, 0x73, 0x29, 0xd5, 0xb9, 0x28, 0x4e, 0xc7, 0x8b,
	0x17, 0x2c, 0x2f, 0xcf, 0xd8, 0xde, 0x58, 0x73, 0xb5, 0x10, 0x27, 0xfc, 0x88, 0x17, 0x46, 0x5d,
	0xc5, 0xa5, 0x92, 0x46, 0xe2, 0x47, 0x42, 0x1b, 0x21, 0xe3, 0xb6, 0x3c, 0x6e, 0xca, 0x07, 0x4f,
	0xaf, 0x63, 0x9c, 0x32, 0xc3, 0x2f, 0x59, 0xad, 0x1e, 0xfd, 0xf0, 0xe0, 0xde, 0xac, 0xa2, 0x26,
	0x2b, 0x28, 0x7e, 0x08, 0xbd, 0x33, 0xa9, 0x8d, 0x26, 0x28, 0x74, 0x23, 0x9f, 0x56, 0x01, 0x7e,
	0x0c, 0x3e, 0x9b, 0xcf, 0x15, 0xd7, 0x9a, 0x6b, 0xe2, 0xd8, 0x3f, 0x6d, 0x02, 0xbf, 0x84, 0x5e,
	0x29, 0x95, 0xd1, 0xc4, 0x0d, 0xdd, 0x68, 0x7b, 0xf2, 0x24, 0xfe, 0xe7, 0x48, 0xf1, 0x17, 0xa9,
	0x0c, 0xad, 0xaa, 0xf1, 0x14, 0xfa, 0xb9, 0x3c, 0x61, 0x46, 0xc8, 0x82, 0x6c, 0x86, 0x28, 0xba,
	0x3f, 0xd9, 0xfd, 0x8f, 0xb2, 0x3b, 0x65, 0x3c, 0xad, 0x75, 0x74, 0x4d, 0xc0, 0x14, 0x40, 0x71,
	0x2d, 0xf3, 0x0b, 0xcb, 0xeb, 0x59, 0xde, 0xe4, 0xb6, 0x3c, 0xba, 0x56, 0xd2, 0x0e, 0x05, 0xa7,
	0xe0, 0xf3, 0x62, 0x5e, 0x4a, 0x51, 0x18, 0x4d, 0x3c, 0x7b, 0xb8, 0x5b, 0x8f, 0x98, 0xd4, 0x42,
	0xda, 0x22, 0x06, 0x3f, 0x1d, 0xe8, 0x37, 0x79, 0x4c, 0x60, 0xab, 0xb6, 0x90, 0xa0, 0x10, 0x45,
	0x3e, 0x6d, 0x42, 0xfc, 0xb5, 0xf1, 0xd3, 0xb1, 0x2d, 0xdf, 0xdc, 0xb5, 0xa5, 0x75, 0x59, 0xdb,
	0x5c, 0xe3, 0x75, 0x06, 0x5e, 0xce, 0x8e, 0x79, 0xde, 0xec, 0xe8, 0xed, 0x9d, 0x99, 0x53, 0x2b,
	0xaf, 0xa0, 0x35, 0x6b, 0xb0, 0x0f, 0xd0, 0xb6, 0xc2, 0x01, 0xb8, 0xe7, 0xfc, 0xaa, 0x3e, 0xcc,
	0xea, 0x73, 0x75, 0x99, 0x16, 0x2c, 0xbf, 0xe0, 0xc4, 0x09, 0x51, 0xb4, 0x43, 0xab, 0xe0, 0xb5,
	0xb3, 0x8f, 0x06, 0xaf, 0x60, 0xbb, 0x03, 0xbc, 0x49, 0xea, 0x77, 0xa4, 0xa3, 0x5d, 0xe8, 0x37,
	0xeb, 0xc7, 0x0f, 0x60, 0xe7, 0x53, 0x32, 0xfb, 0x78, 0x94, 0x7c, 0xcb, 0x12, 0x9a, 0x1e, 0x4c,
	0x83, 0x8d, 0x75, 0xea, 0x30, 0xad, 0x53, 0x68, 0xf4, 0x1c, 0xa0, 0x5d, 0x30, 0xee, 0xc3, 0x66,
	0xfa, 0x39, 0x4d, 0x82, 0x0d, 0x0c, 0xe0, 0xcd, 0xb2, 0x83, 0xec, 0xf0, 0x7d, 0x80, 0xf0, 0x16,
	0xb8, 0x1f, 0xd2, 0x59, 0xe0, 0xbc, 0x8b, 0x7f, 0x2d, 0x87, 0xe8, 0xf7, 0x72, 0x88, 0xfe, 0x2c,
	0x87, 0xe8, 0x7b, 0x58, 0xd9, 0x24, 0xe4, 0x98, 0x95, 0x62, 0x7c, 0xcd, 0x7b, 0x3a, 0xf6, 0xec,
	0x43, 0xda, 0xfb, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x9f, 0xa1, 0x1e, 0x0c, 0xb1, 0x03, 0x00, 0x00,
}
