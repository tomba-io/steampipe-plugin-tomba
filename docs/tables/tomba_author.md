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

```
+------------------------------------------------------------------------------------------------------------------+---------------------------+---------------+
| uri                                                                                                              | extracted_on              | still_on_page |
+------------------------------------------------------------------------------------------------------------------+---------------------------+---------------+
| https://clearbit.com/                                                                                            | 2021-01-12T12:53:02+01:00 | true          |
| https://clearbit.com/prospector                                                                                  | 2021-01-12T12:53:03+01:00 | true          |
| https://clearbit.com/pardot                                                                                      | 2021-01-12T12:53:03+01:00 | true          |
| https://pypi.org/project/clearbit/                                                                               | 2021-01-29T20:32:08+01:00 | true          |
| https://growjo.com/company/Clearbit                                                                              | 2021-07-23T09:11:15+01:00 | true          |
| https://help.clearbit.com/hc/en-us/articles/360019332974-dashdash-Setup-Guide                                    | 2021-07-25T07:53:54+01:00 | true          |
| https://pdfcookie.com/documents/book1xlsx-025ekom5x721                                                           | 2021-10-01T14:53:49+01:00 | true          |
| https://github.com/clearbit/clearbit-ruby/blob/master/clearbit.gemspec                                           | 2021-12-22T15:54:30+01:00 | true          |
| https://github.com/clearbit/clearbit-python                                                                      | 2021-12-22T15:54:33+01:00 | true          |
| https://clearbitjs.com/hubspot                                                                                   | 2021-12-22T21:54:13+01:00 | true          |
| https://clearbitjs.com/marketo                                                                                   | 2021-12-22T21:54:13+01:00 | true          |
| https://clearbitjs.com/pardot                                                                                    | 2021-12-22T21:54:13+01:00 | true          |
| https://dashboard.clearbit.com/docs                                                                              | 2021-12-22T13:44:15+01:00 | true          |
| https://dashboard.clearbit.com/prospector                                                                        | 2021-12-22T13:44:18+01:00 | true          |
| https://dashboard.clearbit.com/hubspot                                                                           | 2021-12-22T13:44:20+01:00 | true          |
| https://www.skymem.info/srch?q=https://app.ahrefs.com/v2-site-explorer/refdomains/subdomains?target=clearbit.com | 2022-09-05T15:01:00+01:00 | true          |
| https://www.skymem.info/srch?q=https://clearbit.com/                                                             | 2022-09-16T19:57:19+01:00 | true          |
| https://newwebsiteinterface.tomba.workers.dev/api/email-finder                                                   | 2022-09-17T02:20:25+01:00 | true          |
| https://newwebsiteinterface.tomba.workers.dev/api/email-verifier                                                 | 2022-09-17T02:20:25+01:00 | true          |
| https://newwebsiteinterface.tomba.workers.dev/api/author-finder                                                  | 2022-09-17T02:20:25+01:00 | true          |
+------------------------------------------------------------------------------------------------------------------+---------------------------+---------------+
```