# Projet de Gautherot-Théo Essir22-prog-go

## Port Scanner

Dans le cadre de ma licence j'ai du réaliser un programme en Go qui va me permettre de faire un scanner de ports qui peut scanner une plage de ports sur une adresse IP cible.

### Prérequis

Tu auras besoin d'avoir déjà installer :
- Go (vérification avec ```go version```)
Pour les deux fichiers suivant il sont installé par défaut dans mon Git 
- go.mod ( ou ```go mod init ipscan``` pour re-génerer)
- go.sum (ou ```go mod tidy``` depuis l'export pour mettre à jour)

### Utilisation

Pour utiliser le scanner de ports, exécutez la commande suivante :

```
go run package-main.go --target <target-ip> --ports <ports-range>
```

`<target-ip>` : l'adresse IP cible à scanner.

`<ports-range>` : la plage de ports à scanner (exemples : 1024-65535, all).


## Exemples

Pour scanner tous les ports sur l'adresse IP localhost, exécutez la commande suivante :

```
port-scanner --target localhost --ports all
```

## Options script

`--target` ou `-t` : l'adresse IP de la machine à scanner (obligatoire)
`--ports` ou `-p` : la plage de ports à scanner (obligatoire). Cette option peut prendre la valeur all pour scanner tous les ports.
`--workers` ou `-w` : le nombre de threads à utiliser pour le scan (par défaut: 1)
`--quiet` ou `-q` : ne pas afficher les logs, afficher uniquement les résultats (par défaut: false)

## Auteur

* **Gautherot Théo** - *Initial work* - [PurpleBooth](https://github.com/AsTheLow)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
