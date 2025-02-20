// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResourceBuilder(t *testing.T) {
	for _, tt := range []string{"default", "all_set", "none_set"} {
		t.Run(tt, func(t *testing.T) {
			cfg := loadResourceAttributesConfig(t, tt)
			rb := NewResourceBuilder(cfg)
			rb.SetHostArch("host.arch-val")
			rb.SetHostCPUCacheL2Size(22)
			rb.SetHostCPUFamily("host.cpu.family-val")
			rb.SetHostCPUModelID("host.cpu.model.id-val")
			rb.SetHostCPUModelName("host.cpu.model.name-val")
			rb.SetHostCPUStepping("host.cpu.stepping-val")
			rb.SetHostCPUVendorID("host.cpu.vendor.id-val")
			rb.SetHostID("host.id-val")
			rb.SetHostIP([]any{"host.ip-item1", "host.ip-item2"})
			rb.SetHostMac([]any{"host.mac-item1", "host.mac-item2"})
			rb.SetHostName("host.name-val")
			rb.SetOsDescription("os.description-val")
			rb.SetOsType("os.type-val")
			rb.SetOsVersion("os.version-val")

			res := rb.Emit()
			assert.Equal(t, 0, rb.Emit().Attributes().Len()) // Second call should return empty Resource

			switch tt {
			case "default":
				assert.Equal(t, 3, res.Attributes().Len())
			case "all_set":
				assert.Equal(t, 14, res.Attributes().Len())
			case "none_set":
				assert.Equal(t, 0, res.Attributes().Len())
				return
			default:
				assert.Failf(t, "unexpected test case: %s", tt)
			}

			val, ok := res.Attributes().Get("host.arch")
			assert.Equal(t, tt == "all_set", ok)
			if ok {
				assert.EqualValues(t, "host.arch-val", val.Str())
			}
			val, ok = res.Attributes().Get("host.cpu.cache.l2.size")
			assert.Equal(t, tt == "all_set", ok)
			if ok {
				assert.EqualValues(t, 22, val.Int())
			}
			val, ok = res.Attributes().Get("host.cpu.family")
			assert.Equal(t, tt == "all_set", ok)
			if ok {
				assert.EqualValues(t, "host.cpu.family-val", val.Str())
			}
			val, ok = res.Attributes().Get("host.cpu.model.id")
			assert.Equal(t, tt == "all_set", ok)
			if ok {
				assert.EqualValues(t, "host.cpu.model.id-val", val.Str())
			}
			val, ok = res.Attributes().Get("host.cpu.model.name")
			assert.Equal(t, tt == "all_set", ok)
			if ok {
				assert.EqualValues(t, "host.cpu.model.name-val", val.Str())
			}
			val, ok = res.Attributes().Get("host.cpu.stepping")
			assert.Equal(t, tt == "all_set", ok)
			if ok {
				assert.EqualValues(t, "host.cpu.stepping-val", val.Str())
			}
			val, ok = res.Attributes().Get("host.cpu.vendor.id")
			assert.Equal(t, tt == "all_set", ok)
			if ok {
				assert.EqualValues(t, "host.cpu.vendor.id-val", val.Str())
			}
			val, ok = res.Attributes().Get("host.id")
			assert.Equal(t, tt == "all_set", ok)
			if ok {
				assert.EqualValues(t, "host.id-val", val.Str())
			}
			val, ok = res.Attributes().Get("host.ip")
			assert.Equal(t, tt == "all_set", ok)
			if ok {
				assert.EqualValues(t, []any{"host.ip-item1", "host.ip-item2"}, val.Slice().AsRaw())
			}
			val, ok = res.Attributes().Get("host.mac")
			assert.Equal(t, tt == "all_set", ok)
			if ok {
				assert.EqualValues(t, []any{"host.mac-item1", "host.mac-item2"}, val.Slice().AsRaw())
			}
			val, ok = res.Attributes().Get("host.name")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "host.name-val", val.Str())
			}
			val, ok = res.Attributes().Get("os.description")
			assert.Equal(t, tt == "all_set", ok)
			if ok {
				assert.EqualValues(t, "os.description-val", val.Str())
			}
			val, ok = res.Attributes().Get("os.type")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "os.type-val", val.Str())
			}
			val, ok = res.Attributes().Get("os.version")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "os.version-val", val.Str())
			}
		})
	}
}
