package codersdk

import (
	"encoding/json"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestAPIAllowListTarget_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	all := AllowAllTarget()
	b, err := json.Marshal(all)
	require.NoError(t, err)
	require.JSONEq(t, `"*:*"`, string(b))
	var rt APIAllowListTarget
	require.NoError(t, json.Unmarshal(b, &rt))
	require.True(t, rt.Type.IsAny())
	require.True(t, rt.ID.IsAny())

	ty := AllowTypeTarget(ResourceWorkspace)
	b, err = json.Marshal(ty)
	require.NoError(t, err)
	require.JSONEq(t, `"workspace:*"`, string(b))
	require.NoError(t, json.Unmarshal(b, &rt))
	r, ok := rt.Type.Value()
	require.True(t, ok)
	require.Equal(t, ResourceWorkspace, r)
	require.True(t, rt.ID.IsAny())

	id := uuid.New()
	res := AllowResourceTarget(ResourceTemplate, id)
	b, err = json.Marshal(res)
	require.NoError(t, err)
	exp := `"template:` + id.String() + `"`
	require.JSONEq(t, exp, string(b))
}
