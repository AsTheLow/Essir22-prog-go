# Essir22-prog-go
Projet de Gautherot-Théo 
You can run this script from the command line by using the go run command:

go run main.go https://www.example.com

On va initié un module ( appellé package en cours) 

go mod init ipscan
go mod tidy


Fonction expliqué 

func parsePorts(ports string) []int {
	var portsList []int
	if ports == "all" {
		for i := 1; i <= 65535; i++ {
			portsList = append(portsList, i)
		}
	} else {
		parts := strings.Split(ports, "-")
		if len(parts) != 2 {
			log.Fatalf("Invalid port range: %s", ports)
		}
		start, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("Invalid start port: %s", parts[0])
		}
		end, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("Invalid end port: %s", parts[1])
		}
		if end < start {
			log.Fatalf("Invalid port range: %s", ports)
		}
		for i := start; i <= end; i++ {
			portsList = append(portsList, i)
		}
	}
	sort.Ints(portsList)
	return portsList
}
parsePorts est une fonction qui prend en entrée une chaîne de ports et renvoie une liste triée d'entiers représentant les ports.
La fonction débute par initialiser une liste vide d'entiers appelée portsList.

Si la chaîne d'entrée est "all", la fonction parcourt tous les ports de 1 à 65535 et les ajoute à la liste portsList. Sinon, la chaîne d'entrée est divisée en deux parties en utilisant la fonction strings.Split avec le délimiteur "-".

Si la longueur de la liste de parties n'est pas égale à 2, cela signifie qu'il y a un problème avec la plage de ports donnée et la fonction génère une erreur avec log.Fatalf. Sinon, les deux parties sont converties en entiers avec strconv.Atoi.

Si la fin est inférieure au début, la plage de ports donnée est également considérée comme invalide et génère une erreur. Sinon, la fonction parcourt la plage de ports donnée de début à fin et ajoute chaque port à la liste portsList.

Enfin, la liste est triée avec sort.Ints et est renvoyée.
