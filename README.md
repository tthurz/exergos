# exergos

Instruções para usar o exergos.

## Como inicializar um repositorio Go

    go mod init nome-do-repositorio

Exemplo:

    go mod init github.com/tthurz/exergos

## Como compilar tudo

   go install ./...

## Como usar o git

Pode fazer em cada arquivo:

    git add README.md
    git add cmd\teste\main.go

Ou pode fazer em tudo:

    git add . 

Como criar um commit:

    git commit -m "Mensagem"

Enviar para o repo:

    git push

### Para baixar alterações

Repositório todo:

    git clone https://github.com/tthurz/exergos

Apenas atualiar um repo:

    git pull


    
