def input_list(prompt=None, delimiter=None):
    return input(prompt).split(delimiter)


def input_int_list(prompt=None, delimiter=None):
    return [int(i) for i in input_list(prompt, delimiter)]


def input_int(prompt):
    return int(input(prompt))
