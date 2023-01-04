// Lista de Flags para usar en paquete fmt

// Generales

// %v es el valor por defecto
// %#v Una sintaxis de go para la representación de un valor
// %T Para mostrar el tipo de dato
// Sin valor

// Booleanos

// %t true o false

// Enteros

// %b base 2
// %c carácter representado por su correspondiente valor unicode
// %d base 10
// %o base 8
// %O base 8 con prefijo 0o
// %q Citas textuales entre comillas
// %x base 16 lowercase
// %X base 16 UPPERCASE
// %U Formato unicode U+1234 al igual que "U+%04X"

// Flotantes y componentes sustitutos

// %b decimales con notación científica de la forma strconv.FormatFloat -> -123456p-78
// %e notación científica -1.234456e+78
// %E notación científica -1.234456E+78
// %f decimal con punto sin exponente 123.456
// %F sinónimo de %f
// %g %e para exponentes grandes, %f de lo contrario. La precision se discute después
// %G %E para exponentes grandes, %F de lo contrario.
// %x notación hexadecimal -0x1.23abcp+20
// %X Hexadecimal en mayúscula -0X1.23ABCP+20

// Strings



