// Copyright 2017 Drew J. Sonne. All rights reserved.
//
// Use of this source code is governed by a LGPLv3-style
// license that can be found in the LICENSE file.

/*
Package cli provides a gocd cli tool wrapping the gocd package.

Usage:

	gocd --help

	export GOCD_PASSWORD=mysecret
	gocd -server https://goserver:8154/go -username user list-agents

*/
package cli
