//
//    Copyright 2019 Insolar Technologies
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
//

package critlog

import (
	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/log/logoutput"
	"io"
)

func NewFatalDirectWriter(output *logoutput.Adapter) *FatalDirectWriter {
	if output == nil {
		panic("illegal value")
	}

	return &FatalDirectWriter{
		output: output,
	}
}

var _ insolar.LogLevelWriter = &FatalDirectWriter{}
var _ io.WriteCloser = &FatalDirectWriter{}

type FatalDirectWriter struct {
	output *logoutput.Adapter
}

func (p *FatalDirectWriter) Close() error {
	return p.output.Close()
}

func (p *FatalDirectWriter) Flush() error {
	return p.output.Flush()
}

func (p *FatalDirectWriter) Write(b []byte) (n int, err error) {
	return p.output.Write(b)
}

func (p *FatalDirectWriter) LowLatencyWrite(level insolar.LogLevel, b []byte) (int, error) {
	return p.LogLevelWrite(level, b)
}

func (p *FatalDirectWriter) IsLowLatencySupported() bool {
	return false
}

func (p *FatalDirectWriter) LogLevelWrite(level insolar.LogLevel, b []byte) (n int, err error) {
	switch level {
	case insolar.FatalLevel:
		if !p.output.SetFatal() {
			break
		}
		n, _ = p.output.DirectLevelWrite(level, b)
		_ = p.output.DirectFlushFatal()
		return n, nil

	case insolar.PanicLevel:
		n, err = p.output.LogLevelWrite(level, b)
		_ = p.output.Flush()
		return n, err
	}
	return p.output.LogLevelWrite(level, b)
}
