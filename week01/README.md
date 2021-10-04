# Submission Assignment Week 01

## My History in Technology

I started out my career in Information Technology in the early 2000s. My first job was at a research institute as a programmer, first in Perl and then in Java. After a few years, I got the opportunity to switch to Linux system administration. Being in the middle of the [dot-com bubble](https://en.wikipedia.org/wiki/Dot-com_bubble) I felt that system administration would provide more job security. There were hardly any programming jobs available and system administration is something that always needs to be done, was my reasoning. So I was glad I could make the change.

So for the next few years, I worked as a Linux system administrator. Pretty soon I got in touch with servers with Oracle databases running on them, and that caught my interest. I soon started specializing in that and I worked as a DBA roughly until 2015. During that time, my programming experience was limited to shell and Perl scripting, and some Python programming.

Around 2015 I felt that the field of Oracle databases did not provide me with many challenges anymore. The main features of the database software hardly changed, hence the work hardly changed. Luckily, I got the opportunity to start working with the then-emerging public cloud providers. This was a chance I was very eager to take. I must say that this has most certainly provided me with enough challenges. New features and technologies are implemented at a rapid pace, Kubernetes being one of the more dominant ones.

The way that Go fits in my plans for the future, is that I want to be able to write more complex software. When I'm able to do this, it provides me the possibility to automate more complex tasks. But I also hope that I get more feeling for the daily struggles of the programmers that I work with.

I really enjoyed programming when I did that at the start of my career, so I hope to have a lot of fun while learning to program in Go.

## Exploring the `fmt` package

For the assignment of week 01, I had to explore the `fmt` package and describe any surprises that I had during the exploration. I must say that there were hardly any surprises for me.

I find the structure of the package very straightforward. The structure of the documentation is also straightforward and predictable. The verbs that you can use to format the printing are listed directly at the top of the documentation page. When you scroll down, you will immediately find an example on how to use these verbs in a `fmt.Printf()` statement. Further down, even more examples are listed.

The Go documentation reminds me very much of the Javadoc documentation that I worked with while learning Java. Printing 'Hello World' in Java would be something along the lines of:

```java
System.out.println("Hello World")
```

The big difference that I see here between Java and Go is that for Go, the import of the `fmt` package is explicitly needed. For Java, you do not need to import the System class because it is always available.

I cannot remember whether I ever got to work with formatted printing in Java as is described on [this page](https://docs.oracle.com/javase/tutorial/java/data/numberformat.html). So I cannot make a decent comparison on formatted printing between Java and Go.

The first programming language that I got to work with is Perl, in which formatted printing is done as follows:

```perl
 #!/usr/bin/perl
 $a =  "text";
 printf("%s\n", $a);
 ```

 This works out of the box; no imports or packages needed.
