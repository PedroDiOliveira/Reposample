import random
#Introducao
print("|------------------------------------|")
print("|  Bem vindo ao jogo de advinhacao   |")
print("|------------------------------------|")

#Definindo a selecao de dificuldade
dificuldade = int(input("(1)Facil  (2)Medio  (3)Dificil: "))
print("*************************************************")
if (dificuldade > 3 or dificuldade < 1):
    print("O valor tem que ser entre 1 e 3")
elif(dificuldade == 1):
    y = 10
    z = 3
elif(dificuldade == 2):
    y = 20
    z = 4
elif(dificuldade == 3):
    y = 100
    z = 6

#Declaracao de variaveis    
rodada = 1
numero_secreto = random.randrange(1, y + 1)
total_de_rodadas = z
pontos = 100

#for que delimita o maximo de rodadas permitida por dificuldade
for rodada in range( 1,total_de_rodadas+1):                           #+1 pois a funcao range é excludente, ou seja precisamos para no numero seguinte ao total de rodadas
    print("Tentativa {} de {}".format(rodada, total_de_rodadas) )
    str = input("Digite um numero entre 1 e {}:".format(y))
    chute = int(str)

    if (chute < 1 or chute > y):
        print("Voce deve digitar um valor entre 1 e {}!".format(y))
        continue

    acertou = chute == numero_secreto
    maior = chute > numero_secreto
    menor = chute < numero_secreto

    if (acertou):
        print("Você ACERTOU e fez {}".format(pontos))
        break
    else:
        if (maior):
            print("Você errou! O seu chute foi MAIOR que o número secreto.")
            print("*************************************************")
        elif (menor):
            print("Você errou! O seu chute foi MENOR que o número secreto.")
            print("*************************************************")
        pontos_perdidos = abs(numero_secreto - chute)
        pontos = pontos - pontos_perdidos
if acertou == False:
        print("O numero secreto era {}".format(numero_secreto))
print("Fim do jogo!")