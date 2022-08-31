select:<br>
The select statement is used to choose from multiple send/receive channel operations<br><br>
The select statement is used to choose from multiple send/receive channel operations<br>
The default case in a select statement is executed when none of the other cases is ready. This is generally used to prevent the select statement from blocking.
<br><br>
When multiple cases in a select statement are ready, one of them will be executed at random.

<br>**empty select**
block forever resulting in a deadlock
