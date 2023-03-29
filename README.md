# Projet de Gautherot-Théo Essir22-prog-go

## Port Scanner

Dans le cadre de ma licence j'ai du réaliser un programme en Go qui va me permettre de faire un scanner de ports en Go qui peut scanner une plage de ports sur une adresse IP cible.

### Prerequisites

Tu auras besoin d'avoir déjà installer 
- Go (vérification avec ```go version```)
- go.mod
- go.sum

### Utilisation

Pour utiliser le scanner de ports, exécutez la commande suivante :

```
go run package-main.go --target <target-ip> --ports <ports-range>
```
`<target-ip>` : l'adresse IP cible à scanner.

`<ports-range>` : la plage de ports à scanner (exemples : 1024-65535, all).

You can run this script from the command line by using the go run command:

go run main.go package-main.go

On va initié un module ( appellé package en cours) 

go mod init ipscan
go mod tidy




Fonction expliqué 
Ce script utilise un package appelé "cobra" pour créer une commande de ligne appelée "portscan". Il définit plusieurs variables globales, telles que target, ports, workers, et quiet, qui sont utilisées pour spécifier les détails de l'analyse de ports à effectuer.

La commande racine, rootCmd, est définie en utilisant cobra.Command, et définit les options d'utilisation, la description courte et longue, ainsi que la fonction à exécuter lorsque la commande est appelée. La fonction Run appelle la fonction scanPorts en passant les variables globales comme arguments.

La fonction init() est appelée automatiquement lorsque le script est chargé. Elle utilise les fonctions de drapeau de Cobra pour définir les options de ligne de commande pour les variables globales définies précédemment.

rootCmd.Flags().StringVarP(&target, "target", "t", "", "the target IP to scan") définit une option de ligne de commande pour la variable target avec les étiquettes "t" et "target", une valeur par défaut vide et une description "the target IP to scan". La fonction MarkFlagRequired("target") indique que cette option de ligne de commande est obligatoire.

rootCmd.Flags().StringVarP(&ports, "ports", "p", "", "the range of ports to scan (examples: 1024-65535, all)") définit une option de ligne de commande pour la variable ports avec les étiquettes "p" et "ports", une valeur par défaut vide et une description "the range of ports to scan (examples: 1024-65535, all)". La fonction MarkFlagRequired("ports") indique que cette option de ligne de commande est obligatoire.

rootCmd.Flags().IntVarP(&workers, "workers", "w", 1, "the number of workers to use for scanning in parallel") définit une option de ligne de commande pour la variable workers avec les étiquettes "w" et "workers", une valeur par défaut de 1 et une description "the number of workers to use for scanning in parallel".

rootCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "don't log, only show results") définit une option de ligne de commande pour la variable quiet avec les étiquettes "q" et "quiet", une valeur par défaut de false et une description "don't log, only show results".

En résumé, cette fonction initialise les options de ligne de commande pour le script en définissant les options requises et facultatives pour l'analyse de ports.

parsePorts est une fonction qui prend en entrée une chaîne de ports et renvoie une liste triée d'entiers représentant les ports.
La fonction débute par initialiser une liste vide d'entiers appelée portsList.

Si la chaîne d'entrée est "all", la fonction parcourt tous les ports de 1 à 65535 et les ajoute à la liste portsList. Sinon, la chaîne d'entrée est divisée en deux parties en utilisant la fonction strings.Split avec le délimiteur "-".

Si la longueur de la liste de parties n'est pas égale à 2, cela signifie qu'il y a un problème avec la plage de ports donnée et la fonction génère une erreur avec log.Fatalf. Sinon, les deux parties sont converties en entiers avec strconv.Atoi.

Si la fin est inférieure au début, la plage de ports donnée est également considérée comme invalide et génère une erreur. Sinon, la fonction parcourt la plage de ports donnée de début à fin et ajoute chaque port à la liste portsList.

Enfin, la liste est triée avec sort.Ints et est renvoyée.
