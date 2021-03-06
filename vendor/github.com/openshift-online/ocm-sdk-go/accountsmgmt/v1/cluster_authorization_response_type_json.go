/*
Copyright (c) 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1

import (
	"io"

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// MarshalClusterAuthorizationResponse writes a value of the 'cluster_authorization_response' type to the given writer.
func MarshalClusterAuthorizationResponse(object *ClusterAuthorizationResponse, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	writeClusterAuthorizationResponse(object, stream)
	stream.Flush()
	return stream.Error
}

// writeClusterAuthorizationResponse writes a value of the 'cluster_authorization_response' type to the given stream.
func writeClusterAuthorizationResponse(object *ClusterAuthorizationResponse, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	if object.allowed != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("allowed")
		stream.WriteBool(*object.allowed)
		count++
	}
	if object.excessResources != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("excess_resources")
		writeReservedResourceList(object.excessResources, stream)
		count++
	}
	if object.subscription != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("subscription")
		writeSubscription(object.subscription, stream)
		count++
	}
	stream.WriteObjectEnd()
}

// UnmarshalClusterAuthorizationResponse reads a value of the 'cluster_authorization_response' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalClusterAuthorizationResponse(source interface{}) (object *ClusterAuthorizationResponse, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = readClusterAuthorizationResponse(iterator)
	err = iterator.Error
	return
}

// readClusterAuthorizationResponse reads a value of the 'cluster_authorization_response' type from the given iterator.
func readClusterAuthorizationResponse(iterator *jsoniter.Iterator) *ClusterAuthorizationResponse {
	object := &ClusterAuthorizationResponse{}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "allowed":
			value := iterator.ReadBool()
			object.allowed = &value
		case "excess_resources":
			value := readReservedResourceList(iterator)
			object.excessResources = value
		case "subscription":
			value := readSubscription(iterator)
			object.subscription = value
		default:
			iterator.ReadAny()
		}
	}
	return object
}
