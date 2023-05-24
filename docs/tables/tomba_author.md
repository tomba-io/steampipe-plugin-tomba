# Table: tomba_author

Get detailed information of email address from a blog post url.

The `tomba_author` table requires the `url` field to be specified in all queries, defining the URL to lookup.

## Examples

### Info about an [Author Finder](https://tomba.io/author-finder)

query any blog post URL

```sql
select
   email,
   first_name,
   last_name 
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

### Find Author Finder Emails Sources

```sql
select
   p ->> 'uri' as uri,
   p ->> 'extracted_on' as extracted_on,
   p ->> 'still_on_page' as still_on_page 
from
   tomba_author,
   jsonb_array_elements(sources) as p 
where
   url = 'https://clearbit.com/blog/company-name-to-domain-api';
```