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


```
+----------------------------------------------------------------------------------------------------------------------+---------------------------+---------------+
| uri                                                                                                                  | extracted_on              | still_on_page |
+----------------------------------------------------------------------------------------------------------------------+---------------------------+---------------+
| https://github.com/tomba-io/generic-emails/blob/084fc1a63d3cdaf9a34f255bedc2baea49a8e8b9/src/lib/validation/hash.ts  | 2021-02-08T20:09:54+01:00 | true          |
| https://github.com/tomba-io/generic-emails/blame/084fc1a63d3cdaf9a34f255bedc2baea49a8e8b9/src/lib/validation/hash.ts | 2021-02-08T20:09:59+01:00 | true          |
| https://katta.co/tomba-review/                                                                                       | 2021-06-02T13:41:50+01:00 | true          |
| https://shards.info/github/benemohamed/                                                                              | 2021-10-26T15:36:04+01:00 | true          |
| https://githubmemory.com/repo/tomba-io/python                                                                        | 2021-10-26T15:36:37+01:00 | true          |
| https://pypi.org/project/tomba-io/                                                                                   | 2021-10-26T15:37:41+01:00 | true          |
| https://www.npmjs.com/package/tomba                                                                                  | 2021-10-26T15:39:28+01:00 | true          |
| https://cnpmjs.org/package/tomba                                                                                     | 2021-10-26T18:58:53+01:00 | true          |
| https://packosphere.com/tombaio/tomba-meteor                                                                         | 2021-10-28T00:52:33+01:00 | true          |
| https://www.pkgstats.com/pkg:tomba                                                                                   | 2021-10-28T00:53:39+01:00 | true          |
| https://pub.dev/packages/tomba                                                                                       | 2021-10-30T15:28:18+01:00 | true          |
| https://pub.dev/packages/tomba/license                                                                               | 2021-10-30T15:28:19+01:00 | true          |
| https://pub.dev/packages/tomba/versions/1.0.0                                                                        | 2021-10-30T15:28:25+01:00 | true          |
| https://cnpmjs.org/package/tomba/v/1.0.0                                                                             | 2021-12-22T14:21:30+01:00 | true          |
| https://r.cnpmjs.org/tomba                                                                                           | 2021-12-22T14:23:16+01:00 | true          |
| https://github.com/tomba-io/r                                                                                        | 2021-12-22T14:41:25+01:00 | true          |
| https://rdrr.io/cran/tomba/f/README.md                                                                               | 2021-12-22T14:41:37+01:00 | true          |
| https://packosphere.com/tombaio/tomba-meteor/1.0.0                                                                   | 2021-12-22T14:42:21+01:00 | true          |
| https://githubmemory.com/repo/tomba-io/node                                                                          | 2021-12-22T15:44:03+01:00 | true          |
| https://packagist.org/packages/tomba-io/php                                                                          | 2023-01-06T02:46:36+01:00 | true          |
+----------------------------------------------------------------------------------------------------------------------+---------------------------+---------------+
```