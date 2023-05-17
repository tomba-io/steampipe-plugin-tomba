![image](https://hub.steampipe.io/images/plugins/tomba-io/tomba-social-graphic.png)

# tomba.io Plugin for Steampipe

Use SQL to search or verify lists of email addresses in minutes using [tomba.io](https://tomba.io).

- **[Get started →](https://hub.steampipe.io/plugins/tomba-io/tomba)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/tomba-io/tomba/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/tomba-io/steampipe-plugin-tomba/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install tomba-io/tomba
```

Configure the server address in `~/.steampipe/config/tomba.spc`:

```hcl
connection "tomba" {
    plugin = "tomba-io/tomba"
    # `key`: Tomba API Key.(Required)
    # `secret`: Tomba API secret. (Required)

    # Sign up for a free API and SECRET Keys at https://app.tomba.io/auth/register

    # This can also be set via the `TOMBA_KEY` and `TOMBA_SECRET` environment variables.
    
    # key = "ta_d40f89cc3b7f59a0638e1234a22fdfa9d3b86"
    # secret = "ts_121f9gf4-6f90-4856-9017-b12b5079adc9"
}
```

Or through environment variables:

```sh
export TOMBA_KEY=ta_d40f89cc3b7f59a0638e1234a22fdfa9d3b86
export TOMBA_SECRET=ts_121f9gf4-6f90-4856-9017-b12b5079adc9
```

Run steampipe:

```shell
steampipe query
```

enrich any Email:

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

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/tomba-io/steampipe-plugin-tomba.git
cd steampipe-plugin-tomba
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/tomba.spc
```

Try it!

```
steampipe query
> .inspect tomba
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/tomba-io/steampipe-plugin-tomba/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [tomba.io Plugin](https://github.com/tomba-io/steampipe-plugin-tomba/labels/help%20wanted)
