
  
  

This mini project has only one goal that to present some parts of my knowledge about Nodejs.

  

Please, to be aware that is not the best practice demo, because all practise can be discussed, and it depends mainly on you and your team power.

  

Used technologies (Please perform the necessary installations): makefile, m4, docker, node, npm, nodemon, toml, open-api, validator, express, prometheus client, log service, mocha, Nodejs clustering service.

  

## Description


Based on the original fizz-buzz, this api has one endpoint (see api.yaml)

Accepts five parameters : three integers int1, int2 and limit, and two strings str1 and str2.

Returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.
  
### For test

```bash
> make clean
> make cert-files; make install
> make test
```

### For try and dev

```bash
> make cert-files; make install (optional if already did)
> make run

```

### For build (prerequisite: make cert-files and make install)

```bash
> make build
```

### For deploy

```bash
> make deploy
```

### For release

```bash
make release
```

(This mini project comes from a job interview test)