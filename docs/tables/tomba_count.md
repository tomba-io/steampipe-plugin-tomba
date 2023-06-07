# Table: tomba_count

Get detailed information about an total email addresses we have for one domain.

The `tomba_count` table requires the `domain` field to be specified in all queries, defining the Domain to lookup.

## Examples

### Info about an Email Count

```sql
select
  total as "Total Emails",
  personal_emails as "Total Personal Emails",
  generic_emails as "Total Generic Emails"
from
  tomba_count
where
  domain = 'tomba.io';
```

### Find the total number of emails in executive and sales department

```sql
select
  department ->> 'executive' as executive,
  department ->> 'sales' as sales
from
  tomba_count
where
  domain = 'clearbit.com';
```
