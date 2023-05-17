# Table: tomba_account

Get detailed information of current account.

## Examples

### Info about an Account Information

query

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

```
+------------+-----------+--------------------+----------------+---------+------------+
| first_name | last_name | email              | timezone       | country | pricing    |
+------------+-----------+--------------------+----------------+---------+------------+
| Mohamed    | Ben       | b.mohamed@tomba.io | Africa/Algiers | DZ      | Enterprise |
+------------+-----------+--------------------+----------------+---------+------------+
```