# Table: tomba_author

Get detailed information of email address from a blog post url.

**NOTES**:

* the `tomba_author` table
  requires the `url` field to be specified in all queries, defining the URL
  to lookup.

## Examples

### Info about an [Author Finder](https://tomba.io/author-finder)

query any blog post URL

```sql
select
  email, first_name, last_name
from
  tomba_author
where
  url = 'https://clearbit.com/blog/company-name-to-domain-api';
```

```
+-------------------+------------+-----------+
| email             | first_name | last_name |
+-------------------+------------+-----------+
| alex@clearbit.com | Alex       | Maccaw    |
+-------------------+------------+-----------+
```