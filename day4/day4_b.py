import hashlib

key = "bgvyzdsv"
integer = 0

while True:
        m = hashlib.md5()
        m.update(key)
        m.update(str(integer))
        hash = m.hexdigest()

        if hash[:6] == "000000":
                print("Lowest int to generate hash is: ", integer)
                break

        integer += 1
