
f = open("day1_input.txt", 'r')
raw = f.read()

braces = list(raw)
floor = 0
position = 0;

for b in braces:
        if b == '(':
                floor = floor + 1
        elif b == ')':
                floor = floor - 1
        elif b == '\n':
                break
        else:
                print "BAD INPUT. BAD."
                exit()

        position = position + 1
        if floor < 0:
                break

print position
