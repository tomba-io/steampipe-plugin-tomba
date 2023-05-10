# Table: tomba_account

Get detailed information of current account.

**NOTES**:

* the `tomba_account` table

## Examples

### Info about an Account Information

query

```sql
select
  *
from
  tomba_account;
```

```
+------------+-----------+--------------------+-----------------------------------------+-----------------------------+
| first_name | last_name | email              | secret_token                            | _ctx                        |
+------------+-----------+--------------------+-----------------------------------------+-----------------------------+
| Mohamed    | Ben       | b.mohamed@tomba.io | ts_xxxx                                 | {"connection_name":"tomba"} |
+------------+-----------+--------------------+-----------------------------------------+-----------------------------+
```