# -*- coding: utf-8 -*-
# https://codewords.recurse.com/issues/one/an-introduction-to-functional-programming
import random


if __name__ == '__main__':
    squares = map(lambda x: x * x, xrange(10))
    print squares

    names = ['Mary', 'Isla', 'Sam']
    secret_names = map(
        lambda x: random.choice([
            'Mr. Pink',
            'Mr. Orange',
            'Mr. Blonde',
        ]),
        names
    )
    print secret_names

    secret_names = map(hash, names)
    print secret_names

    total = reduce(lambda a, x: a + x, xrange(10))
    print total

    sentences = [
        'Mary read a story to Sam and Isla.',
        'Isla cuddled Sam.',
        'Sam chortled.'
    ]
    sam_count = reduce(lambda a, x: a + x.count('Sam'), sentences, 0)
    print sam_count

    people = [
        {'name': 'Mary', 'height': 160},
        {'name': 'Isla', 'height': 80},
        {'name': 'Akira', 'height': 0},
        {'name': 'Sam'}
    ]

    heights = [p['height'] for p in people if 'height' in p and p['height'] > 0]
    if len(heights) > 0:
        print sum(heights) / len(heights)

    def move_cars(car_positions):
        return map(lambda x: x + 1 if random.random() > 0.3 else x, car_positions)

    def output_car(car_positions):
        return '-' * car_positions

    def run_step_of_race(state):
        return {'time': state['time'] - 1,
                'car_positions': move_cars(state['car_positions'])}

    def draw(state):
        print ''
        print '\n'.join(map(output_car, state['car_positions']))

    def race(state):
        draw(state)
        if state['time']:
            race(run_step_of_race(state))

    race(
        {'time': 5, 'car_positions': [1, 1, 1]}
    )

    def zero(s):
        if s[0] == '0':
            return s[1:]

    def one(s):
        if s[0] == '1':
            return s[1:]
