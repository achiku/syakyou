#import <Foundation/Foundation.h>

// http://rypress.com/tutorials/objective-c/data-types/nsarray


int main() {
    NSLog(@"start\n");

    // literal
    NSArray *ary = @[@"moqada", @"_ideyuta", @"8maki"];
    NSLog(@"%@", ary);
    NSLog(@"%@", ary[0]);
    NSLog(@"%@", [ary objectAtIndex:2]);

    for (NSString *item in ary) {
        NSLog(@"%@", item);
    }

    for (int i = 0; i < [ary count]; i++) {
        NSLog(@"%d: %@", i, ary[i]);
    }

    if ([ary containsObject: @"achiku"]) {
        NSLog(@"achiku is in the array");
    } else {
        NSLog(@"achiku is not in the array");
    }
    return 0;
}
