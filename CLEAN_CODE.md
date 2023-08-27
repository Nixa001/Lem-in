<h1 align="center">CLEAN-CODE-GOLANG</h1>

<table align="center" border="0">
  <tr>
    <td ><a href="https://github.com/Nixa001"><img src="https://avatars.githubusercontent.com/u/85998620?v=4" alt="zone01" border="0" width="100px"></a></td>
    <td><a href="#"><img src="https://i.ibb.co/tL6SpfD/zone01.png" alt="zone01" border="0" width="100px"></a></td>
    <td><a href="#"><img src="https://i.ibb.co/W2vM8Ws/golang.png" alt="golang" border="0" width="120px"></a></td>
   </tr>
</table>


 Les principes du Clean Code appliqués en Golang pour faciliter la lecture, la maintenance et la compréhension du code.



## Nommage
- Utiliser des noms significatifs et descriptifs pour les variables, les fonctions, les types, etc.
- Les noms doivent refléter leur fonction ou leur rôle dans le programme.
- Éviter les noms abrégés ou trop courts, qui peuvent rendre le code difficile à comprendre.

Exemples :
```go
// Mauvais nommage
var v int // "v" est un nom trop court et peu descriptif

// Bon nommage
var totalScore int // "totalScore" est plus descriptif

// Mauvais nommage
func clc(num1 int, num2 int) int { // "clc" est un nom abrégé peu explicite
    return num1 + num2
}

// Bon nommage
func calculateSum(num1 int, num2 int) int { // "calculateSum" est plus explicite
    return num1 + num2

// Mauvais nommage
type s struct { // "s" est un nom de type trop court et peu descriptif
    name string
}

// Bon nommage
type person struct { // "person" est plus descriptif
    name string
}
```
## Fonctions
- Les fonctions doivent être courtes et faire une seule chose.
- Les noms de fonctions doivent être significatifs et refléter leur rôle dans le programme.
- Les fonctions ne doivent pas avoir d'effets secondaires indésirables.

Exemples :
```go
// Fonction avec trop de responsabilités
func processInput(input string) {
    // Valide l'entrée
    // Traite l'entrée
    // Enregistre les résultats
}

// Fonction avec une seule responsabilité
func processInput(input string) error {
    if len(input) == 0 {
        return errors.New("Entrée vide")
    }
    // Traite l'entrée
    return nil
}
```

## Gestion des erreurs
- Les erreurs doivent être renvoyées à l'appelant de manière appropriée, par exemple en utilisant la valeur de retour de la fonction ou en utilisant le paquetage errors.
- Les erreurs doivent être claires et informatives pour aider à diagnostiquer les problèmes.

Exemples :
```go
Copy code
// Traitement d'une erreur de manière incorrecte
func processInput(input string) string {
    if len(input) == 0 {
        return "Erreur: Entrée vide" // renvoie une chaîne, pas une erreur
    }
    // Traitement de l'entrée
}

// Traitement correct d'une erreur
func processInput(input string) (string, error) {
    if len(input) == 0 {
        return "", errors.New("Entrée vide") // renvoie une erreur, pas une chaîne
    }
    // Traitement de l'entrée
}
```
## Structure
- Les programmes doivent être organisés de manière logique pour faciliter la lecture et la maintenance.
- Les paquetages doivent être organisés pour refléter leur fonction ou leur rôle dans le programme.
- Les fichiers doivent être organisés de manière logique et nommés de manière significative.

Exemples :
```go
mon-programme/
├── main.go
├── monpaquetage/
│   ├── foo.go
│   ├── bar.go
│   └── baz_test.go
└── README.md
```
## Tests
- Les tests doivent être écrits pour chaque fonction ou méthode importante.
- Les tests doivent être organisés de manière logique et nommés de manière significative.
- Les tests doivent être clairs et faciles à comprendre.

Exemples :
```go
// Fonction importante sans tests
func calculateTotalScore(scores []int) int {
    totalScore := 0
    for _, score := range scores {
        totalScore += score
    }
    return totalScore
}

// Fonction importante avec tests
func calculateTotalScore(scores []int) int {
    totalScore := 0
    for _, score := range scores {
        totalScore += score
    }
    return totalScore
}

func TestCalculateTotalScore(t *testing.T) {
    scores := []int{10, 20, 30, 40}
    expectedTotalScore := 100
    actualTotalScore := calculateTotalScore(scores)
    if actualTotalScore != expectedTotalScore {
        t.Errorf("calculateTotalScore(%v) = %v; want %v", scores, actualTotalScore, expectedTotalScore)
    }
}
```

