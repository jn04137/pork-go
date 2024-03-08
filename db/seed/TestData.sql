-- Adminer 4.8.1 PostgreSQL 16.1 (Debian 16.1-1.pgdg120+1) dump

INSERT INTO "UserAccount" ("id", "username", "password", "email", "uuid", "active", "createdat", "updatedat") VALUES
(1,	'Nikola_Tesla',	'$2a$10$LCUDC/ddb7KYp11PffcMfedkR0PZ3eU242i7AmaFsPVSdQkqWRvNe',	'nikola_tesla@example.com',	'_fhBCU0sx-iefYX_rK7ZE',	't',	'2023-12-11 03:32:30.481928',	NULL),
(2,	'test',	'$2a$10$dQdzE07fAjx8oYkyZ3OQh.jUBxVhx80ACJvt2OTe7Yq/9f7/JrYNi',	'test@example.com',	'XF6_wPv3Ls9OaRVPNqB_u',	't',	'2023-12-11 03:32:30.536599',	NULL),
(27,	'wthunder',	'$2a$10$Jy2t7.Y4/eG5wxc/3f/oZet2BKTpB0pb4JeGRLe9Co9TZSuJZctaq',	'jnganguyen3@gmail.com',	'9a-rGsr3XeEU9RtLTSp95',	't',	'2023-12-31 04:59:25.880414',	NULL);

