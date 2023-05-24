# Table: tomba_verifier

Get detailed information of deliverability of an email address.

The `tomba_verifier` table requires the `email` field to be specified in all queries, defining the Email to lookup.

## Examples

### Info about an [Email Verifier](https://tomba.io/email-verifier)

query any b2b email

```sql
select
   email,
   status,
   result,
   score,
   registrar_name,
   referral_url
from
   tomba_verifier 
where
   email = 'b.mohamed@tomba.io';
```