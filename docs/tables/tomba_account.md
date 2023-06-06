# Table: tomba_account

Get detailed information of current account.

## Examples

### Get the Account Information

```sql
select
  first_name,
  last_name,
  email,
  timezone,
  country,
  pricing 
from
  tomba_account;
```

### List the accounts created in the last 30 days

```sql
select
  first_name,
  last_name,
  email,
  timezone,
  country,
  pricing 
from
  tomba_account
where
  created_at >= now() - interval '30' day;
```
