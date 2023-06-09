# Table: tomba_enrich

Get detailed information person and company data based on an email, For example, you could retrieve a person’s name, location and social handles from an email.

The `tomba_enrich` table requires the `email` field to be specified in all queries, defining the Email to lookup.

## Examples

### Info about an [Enrichment](https://tomba.io/enrichment)

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

### Find Enrichment Email Sources

```sql
select
  p ->> 'uri' as uri,
  p ->> 'extracted_on' as extracted_on,
  p ->> 'still_on_page' as still_on_page
from
  tomba_enrich,
  jsonb_array_elements(sources) as p
where
  email = 'b.mohamed@tomba.io';
```

### Get the social media handles of an email

```sql
select
  email,
  first_name,
  last_name,
  twitter,
  linkedin
from
  tomba_enrich
where
  email = 'b.mohamed@tomba.io';
```
