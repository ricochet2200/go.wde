/*
   Copyright 2012 John Asmuth

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

package main

import (
	"github.com/skelterjohn/go.wde/x11"
	"github.com/skelterjohn/go.wde"
	"github.com/skelterjohn/go.wde/wdetest"
)

func wgen() (w wde.Window, err error) {
	w, err = x11.NewWindow()
	return
}

func main() {
	wdetest.Run(wgen)
}
