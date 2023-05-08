#Introducao
print("*******************************")
print("Bem vindo ao jogo de advinhacao")
print("*******************************")

#Declaracao de variaveis

numero_secreto = 8
x = int(input("Digite um numero"))

#Verificacao de acerto
while numero_secreto != x:
    print("voce errou")
    x = int(input("Digite um numero"))
    if x == numero_secreto:
        print("ACERTOU") 
    