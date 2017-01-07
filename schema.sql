
CREATE TABLE files(name varchar(1000) not null, autoName varchar(255) not null, user_id references user(id), created_date timestamp);
CREATE TABLE category( id integer primary key autoincrement ,name varchar(1000) not null, user_id references user(id));
INSERT INTO "category" VALUES(1,'TaskApp',1);
INSERT INTO "category" VALUES(2,'Writing',1);
INSERT INTO "category" VALUES(4,'Reading',1);
INSERT INTO "category" VALUES(5,'antiTextBook',1);
INSERT INTO "category" VALUES(7,'HolidayTasks',1);
CREATE TABLE comments(id integer primary key autoincrement, content ntext, taskID references task(id), created datetime, user_id references user(id));
INSERT INTO "comments" VALUES(58,'_hello_',1,'2017-01-04 16:26:45',1);
CREATE TABLE status (
    id integer primary key autoincrement,
    status varchar(50) not null
);
INSERT INTO "status" VALUES(1,'COMPLETE');
INSERT INTO "status" VALUES(2,'PENDING');
INSERT INTO "status" VALUES(3,'DELETED');
CREATE TABLE user (
    id integer primary key autoincrement,
    username varchar(100),
    password varchar(1000),
    email varchar(100)
);
INSERT INTO "user" VALUES(1,'suraj','suraj','sapatil@live.com');
CREATE TABLE task (
id integer primary key autoincrement,
title varchar(100),
content text,
created_date timestamp,
last_modified_at timestamp,
finish_date timestamp
, priority integer, cat_id references category(id), task_status_id references status(id), due_date timestamp, user_id references user(id), hide int);
INSERT INTO "task" VALUES(1,'Publish on github','Publish the source of tasks and picsort on github','2015-11-12 15:30:59','2016-12-12 06:11:10',NULL,3,1,2,NULL,1,0);
INSERT INTO "task" VALUES(4,'gofmtall','The idea is to run gofmt -w file.go on every go file in the listing, *Edit turns out this is is difficult to do in golang **Edit barely 3 line bash script :P ','2015-11-12 16:58:31','2016-12-12 06:09:06','2015-11-13 13:16:48',3,1,3,NULL,1,0);
INSERT INTO "task" VALUES(7,'webdev anti textbook mods','- [x] handling errors rather than just printing them
- [x] non-cookie based authentication schemes such as JWT
- [x] When explaining how to make a form, you should explain CSRF and white-listing of parameters.
  done >  When explaining how to capture GET vs POST methods, be sure to explain the many security implications. (GET parameters are logged by reverse proxies and browser histories, web accelerators request GET w/o user clicking, GET is more likely to be cached, etc.) Even when explaining web server development, explain the security implications of binding to 127.0.0.1 vs 0.0.0.0. (I.e. when developing, do you trust everyone on the local network? What if you open your laptop in a cafe that happens to be across the street from Defcon?) Explain concepts like "never use a client filename without security scrubbing" and "security problems when using UTF-8"
- [x] Have a chapter on middleware and using now standard context package. Touch on templates and strategies with templates (unless you want to limit the book to RESTful services). Mention advanced topics (do not necessarily expand on them though but keep readers appraised of possibilities) like content negotiation, cookies, HTTP2, custom headers
- [x] Routing is a big deal if you''re making a web application or a server and I think that you should focus more on that. At least separate routing logic from your handlers. Seeing your handler check the method and/or process the URI was a bit painful since that is a clear job of the router.
- [x] I did not see form submission/validation. I would even show how to submit a file.
- [x] Sessions/cookies are very important to a web application','2015-11-13 04:23:27','2016-12-12 06:11:11',NULL,3,5,2,NULL,1,0);
INSERT INTO "task" VALUES(8,'PR for godoc','Print a message when the server starts, because it is irritating when I have no idea when the server has started','2015-11-13 06:54:53','2016-12-12 06:11:12',NULL,3,NULL,2,NULL,1,0);
INSERT INTO "task" VALUES(14,'Issues in vscode','1. more options for replace all, replace all in session?
2. fold functions and code
3. way to do alt+tab to change the current file
4. when I close a working file, the next file in the buffer should be opened rather than blank window','2015-11-21 04:25:55','2016-12-12 06:11:12',NULL,3,1,2,NULL,1,0);
INSERT INTO "task" VALUES(16,'Reading list','- [x] India''s biggest coverup
- [x] Yuganta
- [x] Gita
- [x] the mahabharata secret
- [x] Rise and fall of the third reich
- [x] The hard things about the hard things
- [x] the power of your subconscious mind
- [x] By way of deception
- [x] The coca cola way
- [x] Predictably irrational
- [x] Extreme management
- [x] rework
- [ ] conspiracy of fools
- [ ] security analysis
- [ ] common stocks uncommon profits
- [ ] wealth of nations
- [ ] clash of civilization and reorder
- [ ] HBR on entrepreneurship
- [ ] to kill a mocking bird
- [ ] Illiad
- [ ] A time to kill
- [ ] world economy
- [ ] economics of business enterprise
- [ ] Alexander the great in fact and fiction
- [ ] between mind and nature
- [ ] french revolution
- [ ] idea mapping','2015-12-06 05:37:06','2016-02-15 08:20:21',NULL,1,4,2,NULL,1,1);
INSERT INTO "task" VALUES(19,'Chapter list','[Introduction](README.md)

* Installation and Tools

   * [Installation](content/0.0install.md)

   * [Tools](content/0.1tools.md)

* General

   * [Go Programming Basics](content/1.0general_talk.md)

   * [Sneak Peek](content/1.1servers.md)

* Implementation

   * [Implementation Basics](content/2.0implementbasics.md)

   * [Basic Functionality](content/2.1functionality.md)

   * [Database Handling](content/2.2database.md)

   * [Database Example](content/2.3example.md)

   * [Closing Remarks](content/2.4closingremarks.md)

* Templates

   * [Basics of Templates](   * [Basics of Templates](  [User Authentication](content/4.0authentication.md)

* [Contributors](CONTRIBUTORS.md)
','2016-01-09 05:38:54','2016-01-09 06:16:51',NULL,3,1,3,NULL,1,0);
INSERT INTO "task" VALUES(20,'Tasks modifications','- [x] create file names & store them in database, while showing link fetch the values from db
- [x] random file name
- [x] http handler on /files URL, will need to check sessions first
- [x] markdown
- [x] categories
- [x] edit/update/delete categories
','2016-01-11 15:50:12','2016-02-06 07:23:35','2016-02-06 07:23:35',3,1,3,NULL,1,0);
INSERT INTO "task" VALUES(21,'Writing challenge','- [ ] Write a short story that has a character that is afraid of many things in the world, however is sent on a quest to recover a person or object that is very dear to them.
- [ ] Whenever you get chills, you just died in an alternate universe.','2016-01-26 16:45:48','2016-03-19 04:36:05',NULL,1,2,3,NULL,1,0);
INSERT INTO "task" VALUES(22,'TasksMods: Feb6','- [ ] add a notification queue
- [x] redirect should work to the `referer()` for each page
- [x] all forms''s `Submit` button should have the submit text
- [x] search icon should be visible for each page
- [x] when we are in category page, `category` should be pre selected
- [x] `add file` link to toggle the file upload bar
- [x] edit''s close should redirect to `referer`
- [ ] modify the editor & ','2016-02-06 07:23:19','2016-02-12 16:15:48','2016-02-12 16:15:48',3,1,3,NULL,1,0);
INSERT INTO "task" VALUES(23,'Tasks-Wishlist','- [x] Multi user support
- [ ] More filtering options on front page
- [ ] Show completed Tasks inside categories
- [ ] <s>remove generic /deleted /completed links and do them on the basis of category</s>
- [ ] convert everything to AJAX
- [ ] proper test cases for everything
- [ ] <s>Modify the way we send notifications</s>
- [x] edit/delete comments?
- [x] timestamp for comments
- [x] images should be inline
- [x] display which category the note belongs on the general view
- [x] think of adding due dates :P
- [ ] <s>Recent comments?</s>
- [ ] <s>Replace Bootstrap with semantic UI?</s>','2016-02-15 05:01:49','2016-12-12 06:11:13',NULL,3,1,2,NULL,1,0);
INSERT INTO "task" VALUES(25,'Authentication','1. Take input from form &amp; read cookie from filesystem for a sessionID
1. Check if a sessionID is valid/create a sessionID and store it in cookie
1. Use a key value store to store the sessions and fetch the sessionID in each request rather than have a global map which stores the session information

use boltdb or some key value store, we&#39;ll check which one to use later, but the general logic is this, if we are on target then we can even write a library','2016-02-21 17:20:15','2016-12-12 06:11:13',NULL,3,1,2,NULL,1,0);
INSERT INTO "task" VALUES(26,'Bookmark manager of firefox','A webapplication in Go for managing bookmarks in firefox/chrome tagging bookmarks is hell, so read the sqlite db, store the tagging in our db rather than relying on the tags of firefox a simple UI, first reads all bookmarks, provides a sleek interface for the bookmarks, probably as boards on the domain/ day they were bookmarked and categories allows us to go categories each bookmark','2016-02-21 17:25:25','2016-03-19 04:42:19',NULL,3,6,3,NULL,1,0);
INSERT INTO "task" VALUES(27,'Ping via go','Ping 8.8.8.8 via go if you don&#39;t get a destination unreachable message then play some audio','2016-02-27 07:07:51','2016-03-19 04:41:52',NULL,2,6,3,NULL,1,0);
INSERT INTO "task" VALUES(28,'ML Action Item','- [x]  os3.pdf
- [ ]  ISLR.pdf
- [ ]  Practice with R
','2016-06-01 02:54:29','2016-12-12 06:10:53',NULL,3,7,3,NULL,1,0);
INSERT INTO "task" VALUES(32,'Modifications to the book','- [ ] would be to show how to talk to a DB over http or sockets because a scaled web application will not have its own DB on the same box. You could also show how to do basic caching.
- [ ] Modern web applications are usually written in JS so I think that you could have an entire chapter on that. At the very least, show how to route an Ajax request to a handler
- [ ] Internationalization and Localization ','2016-08-07 09:00:03','2016-08-22 16:16:02','2016-08-22 16:16:02',3,5,1,NULL,1,0);
INSERT INTO "task" VALUES(33,'Release V1','- [x] Convert the formatting from gitbook to leanpub
- [x] Make this as a Go repository
- [x] make images smaller
- [x] change file names
- [x] upgrade Go from 1.5 to 1.7, all examples
- [x] chapter on git
- [x] chapter on db should be coherent, it is all a mess right now.
- [ ] AJAX! and JS
- [ ] Valid code examples from start to end
- [ ] rewrite the task app?
- [x] updated API chapter
- [ ] a command line REST client
- [ ] extensive modifications to the chapters
- [x] chapter on unit testing in detail
- [ ] new features of Go like HTTP2
- [ ] advance topics like content negotiation & compression
- [ ] socket programming (Found an excellent book, asked the author for permission, he gave it. IT is CC licensed, the same as this book https://jan.newmarch.name/go/)
- [ ] websockets
- [ ] working with XML and file IO in general
- [ ] NoSQL chapter','2016-08-22 16:15:49','2016-10-11 11:52:19','2016-10-11 11:52:19',3,5,1,NULL,1,0);
DELETE FROM sqlite_sequence;
INSERT INTO "sqlite_sequence" VALUES('category',9);
INSERT INTO "sqlite_sequence" VALUES('comments',58);
INSERT INTO "sqlite_sequence" VALUES('status',4);
INSERT INTO "sqlite_sequence" VALUES('user',1);
INSERT INTO "sqlite_sequence" VALUES('task',39);
COMMIT;
