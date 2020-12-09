// Licensed to Apache Software Foundation (ASF) under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Apache Software Foundation (ASF) licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package plugin

import "reflect"

// Plugin defines the plugin model in Satellite.
type Plugin interface {
	// Name returns the name of the specific plugin.
	Name() string
	// Description returns the description of the specific plugin.
	Description() string
}

// NameFinderFunc is used to get the plugin name from different plugin configs.
type NameFinderFunc func(config interface{}) string

// InitializingFunc is used to initialize the specific plugin.
type InitializingFunc func(plugin Plugin, config interface{})

// CallbackFunc would be invoked after initializing.
type CallbackFunc func(plugin Plugin)

// RegInfo is a metadata to be registered in the global type registry.
type RegInfo struct {
	// the plugin interface type. (required)
	PluginType reflect.Type
	// the plugin name finder,and the default value is defaultNameFinder (optional,default value is defaultNameFinder)
	NameFinder NameFinderFunc
	// the plugin initializer (optional, default value is defaultInitializing)
	Initializing InitializingFunc
	// the plugin initializer callback func (optional, default value is defaultCallBack)
	Callback CallbackFunc
}
