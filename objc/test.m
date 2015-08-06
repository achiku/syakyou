#import <Foundation/Foundation.h>


@interface Person: NSObject

@property (readonly) NSString *nickname;
@property (readonly) NSString *firstname;
@property (readonly) NSString *lastname;
@property int age;

- (void) greeting;
@end


int main() {
    printf("hello, world!\n");
    return 0;
}
