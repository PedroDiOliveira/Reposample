#titulo
print("***************************************")
print("Bem vindo ao jogo de advinhação!")
print("***************************************")

#Declaração de variaveis
numero = 8
while True:
    resposta = int(input("Digite um numero:"))

#Comparação de valores
    if numero == resposta:
        print("Voce digitou", numero,", e a resposta era", resposta)
        print("Parabens!!")
        break
    else:
        print("Voce errou, tente novamente")
    