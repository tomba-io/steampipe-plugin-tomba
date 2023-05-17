# Table: tomba_enrich

Get detailed information person and company data based on an email, For example, you could retrieve a person’s name, location and social handles from an email.

The `tomba_enrich` table requires the `email` field to be specified in all queries, defining the Email to lookup.

## Examples

### Info about an [Enrichment](https://tomba.io/enrichment)

query any b2b email

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