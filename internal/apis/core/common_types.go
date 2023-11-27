// Copyright 2022 IronCore authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core

const (
	ReconcileRequestAnnotation = "reconcile.apinet.ironcore.dev/requestedAt"

	// APINetletsGroup is the system rbac group all apinetlets are in.
	APINetletsGroup = "apinet.ironcore.dev:system:apinetlets"

	// APINetletUserNamePrefix is the prefix all apinetlet users should have.
	APINetletUserNamePrefix = "apinet.ironcore.dev:system:apinetlet:"

	// MetalnetletsGroup is the system rbac group all metalnetlets are in.
	MetalnetletsGroup = "apinet.ironcore.dev:system:metalnetlets"

	// MetalnetletUserNamePrefix is the prefix all metalnetlet users should have.
	MetalnetletUserNamePrefix = "apinet.ironcore.dev:system:metalnetlet:"
)

// APINetletCommonName constructs the common name for a certificate of an apinetlet user.
func APINetletCommonName(name string) string {
	return APINetletUserNamePrefix + name
}

// MetalnetletCommonName constructs the common name for a certificate of a metalnetlet user.
func MetalnetletCommonName(name string) string {
	return MetalnetletUserNamePrefix + name
}
