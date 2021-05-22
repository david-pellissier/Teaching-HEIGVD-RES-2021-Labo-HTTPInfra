# Load balancing

NGINX est facile à configurer pour du load balancing

Lien utile :
- https://nginx.org/en/docs/http/load_balancing.html


## Configuration

Il suffit d'ajouter deux sections `upstream` dans le fichier de config de nginx:
```
    upstream static_http {
        server 172.17.0.2:80;
        server 172.17.0.5:80;
    }
    
    upstream dynamic_http {
        server 172.17.0.3:80;
        server 172.17.0.6:80;
    }
```

Dans le reverse proxy, on peut alors référencer `static_http` et `dynamic_http`, et NGINX s'occupera de faire le load balancing:
```

    location / {
        proxy_pass http://static_http/;
    }

    location /api/fortunes/ {
        proxy_pass http://dynamic_http/;
    }
    
``` 

## Démo

Voilà le résultat des logs pour une requête vers le reverse_proxy sur les deux serveurs statiques: 

![](../figures/lb_proof.png)

On peut voir qu'ils reçoivent tous les deux des requêtes en provenance de l'adresse
172.17.0.4, qui est notre reverse proxy. Le load balancing est donc fonctionnel.