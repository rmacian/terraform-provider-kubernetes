package kubernetes

import "k8s.io/api/core/v1"

func flattenHostaliases(in []v1.HostAlias) []interface{} {
	att := make([]interface{}, len(in))
	for i, v := range in {
		c := make(map[string]interface{})
		c["ip"] = v.IP
		if len(v.Hostnames) > 0 {
			c["hostnames"] = v.Hostnames
		}
		att[i] = c
	}
	return []interface{}{att}
}
