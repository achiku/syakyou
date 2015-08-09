#import <Foundation/Foundation.h>


@interface Person: NSObject

@property (assign) NSString *nickname;
@property (assign) NSString *firstname;
@property (assign) NSString *lastname;
@property int age;

- (void) greeting;

@end

@implementation Person
- (id) init:(NSString *)nickname : (NSString *)firstname : (NSString *)lastname : (int) age {
    self = [super init];
    self.nickname = nickname;
    self.firstname = firstname;
    self.lastname = lastname;
    self.age = age;
    return self;
}

- (void) greeting {
    NSLog(@"Hello! My nickname is %@!\n", self.nickname);
}

- (void) politeGreeting {
    NSLog(@"Good evening. My name is %@, %@.\n", self.firstname, self.lastname);
}
@end



int main() {
    Person *p = [[Person alloc] init:@"moqada" :@"Masahiko" :@"Okada" :29];
    [p greeting];
    [p politeGreeting];
    return 0;
}
