from sklearn.preprocessing import LabelEncoder
import pandas as pd


if __name__ == '__main__':
    df1 = pd.DataFrame({
        'data1': range(6),
        'key1': ['A', 'B', 'C', 'A', 'B', 'C'],
        'key2': ['A', 'B', 'C', 'A', 'B', 'C']
        })
    df2 = pd.DataFrame({
        'data2': range(3),
        'key1': ['A', 'C', 'D'],
        'key2': ['A', 'E', 'B']
        })

    le = LabelEncoder()
    le.fit(df1['key1'].unique())
    print(le.transform(df1['key1']))

    df1['label'] = le.transform(df1['key1'])
    df1.drop('key1', axis=1, inplace=True)
    print(df1)
