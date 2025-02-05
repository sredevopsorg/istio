// GENERATED FILE -- DO NOT EDIT
//

package kind

import (
	"istio.io/istio/pkg/config"
	"istio.io/istio/pkg/config/schema/gvk"
)

const (
	Address Kind = iota
	AuthorizationPolicy
	ConfigMap
	CustomResourceDefinition
	Deployment
	DestinationRule
	Endpoints
	EnvoyFilter
	GRPCRoute
	Gateway
	GatewayClass
	HTTPRoute
	Ingress
	KubernetesGateway
	MeshConfig
	MeshNetworks
	MutatingWebhookConfiguration
	Namespace
	Node
	PeerAuthentication
	Pod
	ProxyConfig
	ReferenceGrant
	RequestAuthentication
	Secret
	Service
	ServiceAccount
	ServiceEntry
	Sidecar
	TCPRoute
	TLSRoute
	Telemetry
	UDPRoute
	VirtualService
	WasmPlugin
	WorkloadEntry
	WorkloadGroup
)

func (k Kind) String() string {
	switch k {
	case Address:
		return "Address"
	case AuthorizationPolicy:
		return "AuthorizationPolicy"
	case ConfigMap:
		return "ConfigMap"
	case CustomResourceDefinition:
		return "CustomResourceDefinition"
	case Deployment:
		return "Deployment"
	case DestinationRule:
		return "DestinationRule"
	case Endpoints:
		return "Endpoints"
	case EnvoyFilter:
		return "EnvoyFilter"
	case GRPCRoute:
		return "GRPCRoute"
	case Gateway:
		return "Gateway"
	case GatewayClass:
		return "GatewayClass"
	case HTTPRoute:
		return "HTTPRoute"
	case Ingress:
		return "Ingress"
	case KubernetesGateway:
		return "Gateway"
	case MeshConfig:
		return "MeshConfig"
	case MeshNetworks:
		return "MeshNetworks"
	case MutatingWebhookConfiguration:
		return "MutatingWebhookConfiguration"
	case Namespace:
		return "Namespace"
	case Node:
		return "Node"
	case PeerAuthentication:
		return "PeerAuthentication"
	case Pod:
		return "Pod"
	case ProxyConfig:
		return "ProxyConfig"
	case ReferenceGrant:
		return "ReferenceGrant"
	case RequestAuthentication:
		return "RequestAuthentication"
	case Secret:
		return "Secret"
	case Service:
		return "Service"
	case ServiceAccount:
		return "ServiceAccount"
	case ServiceEntry:
		return "ServiceEntry"
	case Sidecar:
		return "Sidecar"
	case TCPRoute:
		return "TCPRoute"
	case TLSRoute:
		return "TLSRoute"
	case Telemetry:
		return "Telemetry"
	case UDPRoute:
		return "UDPRoute"
	case VirtualService:
		return "VirtualService"
	case WasmPlugin:
		return "WasmPlugin"
	case WorkloadEntry:
		return "WorkloadEntry"
	case WorkloadGroup:
		return "WorkloadGroup"
	default:
		return "Unknown"
	}
}

func MustFromGVK(g config.GroupVersionKind) Kind {
	switch g {
	case gvk.AuthorizationPolicy:
		return AuthorizationPolicy
	case gvk.ConfigMap:
		return ConfigMap
	case gvk.CustomResourceDefinition:
		return CustomResourceDefinition
	case gvk.Deployment:
		return Deployment
	case gvk.DestinationRule:
		return DestinationRule
	case gvk.Endpoints:
		return Endpoints
	case gvk.EnvoyFilter:
		return EnvoyFilter
	case gvk.GRPCRoute:
		return GRPCRoute
	case gvk.Gateway:
		return Gateway
	case gvk.GatewayClass:
		return GatewayClass
	case gvk.HTTPRoute:
		return HTTPRoute
	case gvk.Ingress:
		return Ingress
	case gvk.KubernetesGateway:
		return KubernetesGateway
	case gvk.MeshConfig:
		return MeshConfig
	case gvk.MeshNetworks:
		return MeshNetworks
	case gvk.MutatingWebhookConfiguration:
		return MutatingWebhookConfiguration
	case gvk.Namespace:
		return Namespace
	case gvk.Node:
		return Node
	case gvk.PeerAuthentication:
		return PeerAuthentication
	case gvk.Pod:
		return Pod
	case gvk.ProxyConfig:
		return ProxyConfig
	case gvk.ReferenceGrant:
		return ReferenceGrant
	case gvk.RequestAuthentication:
		return RequestAuthentication
	case gvk.Secret:
		return Secret
	case gvk.Service:
		return Service
	case gvk.ServiceAccount:
		return ServiceAccount
	case gvk.ServiceEntry:
		return ServiceEntry
	case gvk.Sidecar:
		return Sidecar
	case gvk.TCPRoute:
		return TCPRoute
	case gvk.TLSRoute:
		return TLSRoute
	case gvk.Telemetry:
		return Telemetry
	case gvk.UDPRoute:
		return UDPRoute
	case gvk.VirtualService:
		return VirtualService
	case gvk.WasmPlugin:
		return WasmPlugin
	case gvk.WorkloadEntry:
		return WorkloadEntry
	case gvk.WorkloadGroup:
		return WorkloadGroup
	}

	panic("unknown kind: " + g.String())
}
