---
organization: tomba-io
category: ["saas"]
icon_url: "/images/plugins/tomba-io/tomba.svg"
brand_color: "#1D6CE2"
display_name: "tomba.io"
short_name: "tomba"
description: "Steampipe plugin to query Domain or Email information from tomba.io."
og_description: "Query tomba.io with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/tomba-io/tomba-social-graphic.png"
---

# tomba.io + Steampipe

[Tomba](https://tomba.io) is a service that enables you to find and verify professional email addresses.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

## Enrichment

query any email

```sql
select
  email,
  first_name,
  last_name
from
  tomba_enrich
where
  email = 'b.mohamed@tomba.io';
```

```
+--------------------+------------+-----------+
| email              | first_name | last_name |
+--------------------+------------+-----------+
| b.mohamed@tomba.io | Mohamed    | Ben rebia |
+--------------------+------------+-----------+
```

## Documentation

- **[Table definitions & examples →](/plugins/tomba-io/tomba/tables)**

## Get started

### Install

Download and install the latest tomba.io plugin:

```bash
steampipe plugin install tomba-io/tomba
```

### Configuration

Installing the latest tomba plugin will create a config file (`~/.steampipe/config/tomba.spc`) with a single connection named `tomba`:

```hcl
connection "tomba" {
  plugin = "tomba-io/tomba"

  # `key`: Tomba API Key. (Required).
  # `secret`: Tomba API secret. (Required).
  # Sign up for a free API and Secret keys at https://app.tomba.io/auth/register  
  # This can also be set via the `TOMBA_KEY` and `TOMBA_SECRET` environment variables.  

  # key = "ta_d40f89cc3b7f59a0638e1234a22fdfa9d3b86"
  # secret = "ts_121f9gf4-6f90-4856-9017-b12b5079adc9"
}
```

- `key` - API key from tomba.io.
- `secret` - API secret from tomba.io.

Alternatively, you can also use the standard Tomba environment variables to obtain credentials **only if other arguments (`key`and `secret`) are not specified** in the connection:

```sh
export TOMBA_KEY=ta_d40f89cc3b7f59a0638e1234a22fdfa9d3b86
export TOMBA_SECRET=ts_121f9gf4-6f90-4856-9017-b12b5079adc9
```

## Get involved

- Open source: https://github.com/tomba-io/steampipe-plugin-tomba
- Community: [Slack Channel](https://steampipe.io/community/join)
