    git tag -a v0.1.1 -m "Nueva versión con el comando init"
    git push origin v0.1.1
    go run github.com/diego-all/peregrinus@v0.1.1 init --name Diego

    git show-ref --tags
    
    git tag -d v0.1.1
    git push origin --delete v0.1.1