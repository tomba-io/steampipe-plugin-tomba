# Table: tomba_linkedin

Get detailed information of email address from a Linkedin url.

The `tomba_linkedin` table requires the `url` field to be specified in all queries, defining the URL to lookup.

## Examples

### Info about an [Linkedin Finder](https://tomba.io/linkedin-finder)

query any Linkedin URL

```sql
select
   email,
   first_name,
   last_name 
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

### Find Enrichment Emails Sources

```sql
select
   p ->> 'uri' as uri,
   p ->> 'extracted_on' as extracted_on,
   p ->> 'still_on_page' as still_on_page 
from
   tomba_linkedin,
   jsonb_array_elements(sources) as p 
where
   url = 'https://www.linkedin.com/in/mohamed-ben-rebia';
```