## Variables
- Les variables doivent être nommées de manière significative et refléter leur utilisation dans le programme.
- Les variables doivent avoir une portée minimale, c'est-à-dire qu'elles doivent être déclarées dans le bloc le plus petit possible.
- Les variables doivent être initialisées dès leur déclaration.

Exemples :
```go

// Mauvaise utilisation de variables
func calculateAverage(numbers []int) float64 {
    var sum int
    for _, number := range numbers {
        sum += number
    }
    // Utilise une variable globale pour stocker la moyenne
    average = float64(sum) / float64(len(numbers))
    return average
}

// Bonne utilisation de variables
func calculateAverage(numbers []int) float64 {
    sum := 0
    for _, number := range numbers {
        sum += number
    }
    average := float64(sum) / float64(len(numbers))
    return average
}
```
## Structures de données
- Les structures de données doivent être nommées de manière significative et refléter leur utilisation dans le programme.
- Les structures de données doivent avoir des champs avec des noms significatifs.
- Les méthodes qui opèrent sur une structure de données doivent être définies dans le même paquet que la structure elle-même.

Exemples :
```go
// Mauvaise utilisation de structures de données
type Employee struct {
    Name string
    Age  int
}

func GetEmployeeAge(employee Employee) int {
    return employee.Age
}

// Bonne utilisation de structures de données
type Employee struct {
    Name string
    Age  int
}

func (e Employee) GetAge() int {
    return e.Age
}
```

## Commentaires
- Les commentaires doivent être utilisés avec parcimonie et uniquement lorsque cela est nécessaire pour clarifier le code.
- Les commentaires ne doivent pas être utilisés pour expliquer du code mal écrit. Si le code est difficile à comprendre, modifiez-le plutôt que de le commenter.
- Les commentaires doivent être utilisés pour expliquer le « pourquoi » du code plutôt que le « comment ».
- Les commentaires doivent être clairs et informatifs, décrivant le but ou le comportement du code.

Exemples :
```go
Copy code
// Commentaire inutile
func add(x, y int) int {
    // Ajoute x et y et retourne le résultat
    return x + y
}

// Commentaire utile
func add(x, y int) int {
    // Ajoute x et y
    result := x + y
    // Retourne le résultat de l'addition
    return result
}
```
## Formatage
- Le code doit être formaté de manière cohérente pour faciliter la lecture et la compréhension.
- Les outils de formatage de code tels que gofmt ou goimports doivent être utilisés pour garantir la cohérence.

Exemples :
```go

// Code mal formaté
func calculateAverage(numbers []int) float64{
    sum:=0
    for _,number:=range numbers {
        sum+=number
    }
    average:=float64(sum)/float64(len(numbers))
    return average
}

// Code bien formaté
func calculateAverage(numbers []int) float64 {
    sum := 0
    for _, number := range numbers {
        sum += number
    }
    average := float64(sum) / float64(len(numbers))
    return average
}
```
## Conclusion
En suivant ces principes du Clean Code, nous pouvons écrire des programmes Go plus propres, plus lisibles et plus faciles à comprendre et à maintenir. Utilisez ces exemples pour vous guider dans votre propre code et améliorez la qualité de vos programmes.


## Authors

- [@Nixa001](https://github.com/Nixa001)


```
 _   _   _                                     $
| \ | | (_)                  ___     ___    _  $
|  \| |  _  __  __   __ _   / _ \   / _ \  / | $
| . ` | | | \ \/ /  / _` | | | | | | | | | | | $
| |\  | | |  >  <  | (_| | | |_| | | |_| | | | $
|_| \_| |_| /_/\_\  \__,_|  \___/   \___/  |_| $
                                               $
                                               $
```

<table align="center" border="0">
  <tr>
    <td ><a href="https://github.com/Nixa001"><img src="https://avatars.githubusercontent.com/u/85998620?v=4" alt="zone01" border="0" width="100px"></a></td>
    <td><a href="https://imgbb.com/"><img src="https://i.ibb.co/tL6SpfD/zone01.png" alt="zone01" border="0" width="100px"></a></td>
    <td><a href="https://imgbb.com/"><img src="https://i.ibb.co/W2vM8Ws/golang.png" alt="golang" border="0" width="120px"></a></td>
   </tr>
</table>

