# Table: tomba_linkedin

Get detailed information of email address from a Linkedin url.

**NOTES**:

* the `tomba_linkedin` table
  requires the `url` field to be specified in all queries, defining the URL
  to lookup.

## Examples

### Info about an [Linkedin Finder](https://tomba.io/linkedin-finder)

query any Linkedin URL

```sql
select
  email,first_name,last_name
from
  tomba_linkedin
where
  url = 'https://www.linkedin.com/in/alex-maccaw-ab592978';
```

```
+------------------+------------+-----------+
| email            | first_name | last_name |
+------------------+------------+-----------+
| alex@clearbit.co | Alex       | Alex      |
+------------------+------------+-----------+
```