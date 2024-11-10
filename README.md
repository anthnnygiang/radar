# Domain Radar

Sourced from [Cloudflare](https://radar.cloudflare.com/domains). Data last updated: 2024/11/04.

## Usage
0. (Optional) Download the lists from [Cloudflare](https://radar.cloudflare.com/domains), copy into `data/`, and rename them into format `top-<number>-domains.csv`.
1. Run `go run main.go`.
2. View result in `result.txt`

## About

Lists of the most popular 1 million domains.

Cloudflare generates a lists of the top:
* 200
* 500
* 1,000
* 2,000
* 5,000
* 10,000
* 20,000
* 50,000
* 100,000
* 200,000
* 500,000 and 
* 1,000,000 domains.

This program collates all the lists, and arranges
them into groups of 
* 1 - 200
* 201 - 500
* 501 - 1,000
* 1,001 - 2,000
* 2,001 - 5,000
* 5,001 - 10,000
* 10,001 - 20,000
* 20,001 - 50,000
* 50,001 - 100,000
* 100,001 - 200,000
* 200,001 - 500,000 and
* 500,001 - 1,000,000.
