factory method hay còn gọi là Virtual Constructor

Factory Method là một Creational Design Pattern, nó cung cấp một interface cho việc tao ra Object trong
một superclass, NHƯNG cho phép các subclass có thể thay đổi kiểu của Object mà nó sẽ tạo

Chấp nhận factory method vi phạm nguyên lý O

Problem:

    Với strategy như ở main.go thì developer có thể tự do tạo struct mới như dạng EmailNotifier
    hoặc SMSNotifier. 
    Điều này dẫn đến 1 struct mới dc inject vào service và sẽ có những Behavior không mong muốn.
    Vì vậy chúng ta cần phải khống chế, chỉ cho developer xài những kiểu struct do chúng ta định
    nghĩa. -> Factory Method (Lưu ý là method, ko phải class/struct)

Solution:
    
    Dùng Factory Method để ép Developer gọi hàm tạo, trong hàm tạo này chúng ta sẽ cung cấp những
    struct/ class mà chúng ta support.

Use case:

    Logger:
        File logger: chúng ta ko muốn user tạo thêm file, vì file này đã được tạo ở 1 nơi nào đó
        Mỗi lần user muốn dùng File logger, chúng ta cung cấp Factory Method để trả về pointer 
        file logger đó

    Connection Pool:
        Get DB context :D
