# emqx-influxdb-exporter

##### environment variables
- INFLUX_URL
- EMQX_URL
- EMQX_APP_SECRET
- EMQX_APP_SECRET


By default exporter will collect stats and metrics from emqx.
If you want to add further data from outside besides the send to this
endpoint.

```text
POST /v1/unified
```
body
```json
{
  "collection_name": "test_measurement",
  "payload": {
    "field_a": 1,
    "field_b": 2,
    "field_c": 3
  }
}
```

\* Currently supports only integer value

Distributed under MIT license.
