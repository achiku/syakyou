#import <Foundation/Foundation.h>

// http://rypress.com/tutorials/objective-c/data-types/nsdictionary

int main() {
    NSLog(@"start");

    // Literal syntax
    NSDictionary *literalDict = @{
        @"moqada": @"Okada",
        @"8maki": @"Yamaki",
        @"_ideyuta": @"Ide",
    };

    // Values and keys as arguments
    NSDictionary *vkDict = [NSDictionary dictionaryWithObjectsAndKeys:
        [NSNumber numberWithInt: 0], @"8maki",
        [NSNumber numberWithInt: 1], @"_ideyuta",
        [NSNumber numberWithInt: 2], @"moqada", nil
    ];

    NSLog(@"%@", literalDict);
    NSLog(@"%@", vkDict);

    NSLog(@"%@", [literalDict allKeys]);
    NSLog(@"%@", [vkDict allValues]);

    NSLog(@"%@", literalDict[@"moqada"]);
    NSLog(@"%@", vkDict[@"moqada"]);

    NSLog(@"%ld", [literalDict count]);
    for (id key in literalDict) {
        NSLog(@"%@ -> %@", key, literalDict[key]);
    }

    NSMutableDictionary *mDict = [NSMutableDictionary dictionaryWithDictionary:
        @{
            @"moqada": @"frontend/backend/infra",
            @"_ideyuta": @"design/frontend",
            @"8maki": @"CEO",
        }
    ];

    NSLog(@"%@", mDict);

    [mDict setObject:@"design/frontend/backend" forKey: @"_ideyuta"];
    NSLog(@"%@", mDict);

    mDict[@"achiku"] = @"backend/infra";
    NSLog(@"%@", mDict);


    return 0;
}
