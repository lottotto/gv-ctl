# gvctl
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Flottotto%2Fgv-ctl.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Flottotto%2Fgv-ctl?ref=badge_shield)


### usage
```
gvctl setup -f sample.yaml
```

### yaml file 
```
"baseurl": <YOUR GITLAB SERVER URL>
"id": <PROJECT or GROUP ID>
"token": <YOUR ACCESS TOKEN>
"type": <"project" or "group">
"variables":
- "key": "HOGE"
  "value": "hoge"
- "key": "FUGA"
  "value": "fuga"
```

## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Flottotto%2Fgv-ctl.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Flottotto%2Fgv-ctl?ref=badge_large)