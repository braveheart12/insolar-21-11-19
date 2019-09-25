//
// Copyright 2019 Insolar Technologies GmbH
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
//

package log

import (
	"bytes"
	"github.com/insolar/insolar/configuration"
	"github.com/insolar/insolar/insolar"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"regexp"
	"testing"
)

func capture(f func()) string {
	defer SaveGlobalLogger()()

	var buf bytes.Buffer

	gl, err := GlobalLogger().Copy().WithOutput(&buf).Build()
	if err != nil {
		panic(err)
	}
	SetGlobalLogger(gl)

	f()

	return buf.String()
}

func assertHelloWorld(t *testing.T, out string) {
	assert.Contains(t, out, "HelloWorld")
}

func TestLog_GlobalLogger_redirection(t *testing.T) {
	defer SaveGlobalLogger()()

	originalG := GlobalLogger()

	var buf bytes.Buffer
	newGL, err := GlobalLogger().Copy().WithOutput(&buf).Build()
	if err != nil {
		panic(err)
	}
	SetGlobalLogger(newGL)
	newCopyLL, _ := GlobalLogger().Copy().BuildLowLatency()

	originalG.Info("viaOriginalGlobal")
	newGL.Info("viaNewInstance")
	GlobalLogger().Info("viaNewGlobal")
	newCopyLL.Info("viaNewLLCopyOfGlobal")

	s := buf.String()
	require.Contains(t, s, "viaOriginalGlobal")
	require.Contains(t, s, "viaNewInstance")
	require.Contains(t, s, "viaNewGlobal")
	require.Contains(t, s, "viaNewLLCopyOfGlobal")
}

func TestLog_GlobalLogger(t *testing.T) {

	assert.NoError(t, SetLevel("debug"))

	assertHelloWorld(t, capture(func() { Debug("HelloWorld") }))
	assertHelloWorld(t, capture(func() { Debugf("%s", "HelloWorld") }))

	assertHelloWorld(t, capture(func() { Info("HelloWorld") }))
	assertHelloWorld(t, capture(func() { Infof("%s", "HelloWorld") }))

	assertHelloWorld(t, capture(func() { Warn("HelloWorld") }))
	assertHelloWorld(t, capture(func() { Warnf("%s", "HelloWorld") }))

	assertHelloWorld(t, capture(func() { Error("HelloWorld") }))
	assertHelloWorld(t, capture(func() { Errorf("%s", "HelloWorld") }))

	assert.Panics(t, func() { capture(func() { Panic("HelloWorld") }) })
	assert.Panics(t, func() { capture(func() { Panicf("%s", "HelloWorld") }) })

	// cyclic run of this test changes loglevel, so revert it back
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	// can't catch os.exit() to test Fatal
	// Fatal("HelloWorld")
	// Fatalln("HelloWorld")
	// Fatalf("%s", "HelloWorld")
}

func TestLog_NewLog_Config(t *testing.T) {
	invalidtests := map[string]configuration.Log{
		"InvalidAdapter":   configuration.Log{Level: "Debug", Adapter: "invalid", Formatter: "text"},
		"InvalidLevel":     configuration.Log{Level: "Invalid", Adapter: "zerolog", Formatter: "text"},
		"InvalidFormatter": configuration.Log{Level: "Debug", Adapter: "zerolog", Formatter: "invalid"},
	}

	for name, test := range invalidtests {
		t.Run(name, func(t *testing.T) {
			logger, err := NewLog(test)
			assert.Nil(t, logger)
			assert.Error(t, err)
		})
	}

	validtests := map[string]configuration.Log{
		"WithAdapter": configuration.Log{Level: "Debug", Adapter: "zerolog", Formatter: "text"},
	}
	for name, test := range validtests {
		t.Run(name, func(t *testing.T) {
			logger, err := NewLog(test)
			assert.NotNil(t, logger)
			assert.NoError(t, err)
		})
	}
}

func TestLog_GlobalLogger_Level(t *testing.T) {
	defer SetLevel("info")
	assert.NoError(t, SetLevel("error"))
	assert.Error(t, SetLevel("errorrr"))
}

func TestLog_AddFields(t *testing.T) {
	errtxt1 := "~CHECK_ERROR_OUTPUT_WITH_FIELDS~"
	errtxt2 := "~CHECK_ERROR_OUTPUT_WITHOUT_FIELDS~"

	var (
		fieldname  = "TraceID"
		fieldvalue = "Trace100500"
	)
	tt := []struct {
		name    string
		fieldfn func(la insolar.Logger) insolar.Logger
	}{
		{
			name: "WithFields",
			fieldfn: func(la insolar.Logger) insolar.Logger {
				fields := map[string]interface{}{fieldname: fieldvalue}
				return la.WithFields(fields)
			},
		},
		{
			name: "WithField",
			fieldfn: func(la insolar.Logger) insolar.Logger {
				return la.WithField(fieldname, fieldvalue)
			},
		},
	}

	for _, tItem := range tt {
		t.Run(tItem.name, func(t *testing.T) {
			la, err := NewLog(configuration.NewLog())
			assert.NoError(t, err)

			var b bytes.Buffer
			logger, err := la.Copy().WithOutput(&b).Build()
			assert.NoError(t, err)

			tItem.fieldfn(logger).Error(errtxt1)
			logger.Error(errtxt2)

			var recitems []string
			for {
				line, err := b.ReadBytes('\n')
				if err != nil && err != io.EOF {
					require.NoError(t, err)
				}

				recitems = append(recitems, string(line))
				if err == io.EOF {
					break
				}
			}
			assert.Contains(t, recitems[0], errtxt1)
			assert.Contains(t, recitems[1], errtxt2)
			assert.Contains(t, recitems[0], fieldvalue)
			assert.NotContains(t, recitems[1], fieldvalue)
		})
	}
}

func TestLog_Timestamp(t *testing.T) {
	for _, adapter := range []string{"zerolog"} {
		adapter := adapter
		t.Run(adapter, func(t *testing.T) {
			logger, err := NewLog(configuration.Log{Level: "info", Adapter: adapter, Formatter: "json"})
			require.NoError(t, err)
			require.NotNil(t, logger)

			var buf bytes.Buffer
			logger, err = logger.Copy().WithOutput(&buf).Build()
			require.NoError(t, err)

			logger.Error("test")

			assert.Regexp(t, regexp.MustCompile("[0-9][0-9]:[0-9][0-9]:[0-9][0-9]"), buf.String())
		})
	}
}

func TestLog_WriteDuration(t *testing.T) {
	for _, adapter := range []string{"zerolog"} {
		adapter := adapter
		t.Run(adapter, func(t *testing.T) {
			logger, err := NewLog(configuration.Log{Level: "info", Adapter: adapter, Formatter: "json"})
			require.NoError(t, err)
			require.NotNil(t, logger)

			var buf bytes.Buffer
			logger, err = logger.Copy().WithOutput(&buf).WithMetrics(insolar.LogMetricsWriteDelayField).Build()
			require.NoError(t, err)

			logger.Error("test")

			s := buf.String()
			assert.Contains(t, s, `,"writeDuration":"`)
		})
	}
}
