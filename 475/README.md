首先这道题的核心思路是  对于一个房子，他要么选前面的供暖气，要么选后面的供暖气

这2个供暖气中选最小的那个，但是在每个房子与每个房子直接选最大的那个，就可以保证所有房子都能覆盖到

![image-20211220181458718](./image-20211220181458718.png)

上述都是按照排序后的思想去做的，所有对于房子和供暖都得先进行一下排序，`关于排序就先直接使用库`