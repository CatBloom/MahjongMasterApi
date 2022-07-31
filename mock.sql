insert into leagues(id,name,manual,start_at,finish_at)
values
('e952fe14-5ddd-4b3e-abf0-8ebd9d57a22a','sample01','sampleManual','2022/06/30','2022/07/01')

insert into players(name,league_id)
values
('sample01','e952fe14-5ddd-4b3e-abf0-8ebd9d57a22a'),
('sample02','e952fe14-5ddd-4b3e-abf0-8ebd9d57a22a'),
('sample03','e952fe14-5ddd-4b3e-abf0-8ebd9d57a22a'),
('sample04','e952fe14-5ddd-4b3e-abf0-8ebd9d57a22a'),
('sample05','e952fe14-5ddd-4b3e-abf0-8ebd9d57a22a'),
('sample06','e952fe14-5ddd-4b3e-abf0-8ebd9d57a22a'),
('sample07','e952fe14-5ddd-4b3e-abf0-8ebd9d57a22a'),
('sample08','e952fe14-5ddd-4b3e-abf0-8ebd9d57a22a')

insert into games(id,league_id)
values
(1,'8f775ab526474978838dc5745aa069ef'),
(2,'8f775ab526474978838dc5745aa069ef'),
(3,'8f775ab526474978838dc5745aa069ef'),
(4,'8f775ab526474978838dc5745aa069ef')

insert into games_players(game_id,player_id)
values
(1,1),
(1,2),
(1,3),
(1,4),
(2,1),
(2,2),
(2,3),
(2,4),
(3,3),
(3,4),
(3,5),
(3,6),
(4,1),
(4,2),
(4,7),
(4,8)

insert into results(player_id,rank,point,calc_point,game_id)
values
(1,1,60000,60,1),
(2,2,50000,40,1),
(3,3,30000,-20,1),
(4,4,-40000,-80,1),
(3,1,40000,40,2),
(4,2,30000,20,2),
(1,3,20000,-20,2),
(2,4,10000,-40,2),
(3,1,45000,45,3),
(4,2,35000,30,3),
(5,3,20000,-30,3),
(6,4,0,-45,3),
(1,1,50000,60,4),
(2,2,40000,20,4),
(7,3,30000,-20,4),
(8,4,-20000,-60,4)
