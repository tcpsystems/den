# Market research

This is research for a later plan. It does not create accounts, trials, or servers.

## Row

A price row has provider, offer, region, price, billing unit, date, and source URL. A claim about SEO, certificates, terms, privacy, design, discount, or trial keeps the same source. A page without a price and date is a lead.

## Groups

GPU capacity, VPS and bare metal, and web or container platforms are separate groups. One worker owns one group. Two workers must record the same price and unit before the row is accepted.

## Radar

Accepted rows are presented in Radar as offers. They do not enter the model `Price` record. An offer has provider, offer, region, price, billing unit, observed date, source URL, and a content hash. Existing Radar records are assessed against those fields. A record missing a source, date, or unit stays unaccepted and visible as a gap.

## Use

Accepted rows can inform a Den, Radar, or Foxcore plan. They do not change the runtime loop or start a deployment.