INSERT INTO "UserPost" ("id", "owner", "title", "body", "createdat", "updatedat", "isdeleted") VALUES
(1,	1,	'Lorem Ipsum',	'<p class=\""text-[14px] font-light\""><strong>Lorem Ipsum</strong> is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.</p>
',	'2023-12-11 03:32:30.538876',	NULL,	'f'),
(2,	1,	'Lorem Ipsum',	'<p class=\""text-[14px] font-light\""><strong>Lorem Ipsum</strong> is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.</p>
',	'2023-12-11 03:32:30.540992',	NULL,	'f'),
(3,	1,	'Lorem Ipsum',	'<p class=\""text-[14px] font-light\""><strong>Lorem Ipsum</strong> is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.</p>
',	'2023-12-11 03:32:30.542741',	NULL,	'f'),
(4,	1,	'Lorem Ipsum',	'<p class=\""text-[14px] font-light\""><strong>Lorem Ipsum</strong> is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.</p>
',	'2023-12-11 03:32:30.544554',	NULL,	'f'),
(5,	1,	'Lorem Ipsum',	'<p class=\""text-[14px] font-light\""><strong>Lorem Ipsum</strong> is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.</p>
',	'2023-12-11 03:32:30.545949',	NULL,	'f'),
(6,	1,	'Lorem Ipsum',	'<p class=\""text-[14px] font-light\""><strong>Lorem Ipsum</strong> is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.</p>
',	'2023-12-11 03:32:30.547247',	NULL,	'f'),
(7,	1,	'Lorem Ipsum',	'<p class=\""text-[14px] font-light\""><strong>Lorem Ipsum</strong> is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.</p>
',	'2023-12-11 03:32:30.548519',	NULL,	'f'),
(8,	1,	'Lorem Ipsum',	'<p class=\""text-[14px] font-light\""><strong>Lorem Ipsum</strong> is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.</p>
',	'2023-12-11 03:32:30.549722',	NULL,	'f'),
(9,	1,	'Lorem Ipsum',	'<p class=\""text-[14px] font-light\""><strong>Lorem Ipsum</strong> is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.</p>
',	'2023-12-11 03:32:30.55085',	NULL,	'f'),
(10,	1,	'Lorem Ipsum',	'<p class=\""text-[14px] font-light\""><strong>Lorem Ipsum</strong> is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.</p>
',	'2023-12-11 03:32:30.551929',	NULL,	'f'),
(11,	1,	'Lorem Ipsum',	'<p class=\""text-[14px] font-light\""><strong>Lorem Ipsum</strong> is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.</p>
',	'2023-12-11 03:32:30.552851',	NULL,	'f'),
(12,	1,	'Lorem Ipsum',	'<p class=\""text-[14px] font-light\""><strong>Lorem Ipsum</strong> is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.</p>
',	'2023-12-11 03:32:30.553784',	NULL,	'f'),
(13,	1,	'Lorem Ipsum',	'<p class=\""text-[14px] font-light\""><strong>Lorem Ipsum</strong> is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.</p>
',	'2023-12-11 03:32:30.55487',	NULL,	'f'),
(14,	1,	'Lorem Ipsum',	'<p class=\""text-[14px] font-light\""><strong>Lorem Ipsum</strong> is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.</p>
',	'2023-12-11 03:32:30.555961',	NULL,	'f'),
(15,	1,	'Lorem Ipsum',	'<p class=\""text-[14px] font-light\""><strong>Lorem Ipsum</strong> is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.</p>
',	'2023-12-11 03:32:30.556943',	NULL,	'f'),
(16,	2,	'More Lorem Ipsum',	'<p class=\""text-[14px] font-light\""><strong>Lorem Ipsum</strong> is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.</p>
',	'2023-12-11 03:32:30.557823',	NULL,	'f'),
(18,	2,	'Sample post',	'<p>This is a test of a new post with react.</p>',	'2024-01-05 02:13:56.407095',	NULL,	'f'),
(19,	2,	'First Post on Satch3l!',	'<p>Second time''s the charm! Had previously launched this application as Knostash with a Svelte frontend. The Svelte app was functional and the development experience was really nice. I made the decision to come back to React out of comfort and the larger community support + more supported libraries.</p><p></p><p>I was serving the frontend application with a server on Digital Ocean from the Golang application. After hearing about Cloudflare Pages, I was thinking that this solution would save my Digital Ocean account from bandwidth costs from the constant loading of the frontend for each user. Leveraging Cloudflare pages should offset a lot of this potential cost and it should be a more performant solution as Cloudflare''s CDN can serve the React App at nodes closer to end users.</p><p></p><p>Looking to implement a few more features such as editing existing posts, adding a points system, adding a reporting system, etc. There''s a lot more to come!</p><p>If you''ve just randomly come across this page, welcome to the site and I hope that this site comes across your browser again!</p><ul><li><p>- Jonathan Nguyen</p></li></ul><p></p>',	'2024-01-15 00:52:14.806145',	NULL,	'f'),
(20,	2,	'Testing create post',	'<p>Test message to see if creating posts are working</p>',	'2024-01-25 04:23:30.07267',	NULL,	'f'),
(22,	2,	'one more post',	'<p>one more post</p><p></p>',	'2024-01-25 04:33:45.731199',	NULL,	'f'),
(24,	2,	'one more post',	'<p>post post post</p>',	'2024-01-25 04:44:49.474914',	NULL,	't'),
(23,	2,	'Create post',	'<p>Creating a post here</p>',	'2024-01-25 04:37:45.986658',	NULL,	't'),
(21,	2,	'ANTOHER TEST POST',	'<p>Giving this post another edit to test the transition to the viewpost page after a successful edit.</p>',	'2024-01-25 04:27:06.597364',	NULL,	'f');

INSERT INTO "Comment" ("id", "owner", "body", "createdat", "updatedat") VALUES
(1,	1,	'<p class=\""text-[14px] font-light\""><strong>Lorem Ipsum</strong> is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.</p>
',	'2023-12-11 03:32:30.558521',	NULL),
(2,	1,	'<p class=\""text-[14px] font-light\""><strong>Lorem Ipsum</strong> is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.</p>
',	'2023-12-11 03:32:30.561629',	NULL),
(3,	1,	'<p class=\""text-[14px] font-light\""><strong>Lorem Ipsum</strong> is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.</p>
',	'2023-12-11 03:32:30.563381',	NULL),
(4,	1,	'<p class=\""text-[14px] font-light\""><strong>Lorem Ipsum</strong> is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.</p>
',	'2023-12-11 03:32:30.564958',	NULL),
(5,	1,	'<p class=\""text-[14px] font-light\""><strong>Lorem Ipsum</strong> is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.</p>
',	'2023-12-11 03:32:30.566397',	NULL),
(6,	1,	'<p class=\""text-[14px] font-light\""><strong>Lorem Ipsum</strong> is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.</p>
',	'2023-12-11 03:32:30.568047',	NULL),
(7,	1,	'<p class=\""text-[14px] font-light\""><strong>Lorem Ipsum</strong> is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.</p>
',	'2023-12-11 03:32:30.56922',	NULL),
(8,	1,	'<p class=\""text-[14px] font-light\""><strong>Lorem Ipsum</strong> is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.</p>
',	'2023-12-11 03:32:30.570313',	NULL),
(9,	1,	'<p class=\""text-[14px] font-light\""><strong>Lorem Ipsum</strong> is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.</p>
',	'2023-12-11 03:32:30.57135',	NULL),
(10,	1,	'<p class=\""text-[14px] font-light\""><strong>Lorem Ipsum</strong> is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.</p>
',	'2023-12-11 03:32:30.572256',	NULL),
(11,	1,	'<p class=\""text-[14px] font-light\""><strong>Lorem Ipsum</strong> is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.</p>
',	'2023-12-11 03:32:30.573212',	NULL),
(12,	1,	'<p class=\""text-[14px] font-light\""><strong>Lorem Ipsum</strong> is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.</p>
',	'2023-12-11 03:32:30.574124',	NULL),
(13,	1,	'<p class=\""text-[14px] font-light\""><strong>Lorem Ipsum</strong> is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.</p>
',	'2023-12-11 03:32:30.575094',	NULL),
(14,	1,	'<p class=\""text-[14px] font-light\""><strong>Lorem Ipsum</strong> is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.</p>
',	'2023-12-11 03:32:30.576078',	NULL),
(15,	1,	'<p class=\""text-[14px] font-light\""><strong>Lorem Ipsum</strong> is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.</p>
',	'2023-12-11 03:32:30.577004',	NULL),
(16,	2,	'<p>Trying again to get comment button to work</p>',	'2024-01-07 07:35:19.410556',	NULL),
(17,	2,	'<p>Testing the disabling of the outline. Wow. This looks pretty clean. Gonna try hitting the submit now</p>',	'2024-01-07 09:06:28.918246',	NULL),
(18,	2,	'<p>Creating a test comment</p>',	'2024-02-04 19:48:44.844657',	NULL),
(19,	2,	'<p>Creating a test comment</p>',	'2024-02-04 19:48:51.185896',	NULL),
(20,	2,	'<p>When creating a comment, will the feed refresh now?</p>',	'2024-02-04 20:49:28.241312',	NULL),
(21,	2,	'<p>Testing again. Now clearing the editor after a successful comment creation.</p>',	'2024-02-04 20:50:22.642198',	NULL),
(22,	2,	'<p>Testing once more. Will it work this time?</p>',	'2024-02-04 20:51:07.369194',	NULL),
(23,	2,	'<p>Testing the new create comment feature. Is it working?</p>',	'2024-02-04 20:51:53.548622',	NULL),
(24,	2,	'<p>How about now?</p>',	'2024-02-04 20:51:59.982118',	NULL),
(25,	2,	'<p>Trying once more. Will this clear the comment editor after a success?</p>',	'2024-02-04 20:58:30.98852',	NULL),
(26,	2,	'<p>Will this test wait for the post to populate before querying?</p>',	'2024-02-04 20:59:38.820313',	NULL),
(27,	2,	'<p>It''s not quite loading the way that I want it to. It''s clearing the comment editor correctly, but it''s kind of duplicating the last comment on the list.</p>',	'2024-02-04 21:02:04.804457',	NULL),
(28,	2,	'<p>Might just be fine for now. Have to come back to fix it.</p>',	'2024-02-04 21:03:14.566433',	NULL),
(29,	2,	'<p>Could cheat and force the page to reload instead.</p>',	'2024-02-04 21:03:44.937302',	NULL);

INSERT INTO "CommentOnPost" ("id", "commentid", "postid") VALUES
(1,	1,	16),
(2,	2,	16),
(3,	3,	16),
(4,	4,	16),
(5,	5,	16),
(6,	6,	16),
(7,	7,	16),
(8,	8,	16),
(9,	9,	16),
(10,	10,	16),
(11,	11,	16),
(12,	12,	16),
(13,	13,	16),
(14,	14,	16),
(15,	15,	16),
(16,	16,	16),
(17,	17,	18),
(18,	18,	24),
(19,	19,	24),
(20,	20,	24),
(21,	21,	24),
(22,	22,	24),
(23,	23,	23),
(24,	24,	23),
(25,	25,	23),
(26,	26,	23),
(27,	27,	24),
(28,	28,	24),
(29,	29,	24);


INSERT INTO "PointsOnPost" ("id", "userid", "postid", "point") VALUES
(10,	2,	24,	'plus'),
(11,	2,	23,	'empty'),
(12,	2,	15,	'plus'),
(8,	2,	18,	'empty'),
(9,	2,	16,	'empty'),
(7,	2,	20,	'empty'),
(6,	2,	19,	'empty');

-- 2024-03-02 00:40:52.480928+00
