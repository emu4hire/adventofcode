lines = [line.rstrip('\n') for line in open("day2_input.txt", 'r')]
total = 0

for line in lines:
        x, y, z = line.split('x')
        x = eval(x)
        y = eval(y)
        z = eval(z)
        this_present = min(2*(x+y), 2*(x+z), 2*(y+z)) + (x*y*z)
        total = total + this_present

print total
