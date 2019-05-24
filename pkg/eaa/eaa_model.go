// Copyright 2019 Intel Corporation and Smart-Edge.com, Inc. All rights reserved
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package eaa

import "encoding/json"

type NotificationDescriptor struct {
	// Name of notification
	Name string `json:"name,omitempty"`
	// Version of notification
	Version string `json:"version,omitempty"`
	// Human readable description of notification
	Description string `json:"description,omitempty"`
}

type NotificationFromProducer struct {
	// Name of notification
	Name string `json:"name,omitempty"`
	// Version of notification
	Version string `json:"version,omitempty"`
	// The payload can be any JSON object with a name
	// and version-specific schema.
	Payload json.RawMessage `json:"payload,omitempty"`
}

type NotificationToConsumer struct {
	// Name of notification
	Name string `json:"name,omitempty"`
	// Version of notification
	Version string `json:"version,omitempty"`
	// The payload can be any JSON object with a name
	// and version-specific schema.
	Payload json.RawMessage `json:"payload,omitempty"`
	// URN of the producer
	URN URN `json:"producer,omitempty"`
}

type ServiceList struct {
	Services []Service `json:"services,omitempty"`
}

type Service struct {
	URN           *URN                     `json:"urn,omitempty"`
	Description   string                   `json:"description,omitempty"`
	EndpointURI   string                   `json:"endpoint_uri,omitempty"`
	Status        string                   `json:"status,omitempty"`
	Notifications []NotificationDescriptor `json:"notifications,omitempty"`
}

type SubscriptionList struct {
	Subscriptions []Subscription `json:"subscriptions,omitempty"`
}

type Subscription struct {

	// The name of the producer app. The unique ID is optional for
	// subscribing and if not given will subscribe to any producer in the
	// namespace.
	URN *URN `json:"urn,omitempty"`

	// The list of all notification types registered by all producers in
	// this namespace.
	Notifications []NotificationDescriptor `json:"notifications,omitempty"`
}

type URN struct {

	// The per-namespace unique portion of the URN that when appended to
	// the namespace with a separator forms the complete URN.
	ID string `json:"id,omitempty"`

	// The non-unique portion of the URN that identifies the class excluding
	// a trailing separator.
	Namespace string `json:"namespace,omitempty"`
}
