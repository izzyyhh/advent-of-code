def cartesianProduct(els, slots=2):
    if slots == 1:
        return els

    ops = els * slots;
    cartesianProduct = []

    for a in ops:
        for b in ops:
            cartesianProduct.append((tuple(a)+tuple(b)))

    res = set()
    for a in ops:
        for b in cartesianProduct:
            res.append((tuple(a)+tuple(b)))

    return res

test = cartesianProduct(["+", "*"], 3)
for t in test:
    print(t)