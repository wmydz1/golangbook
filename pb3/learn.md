²  repeated：在一个格式良好的消息中，这种字段可以重复任意多次（包括0次）。重复的值的顺序会被保留。表示该值可以重复，相当于java中的List。

Repeated：表示该字段可以包含0~N个元素。其特性和optional一样，但是每一次可以包含多个值。可以看作是在传递一个数组的值。
 
 protobuf 消息的repeated字段，可以包含0~N个相同的内容。
 
 当包含的内容大于0时，可以认为是在修改数据或者数据有改变。
 当包含的内容是0时，也就是不包含时，究竟是不改变原来的数据，还是清空呢？
 
 因此在设计协议时，遇到repeated字段时，最好在与某个optional字段相组合，用来指示是否包含相应的repeated字段。
 例如。
 message MSG1
 {
    repeated fixed32 data = 1;
    optional bool include_data = 2
 }
  
 当 include_data = true 时，表示要修改data字段，此时如果data字段包含0个元素则是清空。
 如果include_data = false ，时表示对data字段不做任何改变，此时data字段包含0个元素或者多个元素在业务逻辑上都应该被忽略。
 
 以上方法可以解决了repeated字段的二义性问题。
 
 
 是的，repeated 就是说该字段是指定类型的数组。java里面是个list
 