# -*- coding: utf-8 -*-
import csv


if __name__ == '__main__':
    with open('./sampledata/u.data', 'r') as fh:
        fieldnames = ['user_id', 'item_id', 'rating', 'unixtime']
        critics = csv.DictReader(fh, fieldnames=fieldnames, delimiter='\t')
        result = []
        users = []
        for critic in critics:
            if not critic['user_id'] in users:
                users.append(critic['user_id'])
                result.append({critic['user_id']: [critic['item_id']]})
            else:
                u = [i for i in result if i.keys() == [critic['user_id']]][0]
                print u
