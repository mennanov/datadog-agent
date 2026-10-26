// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2026-present Datadog, Inc.

package status

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	corestatus "github.com/DataDog/datadog-agent/comp/core/status"
)

func TestADPForwarderStatus(t *testing.T) {
	stats := adpForwarderStatsFromJSON(`{"agent-data-plane":{"forwarder":{"Transactions":{"Success":2}}},"process-agent":{"process":{"CPU":1}}}`)
	assert.Contains(t, stats, "agent-data-plane")
	assert.NotContains(t, stats, "process-agent")

	provider := GetADPForwarderProvider()
	assert.Equal(t, "ADP Forwarder", provider.Name())
	assert.Equal(t, "ADP Forwarder", provider.Section())

	// Test the renderers with the parsed remote data; rendering the live expvar is covered by the bridge tests.
	info := map[string]interface{}{"adpForwarderStats": stats}
	var text, html bytes.Buffer
	require.NoError(t, corestatus.RenderText(templatesFS, "adp_forwarder.tmpl", &text, info))
	require.NoError(t, corestatus.RenderHTML(templatesFS, "adp_forwarder_html.tmpl", &html, info))
	assert.Contains(t, text.String(), "agent-data-plane")
	assert.Contains(t, text.String(), "Success: 2")
	assert.Contains(t, html.String(), "agent-data-plane")
	assert.Contains(t, html.String(), "Success: 2")
}
