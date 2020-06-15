import multiprocessing

def SUMA(p,a,b,c):
    print(f"Posicion de arreglo: {p} La suma es: {a+b+c}")
    return

if __name__ == '__main__':
    Matriz1 = [34,55,66,5,87,5]
    Matriz2 = [11,33,44,54,33,3]
    Matriz3 = [5,7,53,66,7,4]

    for i in range(6):
        p = multiprocessing.Process(target=SUMA, args=(i,Matriz1[i],Matriz2[i],Matriz3[i]))
        p.start()