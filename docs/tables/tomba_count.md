# Table: tomba_count

Get detailed information about an total email addresses we have for one domain.

**NOTES**:

* the `tomba_count` table
  requires the `domain` field to be specified in all queries, defining the Domain
  to lookup.

## Examples

### Info about an Email Count

query any domain

```sql
select
  *
from
  tomba_count
where
  domain = 'tomba.io';
```

```
+----------+-------+-----------------+----------------+-----------------------------+
| domain   | total | personal_emails | generic_emails | _ctx                        |
+----------+-------+-----------------+----------------+-----------------------------+
| tomba.io | 20    | 7               | 13             | {"connection_name":"tomba"} |
+----------+-------+-----------------+----------------+-----------------------------+
```