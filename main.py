arreglo_1 = [int(input(f"Ingrese el número {i + 1} para el primer arreglo: ")) for i in range(10)]

arreglo_2 = [int(input(f"Ingrese el número {i + 1} para el segundo arreglo: ")) for i in range(15)]

print("Primer arreglo:", arreglo_1)

# Mostrar el segundo arreglo
print("Segundo arreglo:", arreglo_2)

arreglo_intercalado = []

for elem_1, elem_2 in zip(arreglo_1, arreglo_2):
    arreglo_intercalado.extend([elem_1, elem_2])

if len(arreglo_1) < len(arreglo_2):
    arreglo_intercalado.extend(arreglo_2[len(arreglo_1):])
elif len(arreglo_1) > len(arreglo_2):
    arreglo_intercalado.extend(arreglo_1[len(arreglo_2):])

print("Arreglo intercalado:", arreglo_intercalado)
arreglo_intercalado.sort()
print("Arreglo intercalado ordenado:",arreglo_intercalado )
