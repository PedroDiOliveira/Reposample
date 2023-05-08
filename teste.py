x = int(input("digite o valor de x:"))
y = int(input("digite o valor de y:"))
z = int(input("digite o valor de z:"))

if x < y and x < z:
    print(x," é o menor numero!")
elif y < x and y < z:
     print(y," é o menor numero!")
else:
     print(z," é o menor numero!")