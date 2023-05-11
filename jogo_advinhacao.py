import random
#Introducao
print("*******************************")
print("Bem vindo ao jogo de advinhacao")
print("*******************************")

#Declaracao de variaveis
rodada = 1

dificuldade = int(input("Digite 1 para FACIL, 2 para MEDIO e 3 para DIFICIL: "))
print("*************************************************")
if (dificuldade > 3 or dificuldade < 1):
    print("O valor tem que ser entre 1 e 3")
elif(dificuldade == 1):
    y = 10
    z = 3
elif(dificuldade == 2):
    y = 20
    z = 5
elif(dificuldade == 3):
    y = 100
    z = 10

#Declaracao de variaveis    
rodada = 1
numero_secreto = random.randrange(1, y + 1)
total_de_rodadas = z

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
        print("Você acertou!")
        break
    else:
        if (maior):
            print("Você errou! O seu chute foi MAIOR que o número secreto.")
            print("*************************************************")
        elif (menor):
            print("Você errou! O seu chute foi MENOR que o número secreto.")
            print("*************************************************")

print("Fim do jogo!")