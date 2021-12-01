# https://adventofcode.com/2020/day/13

testInput	= "../test.txt"
testInput2 = "../test2.txt"
mainInput	= "../main.txt"

def divmod(a, b):
	return a // b, a % b

def egcd(a, b):
	old_r, r = a, b
	old_s, s = 1, 0
	old_t, t = 0, 1

	while r != 0:
		quo, rem = divmod(old_r, r)
		old_r, r = r, rem
		old_s, s = s, old_s - quo * s
		old_t, t = t, old_t - quo * t
	return old_r, old_s, old_t

def alignment(a, b, pa, pb):
	gcd, s, _ = egcd(a, b)
	diff = pa - pb
	pdMult, pdRem = divmod(diff, gcd)
	if pdRem != 0:
		raise("phase mismatch")
	period = a / gcd * b
	phase = pa - pdMult * s * (a % period)
	return period, phase

def part2(buses, offsets):
	period = buses[0]
	offset = 0
	i = 1
	while i < len(buses):
		period, phase = alignment(period, buses[i], offset, offsets[i])
		offset = (period + phase) % period
		i+=1
	return period - offset

if __name__ == '__main__':
	lines = []
	with open(mainInput, 'r') as f:
		lines = f.readlines()

	t0 = int(lines[0])
	for line in lines[1:]:
		busesStr = line.split(",")
		buses = []
		offsets = []
		for i, bus in enumerate(busesStr):
			if bus == "x":
				continue
			buses.append(int(bus))
			offsets.append(i)
		print(part2(buses, offsets))