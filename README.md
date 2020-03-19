# OrgGistScraper

OrgGistScraper makes it easy to scrape all Gists from members of a github organization. When bug-hunting one source of information is employee's gists, potentially containing log files and code snippets used on the target platform. 

## Installation

Use the golang package manager

```bash
go get -u github.com/AmIJesse/OrganizationMemberGistScraper
```

or git clone

```bash
git clone https://github.com/AmIJesse/OrganizationMemberGistScraper
```
Change into the directory you saved it in.

Then download the github package

```bash
go get -u
```

Build the file
```bash
go build
```
## Usage
If neither the -d or -o flags are set, it will only output to STDOUT
```bash
Usage of ./OrgGistScraper:
  -d    Automatically Download All Gists
  -o    Output Results To Text File (Enabled if -d flag is set)
  -org string
        Organization Name
```

## Example
GlassDoor was recently added to HackerOne, and has a short output so I'll use them as an example.
```bash
jesse@phoenix:~/OrgGistScraper$ ./OrgGistScraper -org Glassdoor -d
vikrama
-- test.html - https://gist.github.com/eefca776cbc74c9db59ce4ea098dfaee
-- SassMeister-output.css - https://gist.github.com/6a382b307f2a3a2282cb
-- SassMeister-input.scss - https://gist.github.com/6a382b307f2a3a2282cb
-- SassMeister-input.scss - https://gist.github.com/b930b7736f6b34fd02ee
-- SassMeister-output.css - https://gist.github.com/b930b7736f6b34fd02ee
jesse@phoenix:~/OrgGistScraper$ ls Glassdoor-2020-03-19/
output.txt  vikrama-SassMeister-input.scss  vikrama-SassMeister-output.css  vikrama-test.html
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.


## License
[MIT](https://choosealicense.com/licenses/mit/)
