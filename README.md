<center> <h1>Micromango</h1> </center>

## abstract

Micro (services) man (me and you) go (lang) - is my little pet project to train code and learn how to do things. I really love to read manga and keep lists. So I decided that write my own manga catalog and reader is a good idea. 

## how to run

Anyway the first of all you have to clone this repository. 

```console
foo@bar/$ git clone https://github.com/cl1ckname/micromango
```

then  create `.env` file. An example of it you can find in `.env-example`. There is no need to change something in it, but
you can rewrite ports or paths to databases

### run using compose

1. Clone client repository

    ```console
    foo@bar/$ git clone https://github.com/cl1ckname/micromango-client 
    ```

    or just
    
    ```console
    foo@bar/$ make client
    ```
2. Build and run compose set

    ```console
    foo@bar/$ docker compose up -d
    ```
3. have fun

### run using pm2

1. Install pm2 if you don't have
    
    ```console
   foo@bar/$ npm i -g pm2
    ```

2. Run cluster using
   
   ```console
   foo@bar/$ pm2 start ecosystem.config.js
   ```
   
   or using `Makefile`

   ```console
   foo@bar/$ make up
   ```