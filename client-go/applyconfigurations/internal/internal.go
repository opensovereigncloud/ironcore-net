/*
 * Copyright (c) 2022 by the OnMetal authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package internal

import (
	"fmt"
	"sync"

	typed "sigs.k8s.io/structured-merge-diff/v4/typed"
)

func Parser() *typed.Parser {
	parserOnce.Do(func() {
		var err error
		parser, err = typed.NewParser(schemaYAML)
		if err != nil {
			panic(fmt.Sprintf("Failed to parse schema: %v", err))
		}
	})
	return parser
}

var parserOnce sync.Once
var parser *typed.Parser
var schemaYAML = typed.YAMLObject(`types:
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.Affinity
  map:
    fields:
    - name: instanceAntiAffinity
      type:
        namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.InstanceAntiAffinity
    - name: nodeAffinity
      type:
        namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NodeAffinity
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.DaemonSet
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.DaemonSetSpec
      default: {}
    - name: status
      type:
        namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.DaemonSetStatus
      default: {}
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.DaemonSetSpec
  map:
    fields:
    - name: nodeSelector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
    - name: template
      type:
        namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.InstanceTemplate
      default: {}
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.DaemonSetStatus
  map:
    fields:
    - name: collisionCount
      type:
        scalar: numeric
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.IP
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.IPSpec
      default: {}
    - name: status
      type:
        namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.IPStatus
      default: {}
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.IPAddress
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.IPAddressSpec
      default: {}
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.IPAddressClaimRef
  map:
    fields:
    - name: group
      type:
        scalar: string
    - name: name
      type:
        scalar: string
    - name: namespace
      type:
        scalar: string
    - name: resource
      type:
        scalar: string
    - name: uid
      type:
        scalar: string
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.IPAddressSpec
  map:
    fields:
    - name: claimRef
      type:
        namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.IPAddressClaimRef
      default: {}
    - name: ip
      type:
        namedType: com.github.onmetal.onmetal-api-net.apimachinery.api.net.IP
      default: {}
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.IPClaimRef
  map:
    fields:
    - name: group
      type:
        scalar: string
    - name: name
      type:
        scalar: string
    - name: resource
      type:
        scalar: string
    - name: uid
      type:
        scalar: string
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.IPSpec
  map:
    fields:
    - name: claimRef
      type:
        namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.IPClaimRef
    - name: ip
      type:
        namedType: com.github.onmetal.onmetal-api-net.apimachinery.api.net.IP
      default: {}
    - name: ipFamily
      type:
        scalar: string
    - name: type
      type:
        scalar: string
      default: ""
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.IPStatus
  map:
    elementType:
      scalar: untyped
      list:
        elementType:
          namedType: __untyped_atomic_
        elementRelationship: atomic
      map:
        elementType:
          namedType: __untyped_deduced_
        elementRelationship: separable
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.Instance
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.InstanceSpec
      default: {}
    - name: status
      type:
        namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.InstanceStatus
      default: {}
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.InstanceAffinityTerm
  map:
    fields:
    - name: labelSelector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
    - name: topologyKey
      type:
        scalar: string
      default: ""
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.InstanceAntiAffinity
  map:
    fields:
    - name: requiredDuringSchedulingIgnoredDuringExecution
      type:
        list:
          elementType:
            namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.InstanceAffinityTerm
          elementRelationship: atomic
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.InstanceSpec
  map:
    fields:
    - name: affinity
      type:
        namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.Affinity
    - name: ips
      type:
        list:
          elementType:
            namedType: com.github.onmetal.onmetal-api-net.apimachinery.api.net.IP
          elementRelationship: atomic
    - name: loadBalancerPorts
      type:
        list:
          elementType:
            namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.LoadBalancerPort
          elementRelationship: atomic
    - name: loadBalancerType
      type:
        scalar: string
      default: ""
    - name: networkRef
      type:
        namedType: io.k8s.api.core.v1.LocalObjectReference
      default: {}
    - name: nodeRef
      type:
        namedType: io.k8s.api.core.v1.LocalObjectReference
    - name: topologySpreadConstraints
      type:
        list:
          elementType:
            namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.TopologySpreadConstraint
          elementRelationship: atomic
    - name: type
      type:
        scalar: string
      default: ""
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.InstanceStatus
  map:
    fields:
    - name: collisionCount
      type:
        scalar: numeric
    - name: ips
      type:
        list:
          elementType:
            namedType: com.github.onmetal.onmetal-api-net.apimachinery.api.net.IP
          elementRelationship: atomic
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.InstanceTemplate
  map:
    fields:
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.InstanceSpec
      default: {}
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.LoadBalancer
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.LoadBalancerSpec
      default: {}
    - name: status
      type:
        namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.LoadBalancerStatus
      default: {}
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.LoadBalancerDestination
  map:
    fields:
    - name: ip
      type:
        namedType: com.github.onmetal.onmetal-api-net.apimachinery.api.net.IP
      default: {}
    - name: targetRef
      type:
        namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.LoadBalancerTargetRef
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.LoadBalancerIP
  map:
    fields:
    - name: ip
      type:
        namedType: com.github.onmetal.onmetal-api-net.apimachinery.api.net.IP
      default: {}
    - name: ipFamily
      type:
        scalar: string
    - name: name
      type:
        scalar: string
      default: ""
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.LoadBalancerPort
  map:
    fields:
    - name: endPort
      type:
        scalar: numeric
    - name: port
      type:
        scalar: numeric
      default: 0
    - name: protocol
      type:
        scalar: string
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.LoadBalancerRouting
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: destinations
      type:
        list:
          elementType:
            namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.LoadBalancerDestination
          elementRelationship: atomic
    - name: ips
      type:
        list:
          elementType:
            namedType: com.github.onmetal.onmetal-api-net.apimachinery.api.net.IP
          elementRelationship: atomic
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.LoadBalancerSpec
  map:
    fields:
    - name: ips
      type:
        list:
          elementType:
            namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.LoadBalancerIP
          elementRelationship: associative
          keys:
          - name
    - name: networkRef
      type:
        namedType: io.k8s.api.core.v1.LocalObjectReference
      default: {}
    - name: nodeSelector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
    - name: ports
      type:
        list:
          elementType:
            namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.LoadBalancerPort
          elementRelationship: atomic
    - name: template
      type:
        namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.InstanceTemplate
      default: {}
    - name: type
      type:
        scalar: string
      default: ""
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.LoadBalancerStatus
  map:
    fields:
    - name: collisionCount
      type:
        scalar: numeric
    - name: ips
      type:
        list:
          elementType:
            namedType: com.github.onmetal.onmetal-api-net.apimachinery.api.net.IP
          elementRelationship: atomic
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.LoadBalancerTargetRef
  map:
    fields:
    - name: name
      type:
        scalar: string
      default: ""
    - name: nodeRef
      type:
        namedType: io.k8s.api.core.v1.LocalObjectReference
      default: {}
    - name: uid
      type:
        scalar: string
      default: ""
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NATGateway
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NATGatewaySpec
      default: {}
    - name: status
      type:
        namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NATGatewayStatus
      default: {}
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NATGatewayAutoscaler
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NATGatewayAutoscalerSpec
      default: {}
    - name: status
      type:
        namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NATGatewayAutoscalerStatus
      default: {}
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NATGatewayAutoscalerSpec
  map:
    fields:
    - name: maxPublicIPs
      type:
        scalar: numeric
    - name: minPublicIPs
      type:
        scalar: numeric
    - name: natGatewayRef
      type:
        namedType: io.k8s.api.core.v1.LocalObjectReference
      default: {}
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NATGatewayAutoscalerStatus
  map:
    elementType:
      scalar: untyped
      list:
        elementType:
          namedType: __untyped_atomic_
        elementRelationship: atomic
      map:
        elementType:
          namedType: __untyped_deduced_
        elementRelationship: separable
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NATGatewayIP
  map:
    fields:
    - name: ip
      type:
        namedType: com.github.onmetal.onmetal-api-net.apimachinery.api.net.IP
      default: {}
    - name: name
      type:
        scalar: string
      default: ""
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NATGatewaySpec
  map:
    fields:
    - name: ipFamily
      type:
        scalar: string
      default: ""
    - name: ips
      type:
        list:
          elementType:
            namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NATGatewayIP
          elementRelationship: associative
          keys:
          - name
    - name: networkRef
      type:
        namedType: io.k8s.api.core.v1.LocalObjectReference
      default: {}
    - name: portsPerNetworkInterface
      type:
        scalar: numeric
      default: 0
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NATGatewayStatus
  map:
    fields:
    - name: requestedNATIPs
      type:
        scalar: numeric
    - name: usedNATIPs
      type:
        scalar: numeric
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NATIP
  map:
    fields:
    - name: ip
      type:
        namedType: com.github.onmetal.onmetal-api-net.apimachinery.api.net.IP
      default: {}
    - name: sections
      type:
        list:
          elementType:
            namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NATIPSection
          elementRelationship: atomic
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NATIPSection
  map:
    fields:
    - name: endPort
      type:
        scalar: numeric
      default: 0
    - name: ip
      type:
        namedType: com.github.onmetal.onmetal-api-net.apimachinery.api.net.IP
      default: {}
    - name: port
      type:
        scalar: numeric
      default: 0
    - name: targetRef
      type:
        namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NATTableIPTargetRef
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NATTable
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: ips
      type:
        list:
          elementType:
            namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NATIP
          elementRelationship: atomic
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NATTableIPTargetRef
  map:
    fields:
    - name: name
      type:
        scalar: string
      default: ""
    - name: nodeRef
      type:
        namedType: io.k8s.api.core.v1.LocalObjectReference
      default: {}
    - name: uid
      type:
        scalar: string
      default: ""
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.Network
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NetworkSpec
      default: {}
    - name: status
      type:
        namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NetworkStatus
      default: {}
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NetworkID
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NetworkIDSpec
      default: {}
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NetworkIDClaimRef
  map:
    fields:
    - name: group
      type:
        scalar: string
    - name: name
      type:
        scalar: string
    - name: namespace
      type:
        scalar: string
    - name: resource
      type:
        scalar: string
    - name: uid
      type:
        scalar: string
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NetworkIDSpec
  map:
    fields:
    - name: claimRef
      type:
        namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NetworkIDClaimRef
      default: {}
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NetworkInterface
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NetworkInterfaceSpec
      default: {}
    - name: status
      type:
        namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NetworkInterfaceStatus
      default: {}
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NetworkInterfaceNAT
  map:
    fields:
    - name: claimRef
      type:
        namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NetworkInterfaceNATClaimRef
      default: {}
    - name: ipFamily
      type:
        scalar: string
      default: ""
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NetworkInterfaceNATClaimRef
  map:
    fields:
    - name: name
      type:
        scalar: string
      default: ""
    - name: uid
      type:
        scalar: string
      default: ""
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NetworkInterfacePublicIP
  map:
    fields:
    - name: ip
      type:
        namedType: com.github.onmetal.onmetal-api-net.apimachinery.api.net.IP
      default: {}
    - name: ipFamily
      type:
        scalar: string
    - name: name
      type:
        scalar: string
      default: ""
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NetworkInterfaceSpec
  map:
    fields:
    - name: ips
      type:
        list:
          elementType:
            namedType: com.github.onmetal.onmetal-api-net.apimachinery.api.net.IP
          elementRelationship: atomic
    - name: natGateways
      type:
        list:
          elementType:
            namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NetworkInterfaceNAT
          elementRelationship: atomic
    - name: networkRef
      type:
        namedType: io.k8s.api.core.v1.LocalObjectReference
      default: {}
    - name: nodeRef
      type:
        namedType: io.k8s.api.core.v1.LocalObjectReference
      default: {}
    - name: prefixes
      type:
        list:
          elementType:
            namedType: com.github.onmetal.onmetal-api-net.apimachinery.api.net.IPPrefix
          elementRelationship: atomic
    - name: publicIPs
      type:
        list:
          elementType:
            namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NetworkInterfacePublicIP
          elementRelationship: associative
          keys:
          - name
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NetworkInterfaceStatus
  map:
    fields:
    - name: natIPs
      type:
        list:
          elementType:
            namedType: com.github.onmetal.onmetal-api-net.apimachinery.api.net.IP
          elementRelationship: atomic
    - name: pciAddress
      type:
        namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.PCIAddress
    - name: prefixes
      type:
        list:
          elementType:
            namedType: com.github.onmetal.onmetal-api-net.apimachinery.api.net.IPPrefix
          elementRelationship: atomic
    - name: publicIPs
      type:
        list:
          elementType:
            namedType: com.github.onmetal.onmetal-api-net.apimachinery.api.net.IP
          elementRelationship: atomic
    - name: state
      type:
        scalar: string
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NetworkSpec
  map:
    fields:
    - name: id
      type:
        scalar: string
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NetworkStatus
  map:
    elementType:
      scalar: untyped
      list:
        elementType:
          namedType: __untyped_atomic_
        elementRelationship: atomic
      map:
        elementType:
          namedType: __untyped_deduced_
        elementRelationship: separable
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.Node
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NodeSpec
      default: {}
    - name: status
      type:
        namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NodeStatus
      default: {}
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NodeAffinity
  map:
    fields:
    - name: requiredDuringSchedulingIgnoredDuringExecution
      type:
        namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NodeSelector
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NodeSelector
  map:
    fields:
    - name: nodeSelectorTerms
      type:
        list:
          elementType:
            namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NodeSelectorTerm
          elementRelationship: atomic
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NodeSelectorRequirement
  map:
    fields:
    - name: key
      type:
        scalar: string
      default: ""
    - name: operator
      type:
        scalar: string
      default: ""
    - name: values
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NodeSelectorTerm
  map:
    fields:
    - name: matchExpressions
      type:
        list:
          elementType:
            namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NodeSelectorRequirement
          elementRelationship: atomic
    - name: matchFields
      type:
        list:
          elementType:
            namedType: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NodeSelectorRequirement
          elementRelationship: atomic
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NodeSpec
  map:
    elementType:
      scalar: untyped
      list:
        elementType:
          namedType: __untyped_atomic_
        elementRelationship: atomic
      map:
        elementType:
          namedType: __untyped_deduced_
        elementRelationship: separable
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.NodeStatus
  map:
    elementType:
      scalar: untyped
      list:
        elementType:
          namedType: __untyped_atomic_
        elementRelationship: atomic
      map:
        elementType:
          namedType: __untyped_deduced_
        elementRelationship: separable
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.PCIAddress
  map:
    fields:
    - name: bus
      type:
        scalar: string
    - name: domain
      type:
        scalar: string
    - name: function
      type:
        scalar: string
    - name: slot
      type:
        scalar: string
- name: com.github.onmetal.onmetal-api-net.api.core.v1alpha1.TopologySpreadConstraint
  map:
    fields:
    - name: labelSelector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
    - name: maxSkew
      type:
        scalar: numeric
      default: 0
    - name: topologyKey
      type:
        scalar: string
      default: ""
    - name: whenUnsatisfiable
      type:
        scalar: string
      default: ""
- name: com.github.onmetal.onmetal-api-net.apimachinery.api.net.IP
  scalar: untyped
- name: com.github.onmetal.onmetal-api-net.apimachinery.api.net.IPPrefix
  scalar: untyped
- name: io.k8s.api.core.v1.LocalObjectReference
  map:
    fields:
    - name: name
      type:
        scalar: string
    elementRelationship: atomic
- name: io.k8s.apimachinery.pkg.apis.meta.v1.FieldsV1
  map:
    elementType:
      scalar: untyped
      list:
        elementType:
          namedType: __untyped_atomic_
        elementRelationship: atomic
      map:
        elementType:
          namedType: __untyped_deduced_
        elementRelationship: separable
- name: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
  map:
    fields:
    - name: matchExpressions
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelectorRequirement
          elementRelationship: atomic
    - name: matchLabels
      type:
        map:
          elementType:
            scalar: string
    elementRelationship: atomic
- name: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelectorRequirement
  map:
    fields:
    - name: key
      type:
        scalar: string
      default: ""
    - name: operator
      type:
        scalar: string
      default: ""
    - name: values
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
- name: io.k8s.apimachinery.pkg.apis.meta.v1.ManagedFieldsEntry
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: fieldsType
      type:
        scalar: string
    - name: fieldsV1
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.FieldsV1
    - name: manager
      type:
        scalar: string
    - name: operation
      type:
        scalar: string
    - name: subresource
      type:
        scalar: string
    - name: time
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
- name: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
  map:
    fields:
    - name: annotations
      type:
        map:
          elementType:
            scalar: string
    - name: creationTimestamp
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: deletionGracePeriodSeconds
      type:
        scalar: numeric
    - name: deletionTimestamp
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
    - name: finalizers
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
    - name: generateName
      type:
        scalar: string
    - name: generation
      type:
        scalar: numeric
    - name: labels
      type:
        map:
          elementType:
            scalar: string
    - name: managedFields
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ManagedFieldsEntry
          elementRelationship: atomic
    - name: name
      type:
        scalar: string
    - name: namespace
      type:
        scalar: string
    - name: ownerReferences
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.OwnerReference
          elementRelationship: associative
          keys:
          - uid
    - name: resourceVersion
      type:
        scalar: string
    - name: selfLink
      type:
        scalar: string
    - name: uid
      type:
        scalar: string
- name: io.k8s.apimachinery.pkg.apis.meta.v1.OwnerReference
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
      default: ""
    - name: blockOwnerDeletion
      type:
        scalar: boolean
    - name: controller
      type:
        scalar: boolean
    - name: kind
      type:
        scalar: string
      default: ""
    - name: name
      type:
        scalar: string
      default: ""
    - name: uid
      type:
        scalar: string
      default: ""
    elementRelationship: atomic
- name: io.k8s.apimachinery.pkg.apis.meta.v1.Time
  scalar: untyped
- name: __untyped_atomic_
  scalar: untyped
  list:
    elementType:
      namedType: __untyped_atomic_
    elementRelationship: atomic
  map:
    elementType:
      namedType: __untyped_atomic_
    elementRelationship: atomic
- name: __untyped_deduced_
  scalar: untyped
  list:
    elementType:
      namedType: __untyped_atomic_
    elementRelationship: atomic
  map:
    elementType:
      namedType: __untyped_deduced_
    elementRelationship: separable
`)
